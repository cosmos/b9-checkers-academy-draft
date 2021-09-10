import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "xavierlepretre.checkers.checkers";
/** this line is used by starport scaffolding # proto/tx/message */
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
}
export interface MsgCreateGame {
    creator: string;
    red: string;
    black: string;
}
export interface MsgCreateGameResponse {
    idValue: string;
}
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
    PlayMove(request: MsgPlayMove): Promise<MsgPlayMoveResponse>;
    CreateGame(request: MsgCreateGame): Promise<MsgCreateGameResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
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
