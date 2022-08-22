/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../checkers/params";
import { SystemInfo } from "../checkers/system_info";
import { StoredGame } from "../checkers/stored_game";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { PlayerInfo } from "../checkers/player_info";
import { Leaderboard } from "../checkers/leaderboard";

export const protobufPackage = "b9lab.checkers.checkers";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetSystemInfoRequest {}

export interface QueryGetSystemInfoResponse {
  SystemInfo: SystemInfo | undefined;
}

export interface QueryGetStoredGameRequest {
  index: string;
}

export interface QueryGetStoredGameResponse {
  storedGame: StoredGame | undefined;
}

export interface QueryAllStoredGameRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllStoredGameResponse {
  storedGame: StoredGame[];
  pagination: PageResponse | undefined;
}

export interface QueryCanPlayMoveRequest {
  gameIndex: string;
  player: string;
  fromX: number;
  fromY: number;
  toX: number;
  toY: number;
}

export interface QueryCanPlayMoveResponse {
  possible: boolean;
  reason: string;
}

export interface QueryGetPlayerInfoRequest {
  index: string;
}

export interface QueryGetPlayerInfoResponse {
  playerInfo: PlayerInfo | undefined;
}

export interface QueryAllPlayerInfoRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllPlayerInfoResponse {
  playerInfo: PlayerInfo[];
  pagination: PageResponse | undefined;
}

export interface QueryGetLeaderboardRequest {}

export interface QueryGetLeaderboardResponse {
  Leaderboard: Leaderboard | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetSystemInfoRequest: object = {};

export const QueryGetSystemInfoRequest = {
  encode(
    _: QueryGetSystemInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSystemInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSystemInfoRequest,
    } as QueryGetSystemInfoRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryGetSystemInfoRequest {
    const message = {
      ...baseQueryGetSystemInfoRequest,
    } as QueryGetSystemInfoRequest;
    return message;
  },

  toJSON(_: QueryGetSystemInfoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetSystemInfoRequest>
  ): QueryGetSystemInfoRequest {
    const message = {
      ...baseQueryGetSystemInfoRequest,
    } as QueryGetSystemInfoRequest;
    return message;
  },
};

const baseQueryGetSystemInfoResponse: object = {};

export const QueryGetSystemInfoResponse = {
  encode(
    message: QueryGetSystemInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.SystemInfo !== undefined) {
      SystemInfo.encode(message.SystemInfo, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSystemInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSystemInfoResponse,
    } as QueryGetSystemInfoResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.SystemInfo = SystemInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSystemInfoResponse {
    const message = {
      ...baseQueryGetSystemInfoResponse,
    } as QueryGetSystemInfoResponse;
    if (object.SystemInfo !== undefined && object.SystemInfo !== null) {
      message.SystemInfo = SystemInfo.fromJSON(object.SystemInfo);
    } else {
      message.SystemInfo = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetSystemInfoResponse): unknown {
    const obj: any = {};
    message.SystemInfo !== undefined &&
      (obj.SystemInfo = message.SystemInfo
        ? SystemInfo.toJSON(message.SystemInfo)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSystemInfoResponse>
  ): QueryGetSystemInfoResponse {
    const message = {
      ...baseQueryGetSystemInfoResponse,
    } as QueryGetSystemInfoResponse;
    if (object.SystemInfo !== undefined && object.SystemInfo !== null) {
      message.SystemInfo = SystemInfo.fromPartial(object.SystemInfo);
    } else {
      message.SystemInfo = undefined;
    }
    return message;
  },
};

const baseQueryGetStoredGameRequest: object = { index: "" };

export const QueryGetStoredGameRequest = {
  encode(
    message: QueryGetStoredGameRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetStoredGameRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetStoredGameRequest,
    } as QueryGetStoredGameRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStoredGameRequest {
    const message = {
      ...baseQueryGetStoredGameRequest,
    } as QueryGetStoredGameRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetStoredGameRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetStoredGameRequest>
  ): QueryGetStoredGameRequest {
    const message = {
      ...baseQueryGetStoredGameRequest,
    } as QueryGetStoredGameRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetStoredGameResponse: object = {};

export const QueryGetStoredGameResponse = {
  encode(
    message: QueryGetStoredGameResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.storedGame !== undefined) {
      StoredGame.encode(message.storedGame, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetStoredGameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetStoredGameResponse,
    } as QueryGetStoredGameResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storedGame = StoredGame.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStoredGameResponse {
    const message = {
      ...baseQueryGetStoredGameResponse,
    } as QueryGetStoredGameResponse;
    if (object.storedGame !== undefined && object.storedGame !== null) {
      message.storedGame = StoredGame.fromJSON(object.storedGame);
    } else {
      message.storedGame = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetStoredGameResponse): unknown {
    const obj: any = {};
    message.storedGame !== undefined &&
      (obj.storedGame = message.storedGame
        ? StoredGame.toJSON(message.storedGame)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetStoredGameResponse>
  ): QueryGetStoredGameResponse {
    const message = {
      ...baseQueryGetStoredGameResponse,
    } as QueryGetStoredGameResponse;
    if (object.storedGame !== undefined && object.storedGame !== null) {
      message.storedGame = StoredGame.fromPartial(object.storedGame);
    } else {
      message.storedGame = undefined;
    }
    return message;
  },
};

const baseQueryAllStoredGameRequest: object = {};

export const QueryAllStoredGameRequest = {
  encode(
    message: QueryAllStoredGameRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllStoredGameRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllStoredGameRequest,
    } as QueryAllStoredGameRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllStoredGameRequest {
    const message = {
      ...baseQueryAllStoredGameRequest,
    } as QueryAllStoredGameRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllStoredGameRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllStoredGameRequest>
  ): QueryAllStoredGameRequest {
    const message = {
      ...baseQueryAllStoredGameRequest,
    } as QueryAllStoredGameRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllStoredGameResponse: object = {};

export const QueryAllStoredGameResponse = {
  encode(
    message: QueryAllStoredGameResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.storedGame) {
      StoredGame.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllStoredGameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllStoredGameResponse,
    } as QueryAllStoredGameResponse;
    message.storedGame = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storedGame.push(StoredGame.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllStoredGameResponse {
    const message = {
      ...baseQueryAllStoredGameResponse,
    } as QueryAllStoredGameResponse;
    message.storedGame = [];
    if (object.storedGame !== undefined && object.storedGame !== null) {
      for (const e of object.storedGame) {
        message.storedGame.push(StoredGame.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllStoredGameResponse): unknown {
    const obj: any = {};
    if (message.storedGame) {
      obj.storedGame = message.storedGame.map((e) =>
        e ? StoredGame.toJSON(e) : undefined
      );
    } else {
      obj.storedGame = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllStoredGameResponse>
  ): QueryAllStoredGameResponse {
    const message = {
      ...baseQueryAllStoredGameResponse,
    } as QueryAllStoredGameResponse;
    message.storedGame = [];
    if (object.storedGame !== undefined && object.storedGame !== null) {
      for (const e of object.storedGame) {
        message.storedGame.push(StoredGame.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryCanPlayMoveRequest: object = {
  gameIndex: "",
  player: "",
  fromX: 0,
  fromY: 0,
  toX: 0,
  toY: 0,
};

export const QueryCanPlayMoveRequest = {
  encode(
    message: QueryCanPlayMoveRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.gameIndex !== "") {
      writer.uint32(10).string(message.gameIndex);
    }
    if (message.player !== "") {
      writer.uint32(18).string(message.player);
    }
    if (message.fromX !== 0) {
      writer.uint32(24).uint64(message.fromX);
    }
    if (message.fromY !== 0) {
      writer.uint32(32).uint64(message.fromY);
    }
    if (message.toX !== 0) {
      writer.uint32(40).uint64(message.toX);
    }
    if (message.toY !== 0) {
      writer.uint32(48).uint64(message.toY);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryCanPlayMoveRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryCanPlayMoveRequest,
    } as QueryCanPlayMoveRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gameIndex = reader.string();
          break;
        case 2:
          message.player = reader.string();
          break;
        case 3:
          message.fromX = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.fromY = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.toX = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.toY = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCanPlayMoveRequest {
    const message = {
      ...baseQueryCanPlayMoveRequest,
    } as QueryCanPlayMoveRequest;
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = String(object.gameIndex);
    } else {
      message.gameIndex = "";
    }
    if (object.player !== undefined && object.player !== null) {
      message.player = String(object.player);
    } else {
      message.player = "";
    }
    if (object.fromX !== undefined && object.fromX !== null) {
      message.fromX = Number(object.fromX);
    } else {
      message.fromX = 0;
    }
    if (object.fromY !== undefined && object.fromY !== null) {
      message.fromY = Number(object.fromY);
    } else {
      message.fromY = 0;
    }
    if (object.toX !== undefined && object.toX !== null) {
      message.toX = Number(object.toX);
    } else {
      message.toX = 0;
    }
    if (object.toY !== undefined && object.toY !== null) {
      message.toY = Number(object.toY);
    } else {
      message.toY = 0;
    }
    return message;
  },

  toJSON(message: QueryCanPlayMoveRequest): unknown {
    const obj: any = {};
    message.gameIndex !== undefined && (obj.gameIndex = message.gameIndex);
    message.player !== undefined && (obj.player = message.player);
    message.fromX !== undefined && (obj.fromX = message.fromX);
    message.fromY !== undefined && (obj.fromY = message.fromY);
    message.toX !== undefined && (obj.toX = message.toX);
    message.toY !== undefined && (obj.toY = message.toY);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCanPlayMoveRequest>
  ): QueryCanPlayMoveRequest {
    const message = {
      ...baseQueryCanPlayMoveRequest,
    } as QueryCanPlayMoveRequest;
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = object.gameIndex;
    } else {
      message.gameIndex = "";
    }
    if (object.player !== undefined && object.player !== null) {
      message.player = object.player;
    } else {
      message.player = "";
    }
    if (object.fromX !== undefined && object.fromX !== null) {
      message.fromX = object.fromX;
    } else {
      message.fromX = 0;
    }
    if (object.fromY !== undefined && object.fromY !== null) {
      message.fromY = object.fromY;
    } else {
      message.fromY = 0;
    }
    if (object.toX !== undefined && object.toX !== null) {
      message.toX = object.toX;
    } else {
      message.toX = 0;
    }
    if (object.toY !== undefined && object.toY !== null) {
      message.toY = object.toY;
    } else {
      message.toY = 0;
    }
    return message;
  },
};

const baseQueryCanPlayMoveResponse: object = { possible: false, reason: "" };

export const QueryCanPlayMoveResponse = {
  encode(
    message: QueryCanPlayMoveResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.possible === true) {
      writer.uint32(8).bool(message.possible);
    }
    if (message.reason !== "") {
      writer.uint32(18).string(message.reason);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryCanPlayMoveResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryCanPlayMoveResponse,
    } as QueryCanPlayMoveResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.possible = reader.bool();
          break;
        case 2:
          message.reason = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCanPlayMoveResponse {
    const message = {
      ...baseQueryCanPlayMoveResponse,
    } as QueryCanPlayMoveResponse;
    if (object.possible !== undefined && object.possible !== null) {
      message.possible = Boolean(object.possible);
    } else {
      message.possible = false;
    }
    if (object.reason !== undefined && object.reason !== null) {
      message.reason = String(object.reason);
    } else {
      message.reason = "";
    }
    return message;
  },

  toJSON(message: QueryCanPlayMoveResponse): unknown {
    const obj: any = {};
    message.possible !== undefined && (obj.possible = message.possible);
    message.reason !== undefined && (obj.reason = message.reason);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCanPlayMoveResponse>
  ): QueryCanPlayMoveResponse {
    const message = {
      ...baseQueryCanPlayMoveResponse,
    } as QueryCanPlayMoveResponse;
    if (object.possible !== undefined && object.possible !== null) {
      message.possible = object.possible;
    } else {
      message.possible = false;
    }
    if (object.reason !== undefined && object.reason !== null) {
      message.reason = object.reason;
    } else {
      message.reason = "";
    }
    return message;
  },
};

const baseQueryGetPlayerInfoRequest: object = { index: "" };

export const QueryGetPlayerInfoRequest = {
  encode(
    message: QueryGetPlayerInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetPlayerInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetPlayerInfoRequest,
    } as QueryGetPlayerInfoRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPlayerInfoRequest {
    const message = {
      ...baseQueryGetPlayerInfoRequest,
    } as QueryGetPlayerInfoRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetPlayerInfoRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPlayerInfoRequest>
  ): QueryGetPlayerInfoRequest {
    const message = {
      ...baseQueryGetPlayerInfoRequest,
    } as QueryGetPlayerInfoRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetPlayerInfoResponse: object = {};

export const QueryGetPlayerInfoResponse = {
  encode(
    message: QueryGetPlayerInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.playerInfo !== undefined) {
      PlayerInfo.encode(message.playerInfo, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetPlayerInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetPlayerInfoResponse,
    } as QueryGetPlayerInfoResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.playerInfo = PlayerInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPlayerInfoResponse {
    const message = {
      ...baseQueryGetPlayerInfoResponse,
    } as QueryGetPlayerInfoResponse;
    if (object.playerInfo !== undefined && object.playerInfo !== null) {
      message.playerInfo = PlayerInfo.fromJSON(object.playerInfo);
    } else {
      message.playerInfo = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetPlayerInfoResponse): unknown {
    const obj: any = {};
    message.playerInfo !== undefined &&
      (obj.playerInfo = message.playerInfo
        ? PlayerInfo.toJSON(message.playerInfo)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPlayerInfoResponse>
  ): QueryGetPlayerInfoResponse {
    const message = {
      ...baseQueryGetPlayerInfoResponse,
    } as QueryGetPlayerInfoResponse;
    if (object.playerInfo !== undefined && object.playerInfo !== null) {
      message.playerInfo = PlayerInfo.fromPartial(object.playerInfo);
    } else {
      message.playerInfo = undefined;
    }
    return message;
  },
};

const baseQueryAllPlayerInfoRequest: object = {};

export const QueryAllPlayerInfoRequest = {
  encode(
    message: QueryAllPlayerInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllPlayerInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllPlayerInfoRequest,
    } as QueryAllPlayerInfoRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllPlayerInfoRequest {
    const message = {
      ...baseQueryAllPlayerInfoRequest,
    } as QueryAllPlayerInfoRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPlayerInfoRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllPlayerInfoRequest>
  ): QueryAllPlayerInfoRequest {
    const message = {
      ...baseQueryAllPlayerInfoRequest,
    } as QueryAllPlayerInfoRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllPlayerInfoResponse: object = {};

export const QueryAllPlayerInfoResponse = {
  encode(
    message: QueryAllPlayerInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.playerInfo) {
      PlayerInfo.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllPlayerInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllPlayerInfoResponse,
    } as QueryAllPlayerInfoResponse;
    message.playerInfo = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.playerInfo.push(PlayerInfo.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllPlayerInfoResponse {
    const message = {
      ...baseQueryAllPlayerInfoResponse,
    } as QueryAllPlayerInfoResponse;
    message.playerInfo = [];
    if (object.playerInfo !== undefined && object.playerInfo !== null) {
      for (const e of object.playerInfo) {
        message.playerInfo.push(PlayerInfo.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPlayerInfoResponse): unknown {
    const obj: any = {};
    if (message.playerInfo) {
      obj.playerInfo = message.playerInfo.map((e) =>
        e ? PlayerInfo.toJSON(e) : undefined
      );
    } else {
      obj.playerInfo = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllPlayerInfoResponse>
  ): QueryAllPlayerInfoResponse {
    const message = {
      ...baseQueryAllPlayerInfoResponse,
    } as QueryAllPlayerInfoResponse;
    message.playerInfo = [];
    if (object.playerInfo !== undefined && object.playerInfo !== null) {
      for (const e of object.playerInfo) {
        message.playerInfo.push(PlayerInfo.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetLeaderboardRequest: object = {};

export const QueryGetLeaderboardRequest = {
  encode(
    _: QueryGetLeaderboardRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLeaderboardRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLeaderboardRequest,
    } as QueryGetLeaderboardRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryGetLeaderboardRequest {
    const message = {
      ...baseQueryGetLeaderboardRequest,
    } as QueryGetLeaderboardRequest;
    return message;
  },

  toJSON(_: QueryGetLeaderboardRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetLeaderboardRequest>
  ): QueryGetLeaderboardRequest {
    const message = {
      ...baseQueryGetLeaderboardRequest,
    } as QueryGetLeaderboardRequest;
    return message;
  },
};

const baseQueryGetLeaderboardResponse: object = {};

export const QueryGetLeaderboardResponse = {
  encode(
    message: QueryGetLeaderboardResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Leaderboard !== undefined) {
      Leaderboard.encode(
        message.Leaderboard,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLeaderboardResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLeaderboardResponse,
    } as QueryGetLeaderboardResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Leaderboard = Leaderboard.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLeaderboardResponse {
    const message = {
      ...baseQueryGetLeaderboardResponse,
    } as QueryGetLeaderboardResponse;
    if (object.Leaderboard !== undefined && object.Leaderboard !== null) {
      message.Leaderboard = Leaderboard.fromJSON(object.Leaderboard);
    } else {
      message.Leaderboard = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetLeaderboardResponse): unknown {
    const obj: any = {};
    message.Leaderboard !== undefined &&
      (obj.Leaderboard = message.Leaderboard
        ? Leaderboard.toJSON(message.Leaderboard)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLeaderboardResponse>
  ): QueryGetLeaderboardResponse {
    const message = {
      ...baseQueryGetLeaderboardResponse,
    } as QueryGetLeaderboardResponse;
    if (object.Leaderboard !== undefined && object.Leaderboard !== null) {
      message.Leaderboard = Leaderboard.fromPartial(object.Leaderboard);
    } else {
      message.Leaderboard = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a SystemInfo by index. */
  SystemInfo(
    request: QueryGetSystemInfoRequest
  ): Promise<QueryGetSystemInfoResponse>;
  /** Queries a StoredGame by index. */
  StoredGame(
    request: QueryGetStoredGameRequest
  ): Promise<QueryGetStoredGameResponse>;
  /** Queries a list of StoredGame items. */
  StoredGameAll(
    request: QueryAllStoredGameRequest
  ): Promise<QueryAllStoredGameResponse>;
  /** Queries a list of CanPlayMove items. */
  CanPlayMove(
    request: QueryCanPlayMoveRequest
  ): Promise<QueryCanPlayMoveResponse>;
  /** Queries a PlayerInfo by index. */
  PlayerInfo(
    request: QueryGetPlayerInfoRequest
  ): Promise<QueryGetPlayerInfoResponse>;
  /** Queries a list of PlayerInfo items. */
  PlayerInfoAll(
    request: QueryAllPlayerInfoRequest
  ): Promise<QueryAllPlayerInfoResponse>;
  /** Queries a Leaderboard by index. */
  Leaderboard(
    request: QueryGetLeaderboardRequest
  ): Promise<QueryGetLeaderboardResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "b9lab.checkers.checkers.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  SystemInfo(
    request: QueryGetSystemInfoRequest
  ): Promise<QueryGetSystemInfoResponse> {
    const data = QueryGetSystemInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "b9lab.checkers.checkers.Query",
      "SystemInfo",
      data
    );
    return promise.then((data) =>
      QueryGetSystemInfoResponse.decode(new Reader(data))
    );
  }

  StoredGame(
    request: QueryGetStoredGameRequest
  ): Promise<QueryGetStoredGameResponse> {
    const data = QueryGetStoredGameRequest.encode(request).finish();
    const promise = this.rpc.request(
      "b9lab.checkers.checkers.Query",
      "StoredGame",
      data
    );
    return promise.then((data) =>
      QueryGetStoredGameResponse.decode(new Reader(data))
    );
  }

  StoredGameAll(
    request: QueryAllStoredGameRequest
  ): Promise<QueryAllStoredGameResponse> {
    const data = QueryAllStoredGameRequest.encode(request).finish();
    const promise = this.rpc.request(
      "b9lab.checkers.checkers.Query",
      "StoredGameAll",
      data
    );
    return promise.then((data) =>
      QueryAllStoredGameResponse.decode(new Reader(data))
    );
  }

  CanPlayMove(
    request: QueryCanPlayMoveRequest
  ): Promise<QueryCanPlayMoveResponse> {
    const data = QueryCanPlayMoveRequest.encode(request).finish();
    const promise = this.rpc.request(
      "b9lab.checkers.checkers.Query",
      "CanPlayMove",
      data
    );
    return promise.then((data) =>
      QueryCanPlayMoveResponse.decode(new Reader(data))
    );
  }

  PlayerInfo(
    request: QueryGetPlayerInfoRequest
  ): Promise<QueryGetPlayerInfoResponse> {
    const data = QueryGetPlayerInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "b9lab.checkers.checkers.Query",
      "PlayerInfo",
      data
    );
    return promise.then((data) =>
      QueryGetPlayerInfoResponse.decode(new Reader(data))
    );
  }

  PlayerInfoAll(
    request: QueryAllPlayerInfoRequest
  ): Promise<QueryAllPlayerInfoResponse> {
    const data = QueryAllPlayerInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "b9lab.checkers.checkers.Query",
      "PlayerInfoAll",
      data
    );
    return promise.then((data) =>
      QueryAllPlayerInfoResponse.decode(new Reader(data))
    );
  }

  Leaderboard(
    request: QueryGetLeaderboardRequest
  ): Promise<QueryGetLeaderboardResponse> {
    const data = QueryGetLeaderboardRequest.encode(request).finish();
    const promise = this.rpc.request(
      "b9lab.checkers.checkers.Query",
      "Leaderboard",
      data
    );
    return promise.then((data) =>
      QueryGetLeaderboardResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
