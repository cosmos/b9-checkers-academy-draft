import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "xavierlepretre.checkers.checkers";
export interface WinningPlayer {
    playerAddress: string;
    wonCount: number;
    dateAdded: string;
}
export declare const WinningPlayer: {
    encode(message: WinningPlayer, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): WinningPlayer;
    fromJSON(object: any): WinningPlayer;
    toJSON(message: WinningPlayer): unknown;
    fromPartial(object: DeepPartial<WinningPlayer>): WinningPlayer;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
