/* eslint-disable */
import Long from "long"
import _m0 from "protobufjs/minimal"
import { StoredGame } from "../checkers/stored_game"
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination"
import { NextGame } from "../checkers/next_game"

export const protobufPackage = "xavierlepretre.checkers.checkers"

/** this line is used by starport scaffolding # 3 */
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

    fromPartial<I extends Exact<DeepPartial<QueryGetStoredGameRequest>, I>>(object: I): QueryGetStoredGameRequest {
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
        message.StoredGame !== undefined && (obj.StoredGame = message.StoredGame ? StoredGame.toJSON(message.StoredGame) : undefined)
        return obj
    },

    fromPartial<I extends Exact<DeepPartial<QueryGetStoredGameResponse>, I>>(object: I): QueryGetStoredGameResponse {
        const message = createBaseQueryGetStoredGameResponse()
        message.StoredGame = object.StoredGame !== undefined && object.StoredGame !== null ? StoredGame.fromPartial(object.StoredGame) : undefined
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
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
        return obj
    },

    fromPartial<I extends Exact<DeepPartial<QueryAllStoredGameRequest>, I>>(object: I): QueryAllStoredGameRequest {
        const message = createBaseQueryAllStoredGameRequest()
        message.pagination = object.pagination !== undefined && object.pagination !== null ? PageRequest.fromPartial(object.pagination) : undefined
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
            StoredGame: Array.isArray(object?.StoredGame) ? object.StoredGame.map((e: any) => StoredGame.fromJSON(e)) : [],
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
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
        return obj
    },

    fromPartial<I extends Exact<DeepPartial<QueryAllStoredGameResponse>, I>>(object: I): QueryAllStoredGameResponse {
        const message = createBaseQueryAllStoredGameResponse()
        message.StoredGame = object.StoredGame?.map((e) => StoredGame.fromPartial(e)) || []
        message.pagination = object.pagination !== undefined && object.pagination !== null ? PageResponse.fromPartial(object.pagination) : undefined
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
        message.NextGame !== undefined && (obj.NextGame = message.NextGame ? NextGame.toJSON(message.NextGame) : undefined)
        return obj
    },

    fromPartial<I extends Exact<DeepPartial<QueryGetNextGameResponse>, I>>(object: I): QueryGetNextGameResponse {
        const message = createBaseQueryGetNextGameResponse()
        message.NextGame = object.NextGame !== undefined && object.NextGame !== null ? NextGame.fromPartial(object.NextGame) : undefined
        return message
    },
}

/** Query defines the gRPC querier service. */
export interface Query {
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
        this.StoredGame = this.StoredGame.bind(this)
        this.StoredGameAll = this.StoredGameAll.bind(this)
        this.NextGame = this.NextGame.bind(this)
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
export type Exact<P, I extends P> = P extends Builtin ? P : P & { [K in keyof P]: Exact<P[K], I[K]> } & Record<Exclude<keyof I, KeysOfUnion<P>>, never>

if (_m0.util.Long !== Long) {
    _m0.util.Long = Long as any
    _m0.configure()
}

function isSet(value: any): boolean {
    return value !== null && value !== undefined
}
