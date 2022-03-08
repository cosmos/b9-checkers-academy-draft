import { writeFile } from "fs/promises"
import { Server } from "http"
import express, { Express, Request, Response } from "express"
import { Block, IndexedTx } from "@cosmjs/stargate"
import { sha256 } from "@cosmjs/crypto"
import { toHex } from "@cosmjs/encoding"
import {
    ABCIMessageLog,
    Attribute,
    StringEvent,
} from "cosmjs-types/cosmos/base/abci/v1beta1/abci"
import { DbType, PlayerInfo } from "./types"
import { MyStargateClient } from "./mystargateclient"

export const createIndexer = async () => {
    const port = "3001"
    const dbFile = "./db.json"
    const db: DbType = require(dbFile)
    const pollIntervalMs = 5_000 // 5 seconds
    let timer: NodeJS.Timer | undefined
    const rpcPoint = "http://localhost:26657"
    // Web socket https://gist.github.com/findolor/48e02750e24045e65ee59721618623ea
    let client: MyStargateClient

    const app: Express = express()
    app.get("/", (req: Request, res: Response) => {
        res.send({
            error: "Nothing here",
        })
    })

    app.get("/status", (req: Request, res: Response) => {
        res.json({
            block: {
                height: db.status.block.height,
            },
        })
    })

    app.get("/players/:playerAddress", (req: Request, res: Response) => {
        res.json({
            gameCount:
                db.players[req.params.playerAddress]?.gameIds?.length ?? 0,
            gameIds: db.players[req.params.playerAddress]?.gameIds ?? [],
        })
    })

    app.get(
        "/players/:playerAddress/gameIds",
        (req: Request, res: Response) => {
            res.json(db.players[req.params.playerAddress]?.gameIds ?? [])
        }
    )

    app.patch("/games/:gameId", (req: Request, res: Response) => {
        res.json({
            result: "thank you",
        })
    })

    const saveDb = async () => {
        await writeFile(dbFile, JSON.stringify(db, null, 4))
    }

    const init = async () => {
        client = await MyStargateClient.connect(rpcPoint)
        console.log("Connected to chain-id:", await client.getChainId())
        setTimeout(poll, 1)
    }

    const getAttributeValueByKey = (
        attributes: Attribute[],
        key: string
    ): string | undefined => {
        return attributes.find((attribute: Attribute) => attribute.key == key)
            ?.value
    }

    const handleEventCreate = async (event: StringEvent): Promise<void> => {
        const newId: string | undefined = getAttributeValueByKey(
            event.attributes,
            "Index"
        )
        if (!newId) throw `Create event missing newId`
        const blackAddress: string | undefined = getAttributeValueByKey(
            event.attributes,
            "Black"
        )
        if (!blackAddress) throw `Create event missing blackAddress`
        const redAddress: string | undefined = getAttributeValueByKey(
            event.attributes,
            "Red"
        )
        if (!redAddress) throw `Create event missing redAddress`
        console.log(
            `New game: ${newId}, black: ${blackAddress}, red: ${redAddress}`
        )
        const blackInfo: PlayerInfo = db.players[blackAddress] ?? {
            gameIds: [],
        }
        const redInfo: PlayerInfo = db.players[redAddress] ?? {
            gameIds: [],
        }
        if (blackInfo.gameIds.indexOf(newId) < 0) blackInfo.gameIds.push(newId)
        if (redInfo.gameIds.indexOf(newId) < 0) redInfo.gameIds.push(newId)
        db.players[blackAddress] = blackInfo
        db.players[redAddress] = redInfo
        db.games[newId] = {
            redAddress: redAddress,
            blackAddress: blackAddress,
        }
    }

    const handleEventReject = async (event: StringEvent): Promise<void> => {
        const rejectedId: string | undefined = getAttributeValueByKey(
            event.attributes,
            "IdValue"
        )
        if (!rejectedId) throw `Reject event missing rejectedId`
        const blackAddress: string | undefined =
            db.games[rejectedId]?.blackAddress
        const redAddress: string | undefined = db.games[rejectedId]?.redAddress
        const blackGames: string[] = db.players[blackAddress]?.gameIds ?? []
        const redGames: string[] = db.players[redAddress]?.gameIds ?? []
        console.log(
            `Reject game: ${rejectedId}, black: ${blackAddress}, red: ${redAddress}`
        )
        const indexInBlack: number = blackGames.indexOf(rejectedId)
        if (0 <= indexInBlack) blackGames.splice(indexInBlack, 1)
        const indexInRed: number = redGames.indexOf(rejectedId)
        if (0 <= indexInRed) redGames.splice(indexInRed, 1)
    }

    const handleEventForfeit = async (event: StringEvent): Promise<void> => {
        const forfeitedId: string | undefined = getAttributeValueByKey(
            event.attributes,
            "IdValue"
        )
        if (!forfeitedId) throw `Forfeit event missing forfeitedId`
        const winner: string | undefined = getAttributeValueByKey(
            event.attributes,
            "Winner"
        )
        const blackAddress: string | undefined =
            db.games[forfeitedId]?.blackAddress
        const redAddress: string | undefined = db.games[forfeitedId]?.redAddress
        const blackGames: string[] = db.players[blackAddress]?.gameIds ?? []
        const redGames: string[] = db.players[redAddress]?.gameIds ?? []
        console.log(
            `Forfeit game: ${forfeitedId}, black: ${blackAddress}, red: ${redAddress}, winner: ${winner}`
        )
        const indexInBlack: number = blackGames.indexOf(forfeitedId)
        if (0 <= indexInBlack) blackGames.splice(indexInBlack, 1)
        const indexInRed: number = redGames.indexOf(forfeitedId)
        if (0 <= indexInRed) redGames.splice(indexInRed, 1)
        if (winner == "NO_PLAYER") {
            delete db.games[forfeitedId]
        }
    }

    const handleEvent = async (event: StringEvent): Promise<void> => {
        const isActionOf = (actionValue: string): boolean =>
            event.attributes.some(
                (attribute: Attribute) =>
                    attribute.key === "action" && attribute.value == actionValue
            )
        if (isActionOf("NewGameCreated")) {
            await handleEventCreate(event)
        }
        if (isActionOf("GameRejected")) {
            await handleEventReject(event)
        }
        if (isActionOf("GameForfeited")) {
            await handleEventForfeit(event)
        }
    }

    const handleEvents = async (events: StringEvent[]): Promise<void> => {
        try {
            const myEvents: StringEvent[] = events
                .filter((event: StringEvent) => event.type == "message")
                .filter((event: StringEvent) =>
                    event.attributes.some(
                        (attribute: Attribute) =>
                            attribute.key == "module" &&
                            attribute.value == "checkers"
                    )
                )
            let eventIndex = 0
            while (eventIndex < myEvents.length) {
                await handleEvent(myEvents[eventIndex])
                eventIndex++
            }
        } catch (e) {
            // Skipping if the handling failed. Most likely the transaction failed.
        }
    }

    const handleTx = async (indexed: IndexedTx) => {
        const events: StringEvent[] = JSON.parse(indexed.rawLog).flatMap(
            (log: ABCIMessageLog) => log.events
        )
        await handleEvents(events)
    }

    const handleBlock = async (block: Block) => {
        if (0 < block.txs.length) console.log("")
        let txIndex = 0
        while (txIndex < block.txs.length) {
            const txHash: string = toHex(
                sha256(block.txs[txIndex])
            ).toUpperCase()
            const indexed: IndexedTx | null = await client.getTx(txHash)
            if (!indexed) throw `Could not find indexed tx: ${txHash}`
            await handleTx(indexed)
            txIndex++
        }
        const events: StringEvent[] = await client.getEndBlockEvents(
            block.header.height
        )
        if (0 < events.length) console.log("")
        await handleEvents(events)
    }

    const poll = async () => {
        const currentHeight = await client.getHeight()
        if (db.status.block.height <= currentHeight - 100)
            console.log(
                `Catching up ${db.status.block.height}..${currentHeight}`
            )
        while (db.status.block.height < currentHeight) {
            const processing = db.status.block.height + 1
            process.stdout.cursorTo(0)
            const block: Block = await client.getBlock(processing)
            process.stdout.write(
                `Handling block: ${processing} with ${block.txs.length} txs`
            )
            await handleBlock(block)
            db.status.block.height = processing
        }
        await saveDb()
        timer = setTimeout(poll, pollIntervalMs)
    }

    process.on("SIGINT", () => {
        if (timer) clearTimeout(timer)
        saveDb()
            .then(() => {
                console.log(`${dbFile} saved`)
            })
            .catch(console.error)
            .finally(() => {
                server.close(() => {
                    console.log("server closed")
                    process.exit(0)
                })
            })
    })

    const server: Server = app.listen(port, () => {
        init()
            .catch(console.error)
            .then(() => {
                console.log(`\nserver started at http://localhost:${port}`)
            })
    })
}
