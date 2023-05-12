/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "b9lab.checkers.checkers";

export interface PlayerInfo {
  index: string;
  wonCount: number;
  lostCount: number;
  forfeitedCount: number;
}

const basePlayerInfo: object = {
  index: "",
  wonCount: 0,
  lostCount: 0,
  forfeitedCount: 0,
};

export const PlayerInfo = {
  encode(message: PlayerInfo, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.wonCount !== 0) {
      writer.uint32(16).uint64(message.wonCount);
    }
    if (message.lostCount !== 0) {
      writer.uint32(24).uint64(message.lostCount);
    }
    if (message.forfeitedCount !== 0) {
      writer.uint32(32).uint64(message.forfeitedCount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): PlayerInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePlayerInfo } as PlayerInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.wonCount = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.lostCount = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.forfeitedCount = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PlayerInfo {
    const message = { ...basePlayerInfo } as PlayerInfo;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.wonCount !== undefined && object.wonCount !== null) {
      message.wonCount = Number(object.wonCount);
    } else {
      message.wonCount = 0;
    }
    if (object.lostCount !== undefined && object.lostCount !== null) {
      message.lostCount = Number(object.lostCount);
    } else {
      message.lostCount = 0;
    }
    if (object.forfeitedCount !== undefined && object.forfeitedCount !== null) {
      message.forfeitedCount = Number(object.forfeitedCount);
    } else {
      message.forfeitedCount = 0;
    }
    return message;
  },

  toJSON(message: PlayerInfo): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.wonCount !== undefined && (obj.wonCount = message.wonCount);
    message.lostCount !== undefined && (obj.lostCount = message.lostCount);
    message.forfeitedCount !== undefined &&
      (obj.forfeitedCount = message.forfeitedCount);
    return obj;
  },

  fromPartial(object: DeepPartial<PlayerInfo>): PlayerInfo {
    const message = { ...basePlayerInfo } as PlayerInfo;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.wonCount !== undefined && object.wonCount !== null) {
      message.wonCount = object.wonCount;
    } else {
      message.wonCount = 0;
    }
    if (object.lostCount !== undefined && object.lostCount !== null) {
      message.lostCount = object.lostCount;
    } else {
      message.lostCount = 0;
    }
    if (object.forfeitedCount !== undefined && object.forfeitedCount !== null) {
      message.forfeitedCount = object.forfeitedCount;
    } else {
      message.forfeitedCount = 0;
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
