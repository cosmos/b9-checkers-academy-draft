/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "b9lab.checkers.leaderboard";

export interface Leaderboard {
  winners: Winner[];
}

export interface Winner {
  address: string;
  wonCount: number;
  addedAt: number;
}

export interface Candidate {
  address: Uint8Array;
  wonCount: number;
}

const baseLeaderboard: object = {};

export const Leaderboard = {
  encode(message: Leaderboard, writer: Writer = Writer.create()): Writer {
    for (const v of message.winners) {
      Winner.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Leaderboard {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseLeaderboard } as Leaderboard;
    message.winners = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.winners.push(Winner.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Leaderboard {
    const message = { ...baseLeaderboard } as Leaderboard;
    message.winners = [];
    if (object.winners !== undefined && object.winners !== null) {
      for (const e of object.winners) {
        message.winners.push(Winner.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: Leaderboard): unknown {
    const obj: any = {};
    if (message.winners) {
      obj.winners = message.winners.map((e) =>
        e ? Winner.toJSON(e) : undefined
      );
    } else {
      obj.winners = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Leaderboard>): Leaderboard {
    const message = { ...baseLeaderboard } as Leaderboard;
    message.winners = [];
    if (object.winners !== undefined && object.winners !== null) {
      for (const e of object.winners) {
        message.winners.push(Winner.fromPartial(e));
      }
    }
    return message;
  },
};

const baseWinner: object = { address: "", wonCount: 0, addedAt: 0 };

export const Winner = {
  encode(message: Winner, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.wonCount !== 0) {
      writer.uint32(16).uint64(message.wonCount);
    }
    if (message.addedAt !== 0) {
      writer.uint32(24).uint64(message.addedAt);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Winner {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseWinner } as Winner;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.wonCount = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.addedAt = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Winner {
    const message = { ...baseWinner } as Winner;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.wonCount !== undefined && object.wonCount !== null) {
      message.wonCount = Number(object.wonCount);
    } else {
      message.wonCount = 0;
    }
    if (object.addedAt !== undefined && object.addedAt !== null) {
      message.addedAt = Number(object.addedAt);
    } else {
      message.addedAt = 0;
    }
    return message;
  },

  toJSON(message: Winner): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.wonCount !== undefined && (obj.wonCount = message.wonCount);
    message.addedAt !== undefined && (obj.addedAt = message.addedAt);
    return obj;
  },

  fromPartial(object: DeepPartial<Winner>): Winner {
    const message = { ...baseWinner } as Winner;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.wonCount !== undefined && object.wonCount !== null) {
      message.wonCount = object.wonCount;
    } else {
      message.wonCount = 0;
    }
    if (object.addedAt !== undefined && object.addedAt !== null) {
      message.addedAt = object.addedAt;
    } else {
      message.addedAt = 0;
    }
    return message;
  },
};

const baseCandidate: object = { wonCount: 0 };

export const Candidate = {
  encode(message: Candidate, writer: Writer = Writer.create()): Writer {
    if (message.address.length !== 0) {
      writer.uint32(10).bytes(message.address);
    }
    if (message.wonCount !== 0) {
      writer.uint32(16).uint64(message.wonCount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Candidate {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCandidate } as Candidate;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.bytes();
          break;
        case 2:
          message.wonCount = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Candidate {
    const message = { ...baseCandidate } as Candidate;
    if (object.address !== undefined && object.address !== null) {
      message.address = bytesFromBase64(object.address);
    }
    if (object.wonCount !== undefined && object.wonCount !== null) {
      message.wonCount = Number(object.wonCount);
    } else {
      message.wonCount = 0;
    }
    return message;
  },

  toJSON(message: Candidate): unknown {
    const obj: any = {};
    message.address !== undefined &&
      (obj.address = base64FromBytes(
        message.address !== undefined ? message.address : new Uint8Array()
      ));
    message.wonCount !== undefined && (obj.wonCount = message.wonCount);
    return obj;
  },

  fromPartial(object: DeepPartial<Candidate>): Candidate {
    const message = { ...baseCandidate } as Candidate;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = new Uint8Array();
    }
    if (object.wonCount !== undefined && object.wonCount !== null) {
      message.wonCount = object.wonCount;
    } else {
      message.wonCount = 0;
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

const atob: (b64: string) => string =
  globalThis.atob ||
  ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

const btoa: (bin: string) => string =
  globalThis.btoa ||
  ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  for (let i = 0; i < arr.byteLength; ++i) {
    bin.push(String.fromCharCode(arr[i]));
  }
  return btoa(bin.join(""));
}

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
