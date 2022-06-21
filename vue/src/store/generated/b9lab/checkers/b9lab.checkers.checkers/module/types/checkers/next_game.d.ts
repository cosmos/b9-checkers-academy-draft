import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "b9lab.checkers.checkers";
export interface NextGame {
    creator: string;
    idValue: number;
    fifoHead: string;
    fifoTail: string;
}
export declare const NextGame: {
    encode(message: NextGame, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): NextGame;
    fromJSON(object: any): NextGame;
    toJSON(message: NextGame): unknown;
    fromPartial(object: DeepPartial<NextGame>): NextGame;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
