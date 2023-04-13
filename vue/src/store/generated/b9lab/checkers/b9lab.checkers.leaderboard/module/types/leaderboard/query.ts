/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../leaderboard/params";
import { Leaderboard } from "../leaderboard/leaderboard";

export const protobufPackage = "b9lab.checkers.leaderboard";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
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
      "b9lab.checkers.leaderboard.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Leaderboard(
    request: QueryGetLeaderboardRequest
  ): Promise<QueryGetLeaderboardResponse> {
    const data = QueryGetLeaderboardRequest.encode(request).finish();
    const promise = this.rpc.request(
      "b9lab.checkers.leaderboard.Query",
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
