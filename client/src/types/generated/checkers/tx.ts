/* eslint-disable */
import Long from "long"
import _m0 from "protobufjs/minimal"

export const protobufPackage = "xavierlepretre.checkers.checkers"

/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgRejectGame {
    creator: string
    idValue: string
}

export interface MsgRejectGameResponse {}

export interface MsgPlayMove {
    creator: string
    idValue: string
    fromX: Long
    fromY: Long
    toX: Long
    toY: Long
}

export interface MsgPlayMoveResponse {
    idValue: string
    capturedX: Long
    capturedY: Long
    winner: string
}

export interface MsgCreateGame {
    creator: string
    red: string
    black: string
    wager: Long
    /** Denomination of the wager. */
    token: string
}

export interface MsgCreateGameResponse {
    idValue: string
}

function createBaseMsgRejectGame(): MsgRejectGame {
    return { creator: "", idValue: "" }
}

export const MsgRejectGame = {
    encode(message: MsgRejectGame, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator)
        }
        if (message.idValue !== "") {
            writer.uint32(18).string(message.idValue)
        }
        return writer
    },

    decode(input: _m0.Reader | Uint8Array, length?: number): MsgRejectGame {
        const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input)
        let end = length === undefined ? reader.len : reader.pos + length
        const message = createBaseMsgRejectGame()
        while (reader.pos < end) {
            const tag = reader.uint32()
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string()
                    break
                case 2:
                    message.idValue = reader.string()
                    break
                default:
                    reader.skipType(tag & 7)
                    break
            }
        }
        return message
    },

    fromJSON(object: any): MsgRejectGame {
        return {
            creator: isSet(object.creator) ? String(object.creator) : "",
            idValue: isSet(object.idValue) ? String(object.idValue) : "",
        }
    },

    toJSON(message: MsgRejectGame): unknown {
        const obj: any = {}
        message.creator !== undefined && (obj.creator = message.creator)
        message.idValue !== undefined && (obj.idValue = message.idValue)
        return obj
    },

    fromPartial<I extends Exact<DeepPartial<MsgRejectGame>, I>>(object: I): MsgRejectGame {
        const message = createBaseMsgRejectGame()
        message.creator = object.creator ?? ""
        message.idValue = object.idValue ?? ""
        return message
    },
}

function createBaseMsgRejectGameResponse(): MsgRejectGameResponse {
    return {}
}

export const MsgRejectGameResponse = {
    encode(_: MsgRejectGameResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
        return writer
    },

    decode(input: _m0.Reader | Uint8Array, length?: number): MsgRejectGameResponse {
        const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input)
        let end = length === undefined ? reader.len : reader.pos + length
        const message = createBaseMsgRejectGameResponse()
        while (reader.pos < end) {
            const tag = reader.uint32()
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7)
                    break
            }
        }
        return message
    },

    fromJSON(_: any): MsgRejectGameResponse {
        return {}
    },

    toJSON(_: MsgRejectGameResponse): unknown {
        const obj: any = {}
        return obj
    },

    fromPartial<I extends Exact<DeepPartial<MsgRejectGameResponse>, I>>(_: I): MsgRejectGameResponse {
        const message = createBaseMsgRejectGameResponse()
        return message
    },
}

function createBaseMsgPlayMove(): MsgPlayMove {
    return {
        creator: "",
        idValue: "",
        fromX: Long.UZERO,
        fromY: Long.UZERO,
        toX: Long.UZERO,
        toY: Long.UZERO,
    }
}

export const MsgPlayMove = {
    encode(message: MsgPlayMove, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator)
        }
        if (message.idValue !== "") {
            writer.uint32(18).string(message.idValue)
        }
        if (!message.fromX.isZero()) {
            writer.uint32(24).uint64(message.fromX)
        }
        if (!message.fromY.isZero()) {
            writer.uint32(32).uint64(message.fromY)
        }
        if (!message.toX.isZero()) {
            writer.uint32(40).uint64(message.toX)
        }
        if (!message.toY.isZero()) {
            writer.uint32(48).uint64(message.toY)
        }
        return writer
    },

    decode(input: _m0.Reader | Uint8Array, length?: number): MsgPlayMove {
        const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input)
        let end = length === undefined ? reader.len : reader.pos + length
        const message = createBaseMsgPlayMove()
        while (reader.pos < end) {
            const tag = reader.uint32()
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string()
                    break
                case 2:
                    message.idValue = reader.string()
                    break
                case 3:
                    message.fromX = reader.uint64() as Long
                    break
                case 4:
                    message.fromY = reader.uint64() as Long
                    break
                case 5:
                    message.toX = reader.uint64() as Long
                    break
                case 6:
                    message.toY = reader.uint64() as Long
                    break
                default:
                    reader.skipType(tag & 7)
                    break
            }
        }
        return message
    },

    fromJSON(object: any): MsgPlayMove {
        return {
            creator: isSet(object.creator) ? String(object.creator) : "",
            idValue: isSet(object.idValue) ? String(object.idValue) : "",
            fromX: isSet(object.fromX) ? Long.fromString(object.fromX) : Long.UZERO,
            fromY: isSet(object.fromY) ? Long.fromString(object.fromY) : Long.UZERO,
            toX: isSet(object.toX) ? Long.fromString(object.toX) : Long.UZERO,
            toY: isSet(object.toY) ? Long.fromString(object.toY) : Long.UZERO,
        }
    },

    toJSON(message: MsgPlayMove): unknown {
        const obj: any = {}
        message.creator !== undefined && (obj.creator = message.creator)
        message.idValue !== undefined && (obj.idValue = message.idValue)
        message.fromX !== undefined && (obj.fromX = (message.fromX || Long.UZERO).toString())
        message.fromY !== undefined && (obj.fromY = (message.fromY || Long.UZERO).toString())
        message.toX !== undefined && (obj.toX = (message.toX || Long.UZERO).toString())
        message.toY !== undefined && (obj.toY = (message.toY || Long.UZERO).toString())
        return obj
    },

    fromPartial<I extends Exact<DeepPartial<MsgPlayMove>, I>>(object: I): MsgPlayMove {
        const message = createBaseMsgPlayMove()
        message.creator = object.creator ?? ""
        message.idValue = object.idValue ?? ""
        message.fromX =
            object.fromX !== undefined && object.fromX !== null ? Long.fromValue(object.fromX) : Long.UZERO
        message.fromY =
            object.fromY !== undefined && object.fromY !== null ? Long.fromValue(object.fromY) : Long.UZERO
        message.toX =
            object.toX !== undefined && object.toX !== null ? Long.fromValue(object.toX) : Long.UZERO
        message.toY =
            object.toY !== undefined && object.toY !== null ? Long.fromValue(object.toY) : Long.UZERO
        return message
    },
}

function createBaseMsgPlayMoveResponse(): MsgPlayMoveResponse {
    return { idValue: "", capturedX: Long.ZERO, capturedY: Long.ZERO, winner: "" }
}

export const MsgPlayMoveResponse = {
    encode(message: MsgPlayMoveResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
        if (message.idValue !== "") {
            writer.uint32(10).string(message.idValue)
        }
        if (!message.capturedX.isZero()) {
            writer.uint32(16).int64(message.capturedX)
        }
        if (!message.capturedY.isZero()) {
            writer.uint32(24).int64(message.capturedY)
        }
        if (message.winner !== "") {
            writer.uint32(34).string(message.winner)
        }
        return writer
    },

    decode(input: _m0.Reader | Uint8Array, length?: number): MsgPlayMoveResponse {
        const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input)
        let end = length === undefined ? reader.len : reader.pos + length
        const message = createBaseMsgPlayMoveResponse()
        while (reader.pos < end) {
            const tag = reader.uint32()
            switch (tag >>> 3) {
                case 1:
                    message.idValue = reader.string()
                    break
                case 2:
                    message.capturedX = reader.int64() as Long
                    break
                case 3:
                    message.capturedY = reader.int64() as Long
                    break
                case 4:
                    message.winner = reader.string()
                    break
                default:
                    reader.skipType(tag & 7)
                    break
            }
        }
        return message
    },

    fromJSON(object: any): MsgPlayMoveResponse {
        return {
            idValue: isSet(object.idValue) ? String(object.idValue) : "",
            capturedX: isSet(object.capturedX) ? Long.fromString(object.capturedX) : Long.ZERO,
            capturedY: isSet(object.capturedY) ? Long.fromString(object.capturedY) : Long.ZERO,
            winner: isSet(object.winner) ? String(object.winner) : "",
        }
    },

    toJSON(message: MsgPlayMoveResponse): unknown {
        const obj: any = {}
        message.idValue !== undefined && (obj.idValue = message.idValue)
        message.capturedX !== undefined && (obj.capturedX = (message.capturedX || Long.ZERO).toString())
        message.capturedY !== undefined && (obj.capturedY = (message.capturedY || Long.ZERO).toString())
        message.winner !== undefined && (obj.winner = message.winner)
        return obj
    },

    fromPartial<I extends Exact<DeepPartial<MsgPlayMoveResponse>, I>>(object: I): MsgPlayMoveResponse {
        const message = createBaseMsgPlayMoveResponse()
        message.idValue = object.idValue ?? ""
        message.capturedX =
            object.capturedX !== undefined && object.capturedX !== null
                ? Long.fromValue(object.capturedX)
                : Long.ZERO
        message.capturedY =
            object.capturedY !== undefined && object.capturedY !== null
                ? Long.fromValue(object.capturedY)
                : Long.ZERO
        message.winner = object.winner ?? ""
        return message
    },
}

function createBaseMsgCreateGame(): MsgCreateGame {
    return { creator: "", red: "", black: "", wager: Long.UZERO, token: "" }
}

export const MsgCreateGame = {
    encode(message: MsgCreateGame, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator)
        }
        if (message.red !== "") {
            writer.uint32(18).string(message.red)
        }
        if (message.black !== "") {
            writer.uint32(26).string(message.black)
        }
        if (!message.wager.isZero()) {
            writer.uint32(32).uint64(message.wager)
        }
        if (message.token !== "") {
            writer.uint32(42).string(message.token)
        }
        return writer
    },

    decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateGame {
        const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input)
        let end = length === undefined ? reader.len : reader.pos + length
        const message = createBaseMsgCreateGame()
        while (reader.pos < end) {
            const tag = reader.uint32()
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string()
                    break
                case 2:
                    message.red = reader.string()
                    break
                case 3:
                    message.black = reader.string()
                    break
                case 4:
                    message.wager = reader.uint64() as Long
                    break
                case 5:
                    message.token = reader.string()
                    break
                default:
                    reader.skipType(tag & 7)
                    break
            }
        }
        return message
    },

    fromJSON(object: any): MsgCreateGame {
        return {
            creator: isSet(object.creator) ? String(object.creator) : "",
            red: isSet(object.red) ? String(object.red) : "",
            black: isSet(object.black) ? String(object.black) : "",
            wager: isSet(object.wager) ? Long.fromString(object.wager) : Long.UZERO,
            token: isSet(object.token) ? String(object.token) : "",
        }
    },

    toJSON(message: MsgCreateGame): unknown {
        const obj: any = {}
        message.creator !== undefined && (obj.creator = message.creator)
        message.red !== undefined && (obj.red = message.red)
        message.black !== undefined && (obj.black = message.black)
        message.wager !== undefined && (obj.wager = (message.wager || Long.UZERO).toString())
        message.token !== undefined && (obj.token = message.token)
        return obj
    },

    fromPartial<I extends Exact<DeepPartial<MsgCreateGame>, I>>(object: I): MsgCreateGame {
        const message = createBaseMsgCreateGame()
        message.creator = object.creator ?? ""
        message.red = object.red ?? ""
        message.black = object.black ?? ""
        message.wager =
            object.wager !== undefined && object.wager !== null ? Long.fromValue(object.wager) : Long.UZERO
        message.token = object.token ?? ""
        return message
    },
}

function createBaseMsgCreateGameResponse(): MsgCreateGameResponse {
    return { idValue: "" }
}

export const MsgCreateGameResponse = {
    encode(message: MsgCreateGameResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
        if (message.idValue !== "") {
            writer.uint32(10).string(message.idValue)
        }
        return writer
    },

    decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateGameResponse {
        const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input)
        let end = length === undefined ? reader.len : reader.pos + length
        const message = createBaseMsgCreateGameResponse()
        while (reader.pos < end) {
            const tag = reader.uint32()
            switch (tag >>> 3) {
                case 1:
                    message.idValue = reader.string()
                    break
                default:
                    reader.skipType(tag & 7)
                    break
            }
        }
        return message
    },

    fromJSON(object: any): MsgCreateGameResponse {
        return {
            idValue: isSet(object.idValue) ? String(object.idValue) : "",
        }
    },

    toJSON(message: MsgCreateGameResponse): unknown {
        const obj: any = {}
        message.idValue !== undefined && (obj.idValue = message.idValue)
        return obj
    },

    fromPartial<I extends Exact<DeepPartial<MsgCreateGameResponse>, I>>(object: I): MsgCreateGameResponse {
        const message = createBaseMsgCreateGameResponse()
        message.idValue = object.idValue ?? ""
        return message
    },
}

/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    RejectGame(request: MsgRejectGame): Promise<MsgRejectGameResponse>
    PlayMove(request: MsgPlayMove): Promise<MsgPlayMoveResponse>
    CreateGame(request: MsgCreateGame): Promise<MsgCreateGameResponse>
}

export class MsgClientImpl implements Msg {
    private readonly rpc: Rpc
    constructor(rpc: Rpc) {
        this.rpc = rpc
        this.RejectGame = this.RejectGame.bind(this)
        this.PlayMove = this.PlayMove.bind(this)
        this.CreateGame = this.CreateGame.bind(this)
    }
    RejectGame(request: MsgRejectGame): Promise<MsgRejectGameResponse> {
        const data = MsgRejectGame.encode(request).finish()
        const promise = this.rpc.request("xavierlepretre.checkers.checkers.Msg", "RejectGame", data)
        return promise.then((data) => MsgRejectGameResponse.decode(new _m0.Reader(data)))
    }

    PlayMove(request: MsgPlayMove): Promise<MsgPlayMoveResponse> {
        const data = MsgPlayMove.encode(request).finish()
        const promise = this.rpc.request("xavierlepretre.checkers.checkers.Msg", "PlayMove", data)
        return promise.then((data) => MsgPlayMoveResponse.decode(new _m0.Reader(data)))
    }

    CreateGame(request: MsgCreateGame): Promise<MsgCreateGameResponse> {
        const data = MsgCreateGame.encode(request).finish()
        const promise = this.rpc.request("xavierlepretre.checkers.checkers.Msg", "CreateGame", data)
        return promise.then((data) => MsgCreateGameResponse.decode(new _m0.Reader(data)))
    }
}

interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
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
