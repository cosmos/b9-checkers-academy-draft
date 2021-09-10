import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "xavierlepretre.checkers.checkers";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreateGame {
    creator: string;
    red: string;
    black: string;
}
export interface MsgCreateGameResponse {
    idValue: string;
}
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
    CreateGame(request: MsgCreateGame): Promise<MsgCreateGameResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
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
