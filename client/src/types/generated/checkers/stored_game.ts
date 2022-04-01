/* eslint-disable */
import Long from "long"
import _m0 from "protobufjs/minimal"

export const protobufPackage = "xavierlepretre.checkers.checkers"

export interface StoredGame {
    creator: string
    index: string
    game: string
    turn: string
    red: string
    black: string
    moveCount: Long
    /** Pertains to the FIFO. Towards head. */
    beforeId: string
    /** Pertains to the FIFO. Towards tail. */
    afterId: string
    deadline: string
    winner: string
    wager: Long
    /** Denomination of the wager. */
    token: string
}

function createBaseStoredGame(): StoredGame {
    return {
        creator: "",
        index: "",
        game: "",
        turn: "",
        red: "",
        black: "",
        moveCount: Long.UZERO,
        beforeId: "",
        afterId: "",
        deadline: "",
        winner: "",
        wager: Long.UZERO,
        token: "",
    }
}

export const StoredGame = {
    encode(message: StoredGame, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator)
        }
        if (message.index !== "") {
            writer.uint32(18).string(message.index)
        }
        if (message.game !== "") {
            writer.uint32(26).string(message.game)
        }
        if (message.turn !== "") {
            writer.uint32(34).string(message.turn)
        }
        if (message.red !== "") {
            writer.uint32(42).string(message.red)
        }
        if (message.black !== "") {
            writer.uint32(50).string(message.black)
        }
        if (!message.moveCount.isZero()) {
            writer.uint32(56).uint64(message.moveCount)
        }
        if (message.beforeId !== "") {
            writer.uint32(66).string(message.beforeId)
        }
        if (message.afterId !== "") {
            writer.uint32(74).string(message.afterId)
        }
        if (message.deadline !== "") {
            writer.uint32(82).string(message.deadline)
        }
        if (message.winner !== "") {
            writer.uint32(90).string(message.winner)
        }
        if (!message.wager.isZero()) {
            writer.uint32(96).uint64(message.wager)
        }
        if (message.token !== "") {
            writer.uint32(106).string(message.token)
        }
        return writer
    },

    decode(input: _m0.Reader | Uint8Array, length?: number): StoredGame {
        const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input)
        let end = length === undefined ? reader.len : reader.pos + length
        const message = createBaseStoredGame()
        while (reader.pos < end) {
            const tag = reader.uint32()
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string()
                    break
                case 2:
                    message.index = reader.string()
                    break
                case 3:
                    message.game = reader.string()
                    break
                case 4:
                    message.turn = reader.string()
                    break
                case 5:
                    message.red = reader.string()
                    break
                case 6:
                    message.black = reader.string()
                    break
                case 7:
                    message.moveCount = reader.uint64() as Long
                    break
                case 8:
                    message.beforeId = reader.string()
                    break
                case 9:
                    message.afterId = reader.string()
                    break
                case 10:
                    message.deadline = reader.string()
                    break
                case 11:
                    message.winner = reader.string()
                    break
                case 12:
                    message.wager = reader.uint64() as Long
                    break
                case 13:
                    message.token = reader.string()
                    break
                default:
                    reader.skipType(tag & 7)
                    break
            }
        }
        return message
    },

    fromJSON(object: any): StoredGame {
        return {
            creator: isSet(object.creator) ? String(object.creator) : "",
            index: isSet(object.index) ? String(object.index) : "",
            game: isSet(object.game) ? String(object.game) : "",
            turn: isSet(object.turn) ? String(object.turn) : "",
            red: isSet(object.red) ? String(object.red) : "",
            black: isSet(object.black) ? String(object.black) : "",
            moveCount: isSet(object.moveCount) ? Long.fromString(object.moveCount) : Long.UZERO,
            beforeId: isSet(object.beforeId) ? String(object.beforeId) : "",
            afterId: isSet(object.afterId) ? String(object.afterId) : "",
            deadline: isSet(object.deadline) ? String(object.deadline) : "",
            winner: isSet(object.winner) ? String(object.winner) : "",
            wager: isSet(object.wager) ? Long.fromString(object.wager) : Long.UZERO,
            token: isSet(object.token) ? String(object.token) : "",
        }
    },

    toJSON(message: StoredGame): unknown {
        const obj: any = {}
        message.creator !== undefined && (obj.creator = message.creator)
        message.index !== undefined && (obj.index = message.index)
        message.game !== undefined && (obj.game = message.game)
        message.turn !== undefined && (obj.turn = message.turn)
        message.red !== undefined && (obj.red = message.red)
        message.black !== undefined && (obj.black = message.black)
        message.moveCount !== undefined && (obj.moveCount = (message.moveCount || Long.UZERO).toString())
        message.beforeId !== undefined && (obj.beforeId = message.beforeId)
        message.afterId !== undefined && (obj.afterId = message.afterId)
        message.deadline !== undefined && (obj.deadline = message.deadline)
        message.winner !== undefined && (obj.winner = message.winner)
        message.wager !== undefined && (obj.wager = (message.wager || Long.UZERO).toString())
        message.token !== undefined && (obj.token = message.token)
        return obj
    },

    fromPartial<I extends Exact<DeepPartial<StoredGame>, I>>(object: I): StoredGame {
        const message = createBaseStoredGame()
        message.creator = object.creator ?? ""
        message.index = object.index ?? ""
        message.game = object.game ?? ""
        message.turn = object.turn ?? ""
        message.red = object.red ?? ""
        message.black = object.black ?? ""
        message.moveCount =
            object.moveCount !== undefined && object.moveCount !== null
                ? Long.fromValue(object.moveCount)
                : Long.UZERO
        message.beforeId = object.beforeId ?? ""
        message.afterId = object.afterId ?? ""
        message.deadline = object.deadline ?? ""
        message.winner = object.winner ?? ""
        message.wager =
            object.wager !== undefined && object.wager !== null ? Long.fromValue(object.wager) : Long.UZERO
        message.token = object.token ?? ""
        return message
    },
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined

export type DeepPartial<T> = T extends Builtin
    ? T
    : T extends Long
    ? string | number | Long
    : T extends Array<infer U>
    ? Array<DeepPartial<U>>
    : T extends ReadonlyArray<infer U>
    ? ReadonlyArray<DeepPartial<U>>
    : T extends {}
    ? { [K in keyof T]?: DeepPartial<T[K]> }
    : Partial<T>

type KeysOfUnion<T> = T extends T ? keyof T : never
export type Exact<P, I extends P> = P extends Builtin
    ? P
    : P & { [K in keyof P]: Exact<P[K], I[K]> } & Record<Exclude<keyof I, KeysOfUnion<P>>, never>

if (_m0.util.Long !== Long) {
    _m0.util.Long = Long as any
    _m0.configure()
}

function isSet(value: any): boolean {
    return value !== null && value !== undefined
}
