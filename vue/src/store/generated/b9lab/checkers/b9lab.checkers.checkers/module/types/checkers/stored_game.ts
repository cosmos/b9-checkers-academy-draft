/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "b9lab.checkers.checkers";

export interface StoredGame {
  index: string;
  board: string;
  turn: string;
  black: string;
  red: string;
  moveCount: number;
  beforeIndex: string;
  afterIndex: string;
  deadline: string;
  winner: string;
  wager: number;
}

const baseStoredGame: object = {
  index: "",
  board: "",
  turn: "",
  black: "",
  red: "",
  moveCount: 0,
  beforeIndex: "",
  afterIndex: "",
  deadline: "",
  winner: "",
  wager: 0,
};

export const StoredGame = {
  encode(message: StoredGame, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.board !== "") {
      writer.uint32(18).string(message.board);
    }
    if (message.turn !== "") {
      writer.uint32(26).string(message.turn);
    }
    if (message.black !== "") {
      writer.uint32(34).string(message.black);
    }
    if (message.red !== "") {
      writer.uint32(42).string(message.red);
    }
    if (message.moveCount !== 0) {
      writer.uint32(48).uint64(message.moveCount);
    }
    if (message.beforeIndex !== "") {
      writer.uint32(58).string(message.beforeIndex);
    }
    if (message.afterIndex !== "") {
      writer.uint32(66).string(message.afterIndex);
    }
    if (message.deadline !== "") {
      writer.uint32(74).string(message.deadline);
    }
    if (message.winner !== "") {
      writer.uint32(82).string(message.winner);
    }
    if (message.wager !== 0) {
      writer.uint32(88).uint64(message.wager);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): StoredGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStoredGame } as StoredGame;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.board = reader.string();
          break;
        case 3:
          message.turn = reader.string();
          break;
        case 4:
          message.black = reader.string();
          break;
        case 5:
          message.red = reader.string();
          break;
        case 6:
          message.moveCount = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.beforeIndex = reader.string();
          break;
        case 8:
          message.afterIndex = reader.string();
          break;
        case 9:
          message.deadline = reader.string();
          break;
        case 10:
          message.winner = reader.string();
          break;
        case 11:
          message.wager = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StoredGame {
    const message = { ...baseStoredGame } as StoredGame;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.board !== undefined && object.board !== null) {
      message.board = String(object.board);
    } else {
      message.board = "";
    }
    if (object.turn !== undefined && object.turn !== null) {
      message.turn = String(object.turn);
    } else {
      message.turn = "";
    }
    if (object.black !== undefined && object.black !== null) {
      message.black = String(object.black);
    } else {
      message.black = "";
    }
    if (object.red !== undefined && object.red !== null) {
      message.red = String(object.red);
    } else {
      message.red = "";
    }
    if (object.moveCount !== undefined && object.moveCount !== null) {
      message.moveCount = Number(object.moveCount);
    } else {
      message.moveCount = 0;
    }
    if (object.beforeIndex !== undefined && object.beforeIndex !== null) {
      message.beforeIndex = String(object.beforeIndex);
    } else {
      message.beforeIndex = "";
    }
    if (object.afterIndex !== undefined && object.afterIndex !== null) {
      message.afterIndex = String(object.afterIndex);
    } else {
      message.afterIndex = "";
    }
    if (object.deadline !== undefined && object.deadline !== null) {
      message.deadline = String(object.deadline);
    } else {
      message.deadline = "";
    }
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = String(object.winner);
    } else {
      message.winner = "";
    }
    if (object.wager !== undefined && object.wager !== null) {
      message.wager = Number(object.wager);
    } else {
      message.wager = 0;
    }
    return message;
  },

  toJSON(message: StoredGame): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.board !== undefined && (obj.board = message.board);
    message.turn !== undefined && (obj.turn = message.turn);
    message.black !== undefined && (obj.black = message.black);
    message.red !== undefined && (obj.red = message.red);
    message.moveCount !== undefined && (obj.moveCount = message.moveCount);
    message.beforeIndex !== undefined &&
      (obj.beforeIndex = message.beforeIndex);
    message.afterIndex !== undefined && (obj.afterIndex = message.afterIndex);
    message.deadline !== undefined && (obj.deadline = message.deadline);
    message.winner !== undefined && (obj.winner = message.winner);
    message.wager !== undefined && (obj.wager = message.wager);
    return obj;
  },

  fromPartial(object: DeepPartial<StoredGame>): StoredGame {
    const message = { ...baseStoredGame } as StoredGame;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.board !== undefined && object.board !== null) {
      message.board = object.board;
    } else {
      message.board = "";
    }
    if (object.turn !== undefined && object.turn !== null) {
      message.turn = object.turn;
    } else {
      message.turn = "";
    }
    if (object.black !== undefined && object.black !== null) {
      message.black = object.black;
    } else {
      message.black = "";
    }
    if (object.red !== undefined && object.red !== null) {
      message.red = object.red;
    } else {
      message.red = "";
    }
    if (object.moveCount !== undefined && object.moveCount !== null) {
      message.moveCount = object.moveCount;
    } else {
      message.moveCount = 0;
    }
    if (object.beforeIndex !== undefined && object.beforeIndex !== null) {
      message.beforeIndex = object.beforeIndex;
    } else {
      message.beforeIndex = "";
    }
    if (object.afterIndex !== undefined && object.afterIndex !== null) {
      message.afterIndex = object.afterIndex;
    } else {
      message.afterIndex = "";
    }
    if (object.deadline !== undefined && object.deadline !== null) {
      message.deadline = object.deadline;
    } else {
      message.deadline = "";
    }
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = object.winner;
    } else {
      message.winner = "";
    }
    if (object.wager !== undefined && object.wager !== null) {
      message.wager = object.wager;
    } else {
      message.wager = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
