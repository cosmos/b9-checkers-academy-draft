import { WinningPlayer } from '../checkers/winning_player';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "b9lab.checkers.checkers";
export interface Leaderboard {
    winners: WinningPlayer[];
}
export declare const Leaderboard: {
    encode(message: Leaderboard, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Leaderboard;
    fromJSON(object: any): Leaderboard;
    toJSON(message: Leaderboard): unknown;
    fromPartial(object: DeepPartial<Leaderboard>): Leaderboard;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
