import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "b9lab.checkers.checkers";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgRejectGame {
    creator: string;
    idValue: string;
}
export interface MsgRejectGameResponse {
}
export interface MsgPlayMove {
    creator: string;
    idValue: string;
    fromX: number;
    fromY: number;
    toX: number;
    toY: number;
}
export interface MsgPlayMoveResponse {
    idValue: string;
    capturedX: number;
    capturedY: number;
    winner: string;
}
export interface MsgCreateGame {
    creator: string;
    red: string;
    black: string;
}
export interface MsgCreateGameResponse {
    idValue: string;
}
export declare const MsgRejectGame: {
    encode(message: MsgRejectGame, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRejectGame;
    fromJSON(object: any): MsgRejectGame;
    toJSON(message: MsgRejectGame): unknown;
    fromPartial(object: DeepPartial<MsgRejectGame>): MsgRejectGame;
};
export declare const MsgRejectGameResponse: {
    encode(_: MsgRejectGameResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRejectGameResponse;
    fromJSON(_: any): MsgRejectGameResponse;
    toJSON(_: MsgRejectGameResponse): unknown;
    fromPartial(_: DeepPartial<MsgRejectGameResponse>): MsgRejectGameResponse;
};
export declare const MsgPlayMove: {
    encode(message: MsgPlayMove, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgPlayMove;
    fromJSON(object: any): MsgPlayMove;
    toJSON(message: MsgPlayMove): unknown;
    fromPartial(object: DeepPartial<MsgPlayMove>): MsgPlayMove;
};
export declare const MsgPlayMoveResponse: {
    encode(message: MsgPlayMoveResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgPlayMoveResponse;
    fromJSON(object: any): MsgPlayMoveResponse;
    toJSON(message: MsgPlayMoveResponse): unknown;
    fromPartial(object: DeepPartial<MsgPlayMoveResponse>): MsgPlayMoveResponse;
};
export declare const MsgCreateGame: {
    encode(message: MsgCreateGame, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateGame;
    fromJSON(object: any): MsgCreateGame;
    toJSON(message: MsgCreateGame): unknown;
    fromPartial(object: DeepPartial<MsgCreateGame>): MsgCreateGame;
};
export declare const MsgCreateGameResponse: {
    encode(message: MsgCreateGameResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateGameResponse;
    fromJSON(object: any): MsgCreateGameResponse;
    toJSON(message: MsgCreateGameResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreateGameResponse>): MsgCreateGameResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    RejectGame(request: MsgRejectGame): Promise<MsgRejectGameResponse>;
    PlayMove(request: MsgPlayMove): Promise<MsgPlayMoveResponse>;
    CreateGame(request: MsgCreateGame): Promise<MsgCreateGameResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    RejectGame(request: MsgRejectGame): Promise<MsgRejectGameResponse>;
    PlayMove(request: MsgPlayMove): Promise<MsgPlayMoveResponse>;
    CreateGame(request: MsgCreateGame): Promise<MsgCreateGameResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
