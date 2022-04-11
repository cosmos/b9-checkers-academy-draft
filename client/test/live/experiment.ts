import { readFile } from "fs/promises"
import { DirectSecp256k1HdWallet, OfflineDirectSigner } from "@cosmjs/proto-signing"
import Long from "long"
import { CheckersStargateClient } from "../../src/checkers_stargateclient"
import { CheckersSigningStargateClient } from "../../src/checkers_signingstargateclient"
import { DeliverTxResponse, GasPrice } from "@cosmjs/stargate"
import { QueryCanPlayMoveResponse } from "../../src/types/generated/checkers/query"

const starportEndpoint = "http://localhost:26657"

const getSignerFromMnemonic = async (filePath: string): Promise<OfflineDirectSigner> => {
    return DirectSecp256k1HdWallet.fromMnemonic((await readFile(filePath)).toString(), {
        prefix: "cosmos",
    })
}

const getGameId = (createResponse: DeliverTxResponse): string => {
    return JSON.parse(createResponse.rawLog!)[0].events[0].attributes.find(
        (eventInfo: { key: string }) => eventInfo.key == "Index",
    ).value
}

async function runAll() {
    const client: CheckersStargateClient = await CheckersStargateClient.connect(starportEndpoint)
    const checkers = client.checkersQueryClient!.checkers

    // Initial NextGame
    const nextGame0 = await checkers.getNextGame()
    console.log("NextGame:", nextGame0, ", idValue:", nextGame0.idValue.toString(10))

    // All Games
    const allGames0 = await checkers.getAllStoredGames(
        Uint8Array.of(),
        Long.fromInt(0),
        Long.fromInt(0),
        true,
    )
    console.log("All games", allGames0, ", total: ", allGames0.pagination!.total.toString(10))

    // Prepare keys
    const aliceSigner: OfflineDirectSigner = await getSignerFromMnemonic(`${__dirname}/alice.key`)
    const alice: string = (await aliceSigner.getAccounts())[0].address
    const bobSigner: OfflineDirectSigner = await getSignerFromMnemonic(`${__dirname}/bob.key`)
    const bob: string = (await bobSigner.getAccounts())[0].address
    const aliceSigningClient: CheckersSigningStargateClient =
        await CheckersSigningStargateClient.connectWithSigner(starportEndpoint, aliceSigner, {
            gasPrice: GasPrice.fromString("1stake"),
        })
    const bobSigningClient: CheckersSigningStargateClient =
        await CheckersSigningStargateClient.connectWithSigner(starportEndpoint, bobSigner, {
            gasPrice: GasPrice.fromString("1stake"),
        })

    // Create game
    const createResponse: DeliverTxResponse = await aliceSigningClient.createGame(
        alice,
        bob,
        alice,
        "stake",
        Long.fromNumber(5),
        "auto",
    )
    console.log(createResponse)
    const createdGameId = getGameId(createResponse)

    // Play game
    const canPlay: QueryCanPlayMoveResponse = await client.checkersQueryClient!.checkers.canPlayMove(
        createdGameId,
        "b",
        { x: 1, y: 2 },
        { x: 2, y: 3 },
    )
    console.log(canPlay)

    const playResponse: DeliverTxResponse = await bobSigningClient.playMove(
        bob,
        createdGameId,
        { x: 1, y: 2 },
        { x: 2, y: 3 },
        "auto",
    )
    console.log(playResponse)

    // Reject game
    const rejectResponse: DeliverTxResponse = await aliceSigningClient.rejectGame(
        alice,
        createdGameId,
        "auto",
    )
    console.log(rejectResponse)
}

runAll()
