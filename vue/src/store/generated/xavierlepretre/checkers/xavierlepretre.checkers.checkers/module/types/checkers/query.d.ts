import { Reader, Writer } from 'protobufjs/minimal';
import { StoredGame } from '../checkers/stored_game';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { NextGame } from '../checkers/next_game';
export declare const protobufPackage = "xavierlepretre.checkers.checkers";
/** this line is used by starport scaffolding # 3 */
export interface QueryCanPlayMoveRequest {
    idValue: string;
    player: string;
    fromX: number;
    fromY: number;
    toX: number;
    toY: number;
}
export interface QueryCanPlayMoveResponse {
    possible: boolean;
}
export interface QueryGetStoredGameRequest {
    index: string;
}
export interface QueryGetStoredGameResponse {
    StoredGame: StoredGame | undefined;
}
export interface QueryAllStoredGameRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllStoredGameResponse {
    StoredGame: StoredGame[];
    pagination: PageResponse | undefined;
}
export interface QueryGetNextGameRequest {
}
export interface QueryGetNextGameResponse {
    NextGame: NextGame | undefined;
}
export declare const QueryCanPlayMoveRequest: {
    encode(message: QueryCanPlayMoveRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryCanPlayMoveRequest;
    fromJSON(object: any): QueryCanPlayMoveRequest;
    toJSON(message: QueryCanPlayMoveRequest): unknown;
    fromPartial(object: DeepPartial<QueryCanPlayMoveRequest>): QueryCanPlayMoveRequest;
};
export declare const QueryCanPlayMoveResponse: {
    encode(message: QueryCanPlayMoveResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryCanPlayMoveResponse;
    fromJSON(object: any): QueryCanPlayMoveResponse;
    toJSON(message: QueryCanPlayMoveResponse): unknown;
    fromPartial(object: DeepPartial<QueryCanPlayMoveResponse>): QueryCanPlayMoveResponse;
};
export declare const QueryGetStoredGameRequest: {
    encode(message: QueryGetStoredGameRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetStoredGameRequest;
    fromJSON(object: any): QueryGetStoredGameRequest;
    toJSON(message: QueryGetStoredGameRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetStoredGameRequest>): QueryGetStoredGameRequest;
};
export declare const QueryGetStoredGameResponse: {
    encode(message: QueryGetStoredGameResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetStoredGameResponse;
    fromJSON(object: any): QueryGetStoredGameResponse;
    toJSON(message: QueryGetStoredGameResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetStoredGameResponse>): QueryGetStoredGameResponse;
};
export declare const QueryAllStoredGameRequest: {
    encode(message: QueryAllStoredGameRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllStoredGameRequest;
    fromJSON(object: any): QueryAllStoredGameRequest;
    toJSON(message: QueryAllStoredGameRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllStoredGameRequest>): QueryAllStoredGameRequest;
};
export declare const QueryAllStoredGameResponse: {
    encode(message: QueryAllStoredGameResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllStoredGameResponse;
    fromJSON(object: any): QueryAllStoredGameResponse;
    toJSON(message: QueryAllStoredGameResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllStoredGameResponse>): QueryAllStoredGameResponse;
};
export declare const QueryGetNextGameRequest: {
    encode(_: QueryGetNextGameRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetNextGameRequest;
    fromJSON(_: any): QueryGetNextGameRequest;
    toJSON(_: QueryGetNextGameRequest): unknown;
    fromPartial(_: DeepPartial<QueryGetNextGameRequest>): QueryGetNextGameRequest;
};
export declare const QueryGetNextGameResponse: {
    encode(message: QueryGetNextGameResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetNextGameResponse;
    fromJSON(object: any): QueryGetNextGameResponse;
    toJSON(message: QueryGetNextGameResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetNextGameResponse>): QueryGetNextGameResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a list of canPlayMove items. */
    CanPlayMove(request: QueryCanPlayMoveRequest): Promise<QueryCanPlayMoveResponse>;
    /** Queries a storedGame by index. */
    StoredGame(request: QueryGetStoredGameRequest): Promise<QueryGetStoredGameResponse>;
    /** Queries a list of storedGame items. */
    StoredGameAll(request: QueryAllStoredGameRequest): Promise<QueryAllStoredGameResponse>;
    /** Queries a nextGame by index. */
    NextGame(request: QueryGetNextGameRequest): Promise<QueryGetNextGameResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    CanPlayMove(request: QueryCanPlayMoveRequest): Promise<QueryCanPlayMoveResponse>;
    StoredGame(request: QueryGetStoredGameRequest): Promise<QueryGetStoredGameResponse>;
    StoredGameAll(request: QueryAllStoredGameRequest): Promise<QueryAllStoredGameResponse>;
    NextGame(request: QueryGetNextGameRequest): Promise<QueryGetNextGameResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
