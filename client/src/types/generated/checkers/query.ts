/* eslint-disable */
import Long from "long"
import _m0 from "protobufjs/minimal"
import { StoredGame } from "../checkers/stored_game"
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination"
import { NextGame } from "../checkers/next_game"

export const protobufPackage = "xavierlepretre.checkers.checkers"

/** this line is used by starport scaffolding # 3 */
export interface QueryCanPlayMoveRequest {
    idValue: string
    player: string
    fromX: Long
    fromY: Long
    toX: Long
    toY: Long
}

export interface QueryCanPlayMoveResponse {
    possible: boolean
    reason: string
}

export interface QueryGetStoredGameRequest {
    index: string
}

export interface QueryGetStoredGameResponse {
    StoredGame?: StoredGame
}

export interface QueryAllStoredGameRequest {
    pagination?: PageRequest
}

export interface QueryAllStoredGameResponse {
    StoredGame: StoredGame[]
    pagination?: PageResponse
}

export interface QueryGetNextGameRequest {}

export interface QueryGetNextGameResponse {
    NextGame?: NextGame
}

function createBaseQueryCanPlayMoveRequest(): QueryCanPlayMoveRequest {
    return { idValue: "", player: "", fromX: Long.UZERO, fromY: Long.UZERO, toX: Long.UZERO, toY: Long.UZERO }
}

export const QueryCanPlayMoveRequest = {
    encode(message: QueryCanPlayMoveRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
        if (message.idValue !== "") {
            writer.uint32(10).string(message.idValue)
        }
        if (message.player !== "") {
            writer.uint32(18).string(message.player)
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

    decode(input: _m0.Reader | Uint8Array, length?: number): QueryCanPlayMoveRequest {
        const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input)
        let end = length === undefined ? reader.len : reader.pos + length
        const message = createBaseQueryCanPlayMoveRequest()
        while (reader.pos < end) {
            const tag = reader.uint32()
            switch (tag >>> 3) {
                case 1:
                    message.idValue = reader.string()
                    break
                case 2:
                    message.player = reader.string()
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

    fromJSON(object: any): QueryCanPlayMoveRequest {
        return {
            idValue: isSet(object.idValue) ? String(object.idValue) : "",
            player: isSet(object.player) ? String(object.player) : "",
            fromX: isSet(object.fromX) ? Long.fromString(object.fromX) : Long.UZERO,
            fromY: isSet(object.fromY) ? Long.fromString(object.fromY) : Long.UZERO,
            toX: isSet(object.toX) ? Long.fromString(object.toX) : Long.UZERO,
            toY: isSet(object.toY) ? Long.fromString(object.toY) : Long.UZERO,
        }
    },

    toJSON(message: QueryCanPlayMoveRequest): unknown {
        const obj: any = {}
        message.idValue !== undefined && (obj.idValue = message.idValue)
        message.player !== undefined && (obj.player = message.player)
        message.fromX !== undefined && (obj.fromX = (message.fromX || Long.UZERO).toString())
        message.fromY !== undefined && (obj.fromY = (message.fromY || Long.UZERO).toString())
        message.toX !== undefined && (obj.toX = (message.toX || Long.UZERO).toString())
        message.toY !== undefined && (obj.toY = (message.toY || Long.UZERO).toString())
        return obj
    },

    fromPartial<I extends Exact<DeepPartial<QueryCanPlayMoveRequest>, I>>(
        object: I,
    ): QueryCanPlayMoveRequest {
        const message = createBaseQueryCanPlayMoveRequest()
        message.idValue = object.idValue ?? ""
        message.player = object.player ?? ""
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

function createBaseQueryCanPlayMoveResponse(): QueryCanPlayMoveResponse {
    return { possible: false, reason: "" }
}

export const QueryCanPlayMoveResponse = {
    encode(message: QueryCanPlayMoveResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
        if (message.possible === true) {
            writer.uint32(8).bool(message.possible)
        }
        if (message.reason !== "") {
            writer.uint32(18).string(message.reason)
        }
        return writer
    },

    decode(input: _m0.Reader | Uint8Array, length?: number): QueryCanPlayMoveResponse {
        const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input)
        let end = length === undefined ? reader.len : reader.pos + length
        const message = createBaseQueryCanPlayMoveResponse()
        while (reader.pos < end) {
            const tag = reader.uint32()
            switch (tag >>> 3) {
                case 1:
                    message.possible = reader.bool()
                    break
                case 2:
                    message.reason = reader.string()
                    break
                default:
                    reader.skipType(tag & 7)
                    break
            }
        }
        return message
    },

    fromJSON(object: any): QueryCanPlayMoveResponse {
        return {
            possible: isSet(object.possible) ? Boolean(object.possible) : false,
            reason: isSet(object.reason) ? String(object.reason) : "",
        }
    },

    toJSON(message: QueryCanPlayMoveResponse): unknown {
        const obj: any = {}
        message.possible !== undefined && (obj.possible = message.possible)
        message.reason !== undefined && (obj.reason = message.reason)
        return obj
    },

    fromPartial<I extends Exact<DeepPartial<QueryCanPlayMoveResponse>, I>>(
        object: I,
    ): QueryCanPlayMoveResponse {
        const message = createBaseQueryCanPlayMoveResponse()
        message.possible = object.possible ?? false
        message.reason = object.reason ?? ""
        return message
    },
}

function createBaseQueryGetStoredGameRequest(): QueryGetStoredGameRequest {
    return { index: "" }
}

export const QueryGetStoredGameRequest = {
    encode(message: QueryGetStoredGameRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
        if (message.index !== "") {
            writer.uint32(10).string(message.index)
        }
        return writer
    },

    decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetStoredGameRequest {
        const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input)
        let end = length === undefined ? reader.len : reader.pos + length
        const message = createBaseQueryGetStoredGameRequest()
        while (reader.pos < end) {
            const tag = reader.uint32()
            switch (tag >>> 3) {
                case 1:
                    message.index = reader.string()
                    break
                default:
                    reader.skipType(tag & 7)
                    break
            }
        }
        return message
    },

    fromJSON(object: any): QueryGetStoredGameRequest {
        return {
            index: isSet(object.index) ? String(object.index) : "",
        }
    },

    toJSON(message: QueryGetStoredGameRequest): unknown {
        const obj: any = {}
        message.index !== undefined && (obj.index = message.index)
        return obj
    },

    fromPartial<I extends Exact<DeepPartial<QueryGetStoredGameRequest>, I>>(
        object: I,
    ): QueryGetStoredGameRequest {
        const message = createBaseQueryGetStoredGameRequest()
        message.index = object.index ?? ""
        return message
    },
}

function createBaseQueryGetStoredGameResponse(): QueryGetStoredGameResponse {
    return { StoredGame: undefined }
}

export const QueryGetStoredGameResponse = {
    encode(message: QueryGetStoredGameResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
        if (message.StoredGame !== undefined) {
            StoredGame.encode(message.StoredGame, writer.uint32(10).fork()).ldelim()
        }
        return writer
    },

    decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetStoredGameResponse {
        const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input)
        let end = length === undefined ? reader.len : reader.pos + length
        const message = createBaseQueryGetStoredGameResponse()
        while (reader.pos < end) {
            const tag = reader.uint32()
            switch (tag >>> 3) {
                case 1:
                    message.StoredGame = StoredGame.decode(reader, reader.uint32())
                    break
                default:
                    reader.skipType(tag & 7)
                    break
            }
        }
        return message
    },

    fromJSON(object: any): QueryGetStoredGameResponse {
        return {
            StoredGame: isSet(object.StoredGame) ? StoredGame.fromJSON(object.StoredGame) : undefined,
        }
    },

    toJSON(message: QueryGetStoredGameResponse): unknown {
        const obj: any = {}
        message.StoredGame !== undefined &&
            (obj.StoredGame = message.StoredGame ? StoredGame.toJSON(message.StoredGame) : undefined)
        return obj
    },

    fromPartial<I extends Exact<DeepPartial<QueryGetStoredGameResponse>, I>>(
        object: I,
    ): QueryGetStoredGameResponse {
        const message = createBaseQueryGetStoredGameResponse()
        message.StoredGame =
            object.StoredGame !== undefined && object.StoredGame !== null
                ? StoredGame.fromPartial(object.StoredGame)
                : undefined
        return message
    },
}

function createBaseQueryAllStoredGameRequest(): QueryAllStoredGameRequest {
    return { pagination: undefined }
}

export const QueryAllStoredGameRequest = {
    encode(message: QueryAllStoredGameRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
        }
        return writer
    },

    decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllStoredGameRequest {
        const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input)
        let end = length === undefined ? reader.len : reader.pos + length
        const message = createBaseQueryAllStoredGameRequest()
        while (reader.pos < end) {
            const tag = reader.uint32()
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32())
                    break
                default:
                    reader.skipType(tag & 7)
                    break
            }
        }
        return message
    },

    fromJSON(object: any): QueryAllStoredGameRequest {
        return {
            pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined,
        }
    },

    toJSON(message: QueryAllStoredGameRequest): unknown {
        const obj: any = {}
        message.pagination !== undefined &&
            (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
        return obj
    },

    fromPartial<I extends Exact<DeepPartial<QueryAllStoredGameRequest>, I>>(
        object: I,
    ): QueryAllStoredGameRequest {
        const message = createBaseQueryAllStoredGameRequest()
        message.pagination =
            object.pagination !== undefined && object.pagination !== null
                ? PageRequest.fromPartial(object.pagination)
                : undefined
        return message
    },
}

function createBaseQueryAllStoredGameResponse(): QueryAllStoredGameResponse {
    return { StoredGame: [], pagination: undefined }
}

export const QueryAllStoredGameResponse = {
    encode(message: QueryAllStoredGameResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
        for (const v of message.StoredGame) {
            StoredGame.encode(v!, writer.uint32(10).fork()).ldelim()
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
        }
        return writer
    },

    decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllStoredGameResponse {
        const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input)
        let end = length === undefined ? reader.len : reader.pos + length
        const message = createBaseQueryAllStoredGameResponse()
        while (reader.pos < end) {
            const tag = reader.uint32()
            switch (tag >>> 3) {
                case 1:
                    message.StoredGame.push(StoredGame.decode(reader, reader.uint32()))
                    break
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32())
                    break
                default:
                    reader.skipType(tag & 7)
                    break
            }
        }
        return message
    },

    fromJSON(object: any): QueryAllStoredGameResponse {
        return {
            StoredGame: Array.isArray(object?.StoredGame)
                ? object.StoredGame.map((e: any) => StoredGame.fromJSON(e))
                : [],
            pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
        }
    },

    toJSON(message: QueryAllStoredGameResponse): unknown {
        const obj: any = {}
        if (message.StoredGame) {
            obj.StoredGame = message.StoredGame.map((e) => (e ? StoredGame.toJSON(e) : undefined))
        } else {
            obj.StoredGame = []
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
        return obj
    },

    fromPartial<I extends Exact<DeepPartial<QueryAllStoredGameResponse>, I>>(
        object: I,
    ): QueryAllStoredGameResponse {
        const message = createBaseQueryAllStoredGameResponse()
        message.StoredGame = object.StoredGame?.map((e) => StoredGame.fromPartial(e)) || []
        message.pagination =
            object.pagination !== undefined && object.pagination !== null
                ? PageResponse.fromPartial(object.pagination)
                : undefined
        return message
    },
}

function createBaseQueryGetNextGameRequest(): QueryGetNextGameRequest {
    return {}
}

export const QueryGetNextGameRequest = {
    encode(_: QueryGetNextGameRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
        return writer
    },

    decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetNextGameRequest {
        const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input)
        let end = length === undefined ? reader.len : reader.pos + length
        const message = createBaseQueryGetNextGameRequest()
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

    fromJSON(_: any): QueryGetNextGameRequest {
        return {}
    },

    toJSON(_: QueryGetNextGameRequest): unknown {
        const obj: any = {}
        return obj
    },

    fromPartial<I extends Exact<DeepPartial<QueryGetNextGameRequest>, I>>(_: I): QueryGetNextGameRequest {
        const message = createBaseQueryGetNextGameRequest()
        return message
    },
}

function createBaseQueryGetNextGameResponse(): QueryGetNextGameResponse {
    return { NextGame: undefined }
}

export const QueryGetNextGameResponse = {
    encode(message: QueryGetNextGameResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
        if (message.NextGame !== undefined) {
            NextGame.encode(message.NextGame, writer.uint32(10).fork()).ldelim()
        }
        return writer
    },

    decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetNextGameResponse {
        const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input)
        let end = length === undefined ? reader.len : reader.pos + length
        const message = createBaseQueryGetNextGameResponse()
        while (reader.pos < end) {
            const tag = reader.uint32()
            switch (tag >>> 3) {
                case 1:
                    message.NextGame = NextGame.decode(reader, reader.uint32())
                    break
                default:
                    reader.skipType(tag & 7)
                    break
            }
        }
        return message
    },

    fromJSON(object: any): QueryGetNextGameResponse {
        return {
            NextGame: isSet(object.NextGame) ? NextGame.fromJSON(object.NextGame) : undefined,
        }
    },

    toJSON(message: QueryGetNextGameResponse): unknown {
        const obj: any = {}
        message.NextGame !== undefined &&
            (obj.NextGame = message.NextGame ? NextGame.toJSON(message.NextGame) : undefined)
        return obj
    },

    fromPartial<I extends Exact<DeepPartial<QueryGetNextGameResponse>, I>>(
        object: I,
    ): QueryGetNextGameResponse {
        const message = createBaseQueryGetNextGameResponse()
        message.NextGame =
            object.NextGame !== undefined && object.NextGame !== null
                ? NextGame.fromPartial(object.NextGame)
                : undefined
        return message
    },
}

/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a list of canPlayMove items. */
    CanPlayMove(request: QueryCanPlayMoveRequest): Promise<QueryCanPlayMoveResponse>
    /** Queries a storedGame by index. */
    StoredGame(request: QueryGetStoredGameRequest): Promise<QueryGetStoredGameResponse>
    /** Queries a list of storedGame items. */
    StoredGameAll(request: QueryAllStoredGameRequest): Promise<QueryAllStoredGameResponse>
    /** Queries a nextGame by index. */
    NextGame(request: QueryGetNextGameRequest): Promise<QueryGetNextGameResponse>
}

export class QueryClientImpl implements Query {
    private readonly rpc: Rpc
    constructor(rpc: Rpc) {
        this.rpc = rpc
        this.CanPlayMove = this.CanPlayMove.bind(this)
        this.StoredGame = this.StoredGame.bind(this)
        this.StoredGameAll = this.StoredGameAll.bind(this)
        this.NextGame = this.NextGame.bind(this)
    }
    CanPlayMove(request: QueryCanPlayMoveRequest): Promise<QueryCanPlayMoveResponse> {
        const data = QueryCanPlayMoveRequest.encode(request).finish()
        const promise = this.rpc.request("xavierlepretre.checkers.checkers.Query", "CanPlayMove", data)
        return promise.then((data) => QueryCanPlayMoveResponse.decode(new _m0.Reader(data)))
    }

    StoredGame(request: QueryGetStoredGameRequest): Promise<QueryGetStoredGameResponse> {
        const data = QueryGetStoredGameRequest.encode(request).finish()
        const promise = this.rpc.request("xavierlepretre.checkers.checkers.Query", "StoredGame", data)
        return promise.then((data) => QueryGetStoredGameResponse.decode(new _m0.Reader(data)))
    }

    StoredGameAll(request: QueryAllStoredGameRequest): Promise<QueryAllStoredGameResponse> {
        const data = QueryAllStoredGameRequest.encode(request).finish()
        const promise = this.rpc.request("xavierlepretre.checkers.checkers.Query", "StoredGameAll", data)
        return promise.then((data) => QueryAllStoredGameResponse.decode(new _m0.Reader(data)))
    }

    NextGame(request: QueryGetNextGameRequest): Promise<QueryGetNextGameResponse> {
        const data = QueryGetNextGameRequest.encode(request).finish()
        const promise = this.rpc.request("xavierlepretre.checkers.checkers.Query", "NextGame", data)
        return promise.then((data) => QueryGetNextGameResponse.decode(new _m0.Reader(data)))
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
