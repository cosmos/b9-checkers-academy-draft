/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'
import { PlayerInfo } from '../checkers/player_info'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'
import { StoredGame } from '../checkers/stored_game'
import { NextGame } from '../checkers/next_game'

export const protobufPackage = 'xavierlepretre.checkers.checkers'

/** this line is used by starport scaffolding # 3 */
export interface QueryGetPlayerInfoRequest {
  index: string
}

export interface QueryGetPlayerInfoResponse {
  PlayerInfo: PlayerInfo | undefined
}

export interface QueryAllPlayerInfoRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllPlayerInfoResponse {
  PlayerInfo: PlayerInfo[]
  pagination: PageResponse | undefined
}

export interface QueryCanPlayMoveRequest {
  idValue: string
  player: string
  fromX: number
  fromY: number
  toX: number
  toY: number
}

export interface QueryCanPlayMoveResponse {
  possible: boolean
  reason: string
}

export interface QueryGetStoredGameRequest {
  index: string
}

export interface QueryGetStoredGameResponse {
  StoredGame: StoredGame | undefined
}

export interface QueryAllStoredGameRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllStoredGameResponse {
  StoredGame: StoredGame[]
  pagination: PageResponse | undefined
}

export interface QueryGetNextGameRequest {}

export interface QueryGetNextGameResponse {
  NextGame: NextGame | undefined
}

const baseQueryGetPlayerInfoRequest: object = { index: '' }

export const QueryGetPlayerInfoRequest = {
  encode(message: QueryGetPlayerInfoRequest, writer: Writer = Writer.create()): Writer {
    if (message.index !== '') {
      writer.uint32(10).string(message.index)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPlayerInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetPlayerInfoRequest } as QueryGetPlayerInfoRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetPlayerInfoRequest {
    const message = { ...baseQueryGetPlayerInfoRequest } as QueryGetPlayerInfoRequest
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index)
    } else {
      message.index = ''
    }
    return message
  },

  toJSON(message: QueryGetPlayerInfoRequest): unknown {
    const obj: any = {}
    message.index !== undefined && (obj.index = message.index)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetPlayerInfoRequest>): QueryGetPlayerInfoRequest {
    const message = { ...baseQueryGetPlayerInfoRequest } as QueryGetPlayerInfoRequest
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index
    } else {
      message.index = ''
    }
    return message
  }
}

const baseQueryGetPlayerInfoResponse: object = {}

export const QueryGetPlayerInfoResponse = {
  encode(message: QueryGetPlayerInfoResponse, writer: Writer = Writer.create()): Writer {
    if (message.PlayerInfo !== undefined) {
      PlayerInfo.encode(message.PlayerInfo, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPlayerInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetPlayerInfoResponse } as QueryGetPlayerInfoResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.PlayerInfo = PlayerInfo.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetPlayerInfoResponse {
    const message = { ...baseQueryGetPlayerInfoResponse } as QueryGetPlayerInfoResponse
    if (object.PlayerInfo !== undefined && object.PlayerInfo !== null) {
      message.PlayerInfo = PlayerInfo.fromJSON(object.PlayerInfo)
    } else {
      message.PlayerInfo = undefined
    }
    return message
  },

  toJSON(message: QueryGetPlayerInfoResponse): unknown {
    const obj: any = {}
    message.PlayerInfo !== undefined && (obj.PlayerInfo = message.PlayerInfo ? PlayerInfo.toJSON(message.PlayerInfo) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetPlayerInfoResponse>): QueryGetPlayerInfoResponse {
    const message = { ...baseQueryGetPlayerInfoResponse } as QueryGetPlayerInfoResponse
    if (object.PlayerInfo !== undefined && object.PlayerInfo !== null) {
      message.PlayerInfo = PlayerInfo.fromPartial(object.PlayerInfo)
    } else {
      message.PlayerInfo = undefined
    }
    return message
  }
}

const baseQueryAllPlayerInfoRequest: object = {}

export const QueryAllPlayerInfoRequest = {
  encode(message: QueryAllPlayerInfoRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPlayerInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllPlayerInfoRequest } as QueryAllPlayerInfoRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllPlayerInfoRequest {
    const message = { ...baseQueryAllPlayerInfoRequest } as QueryAllPlayerInfoRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllPlayerInfoRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllPlayerInfoRequest>): QueryAllPlayerInfoRequest {
    const message = { ...baseQueryAllPlayerInfoRequest } as QueryAllPlayerInfoRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllPlayerInfoResponse: object = {}

export const QueryAllPlayerInfoResponse = {
  encode(message: QueryAllPlayerInfoResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.PlayerInfo) {
      PlayerInfo.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPlayerInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllPlayerInfoResponse } as QueryAllPlayerInfoResponse
    message.PlayerInfo = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.PlayerInfo.push(PlayerInfo.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllPlayerInfoResponse {
    const message = { ...baseQueryAllPlayerInfoResponse } as QueryAllPlayerInfoResponse
    message.PlayerInfo = []
    if (object.PlayerInfo !== undefined && object.PlayerInfo !== null) {
      for (const e of object.PlayerInfo) {
        message.PlayerInfo.push(PlayerInfo.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllPlayerInfoResponse): unknown {
    const obj: any = {}
    if (message.PlayerInfo) {
      obj.PlayerInfo = message.PlayerInfo.map((e) => (e ? PlayerInfo.toJSON(e) : undefined))
    } else {
      obj.PlayerInfo = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllPlayerInfoResponse>): QueryAllPlayerInfoResponse {
    const message = { ...baseQueryAllPlayerInfoResponse } as QueryAllPlayerInfoResponse
    message.PlayerInfo = []
    if (object.PlayerInfo !== undefined && object.PlayerInfo !== null) {
      for (const e of object.PlayerInfo) {
        message.PlayerInfo.push(PlayerInfo.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryCanPlayMoveRequest: object = { idValue: '', player: '', fromX: 0, fromY: 0, toX: 0, toY: 0 }

export const QueryCanPlayMoveRequest = {
  encode(message: QueryCanPlayMoveRequest, writer: Writer = Writer.create()): Writer {
    if (message.idValue !== '') {
      writer.uint32(10).string(message.idValue)
    }
    if (message.player !== '') {
      writer.uint32(18).string(message.player)
    }
    if (message.fromX !== 0) {
      writer.uint32(24).uint64(message.fromX)
    }
    if (message.fromY !== 0) {
      writer.uint32(32).uint64(message.fromY)
    }
    if (message.toX !== 0) {
      writer.uint32(40).uint64(message.toX)
    }
    if (message.toY !== 0) {
      writer.uint32(48).uint64(message.toY)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryCanPlayMoveRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryCanPlayMoveRequest } as QueryCanPlayMoveRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.idValue = reader.string()
          break
        case 2:
          message.player = reader.string()
          break
        case 3:
          message.fromX = longToNumber(reader.uint64() as Long)
          break
        case 4:
          message.fromY = longToNumber(reader.uint64() as Long)
          break
        case 5:
          message.toX = longToNumber(reader.uint64() as Long)
          break
        case 6:
          message.toY = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryCanPlayMoveRequest {
    const message = { ...baseQueryCanPlayMoveRequest } as QueryCanPlayMoveRequest
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = String(object.idValue)
    } else {
      message.idValue = ''
    }
    if (object.player !== undefined && object.player !== null) {
      message.player = String(object.player)
    } else {
      message.player = ''
    }
    if (object.fromX !== undefined && object.fromX !== null) {
      message.fromX = Number(object.fromX)
    } else {
      message.fromX = 0
    }
    if (object.fromY !== undefined && object.fromY !== null) {
      message.fromY = Number(object.fromY)
    } else {
      message.fromY = 0
    }
    if (object.toX !== undefined && object.toX !== null) {
      message.toX = Number(object.toX)
    } else {
      message.toX = 0
    }
    if (object.toY !== undefined && object.toY !== null) {
      message.toY = Number(object.toY)
    } else {
      message.toY = 0
    }
    return message
  },

  toJSON(message: QueryCanPlayMoveRequest): unknown {
    const obj: any = {}
    message.idValue !== undefined && (obj.idValue = message.idValue)
    message.player !== undefined && (obj.player = message.player)
    message.fromX !== undefined && (obj.fromX = message.fromX)
    message.fromY !== undefined && (obj.fromY = message.fromY)
    message.toX !== undefined && (obj.toX = message.toX)
    message.toY !== undefined && (obj.toY = message.toY)
    return obj
  },

  fromPartial(object: DeepPartial<QueryCanPlayMoveRequest>): QueryCanPlayMoveRequest {
    const message = { ...baseQueryCanPlayMoveRequest } as QueryCanPlayMoveRequest
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = object.idValue
    } else {
      message.idValue = ''
    }
    if (object.player !== undefined && object.player !== null) {
      message.player = object.player
    } else {
      message.player = ''
    }
    if (object.fromX !== undefined && object.fromX !== null) {
      message.fromX = object.fromX
    } else {
      message.fromX = 0
    }
    if (object.fromY !== undefined && object.fromY !== null) {
      message.fromY = object.fromY
    } else {
      message.fromY = 0
    }
    if (object.toX !== undefined && object.toX !== null) {
      message.toX = object.toX
    } else {
      message.toX = 0
    }
    if (object.toY !== undefined && object.toY !== null) {
      message.toY = object.toY
    } else {
      message.toY = 0
    }
    return message
  }
}

const baseQueryCanPlayMoveResponse: object = { possible: false, reason: '' }

export const QueryCanPlayMoveResponse = {
  encode(message: QueryCanPlayMoveResponse, writer: Writer = Writer.create()): Writer {
    if (message.possible === true) {
      writer.uint32(8).bool(message.possible)
    }
    if (message.reason !== '') {
      writer.uint32(18).string(message.reason)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryCanPlayMoveResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryCanPlayMoveResponse } as QueryCanPlayMoveResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.possible = reader.bool()
          break
        case 2:
          message.reason = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryCanPlayMoveResponse {
    const message = { ...baseQueryCanPlayMoveResponse } as QueryCanPlayMoveResponse
    if (object.possible !== undefined && object.possible !== null) {
      message.possible = Boolean(object.possible)
    } else {
      message.possible = false
    }
    if (object.reason !== undefined && object.reason !== null) {
      message.reason = String(object.reason)
    } else {
      message.reason = ''
    }
    return message
  },

  toJSON(message: QueryCanPlayMoveResponse): unknown {
    const obj: any = {}
    message.possible !== undefined && (obj.possible = message.possible)
    message.reason !== undefined && (obj.reason = message.reason)
    return obj
  },

  fromPartial(object: DeepPartial<QueryCanPlayMoveResponse>): QueryCanPlayMoveResponse {
    const message = { ...baseQueryCanPlayMoveResponse } as QueryCanPlayMoveResponse
    if (object.possible !== undefined && object.possible !== null) {
      message.possible = object.possible
    } else {
      message.possible = false
    }
    if (object.reason !== undefined && object.reason !== null) {
      message.reason = object.reason
    } else {
      message.reason = ''
    }
    return message
  }
}

const baseQueryGetStoredGameRequest: object = { index: '' }

export const QueryGetStoredGameRequest = {
  encode(message: QueryGetStoredGameRequest, writer: Writer = Writer.create()): Writer {
    if (message.index !== '') {
      writer.uint32(10).string(message.index)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetStoredGameRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetStoredGameRequest } as QueryGetStoredGameRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetStoredGameRequest {
    const message = { ...baseQueryGetStoredGameRequest } as QueryGetStoredGameRequest
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index)
    } else {
      message.index = ''
    }
    return message
  },

  toJSON(message: QueryGetStoredGameRequest): unknown {
    const obj: any = {}
    message.index !== undefined && (obj.index = message.index)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetStoredGameRequest>): QueryGetStoredGameRequest {
    const message = { ...baseQueryGetStoredGameRequest } as QueryGetStoredGameRequest
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index
    } else {
      message.index = ''
    }
    return message
  }
}

const baseQueryGetStoredGameResponse: object = {}

export const QueryGetStoredGameResponse = {
  encode(message: QueryGetStoredGameResponse, writer: Writer = Writer.create()): Writer {
    if (message.StoredGame !== undefined) {
      StoredGame.encode(message.StoredGame, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetStoredGameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetStoredGameResponse } as QueryGetStoredGameResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.StoredGame = StoredGame.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetStoredGameResponse {
    const message = { ...baseQueryGetStoredGameResponse } as QueryGetStoredGameResponse
    if (object.StoredGame !== undefined && object.StoredGame !== null) {
      message.StoredGame = StoredGame.fromJSON(object.StoredGame)
    } else {
      message.StoredGame = undefined
    }
    return message
  },

  toJSON(message: QueryGetStoredGameResponse): unknown {
    const obj: any = {}
    message.StoredGame !== undefined && (obj.StoredGame = message.StoredGame ? StoredGame.toJSON(message.StoredGame) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetStoredGameResponse>): QueryGetStoredGameResponse {
    const message = { ...baseQueryGetStoredGameResponse } as QueryGetStoredGameResponse
    if (object.StoredGame !== undefined && object.StoredGame !== null) {
      message.StoredGame = StoredGame.fromPartial(object.StoredGame)
    } else {
      message.StoredGame = undefined
    }
    return message
  }
}

const baseQueryAllStoredGameRequest: object = {}

export const QueryAllStoredGameRequest = {
  encode(message: QueryAllStoredGameRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllStoredGameRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllStoredGameRequest } as QueryAllStoredGameRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllStoredGameRequest {
    const message = { ...baseQueryAllStoredGameRequest } as QueryAllStoredGameRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllStoredGameRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllStoredGameRequest>): QueryAllStoredGameRequest {
    const message = { ...baseQueryAllStoredGameRequest } as QueryAllStoredGameRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllStoredGameResponse: object = {}

export const QueryAllStoredGameResponse = {
  encode(message: QueryAllStoredGameResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.StoredGame) {
      StoredGame.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllStoredGameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllStoredGameResponse } as QueryAllStoredGameResponse
    message.StoredGame = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.StoredGame.push(StoredGame.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllStoredGameResponse {
    const message = { ...baseQueryAllStoredGameResponse } as QueryAllStoredGameResponse
    message.StoredGame = []
    if (object.StoredGame !== undefined && object.StoredGame !== null) {
      for (const e of object.StoredGame) {
        message.StoredGame.push(StoredGame.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllStoredGameResponse): unknown {
    const obj: any = {}
    if (message.StoredGame) {
      obj.StoredGame = message.StoredGame.map((e) => (e ? StoredGame.toJSON(e) : undefined))
    } else {
      obj.StoredGame = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllStoredGameResponse>): QueryAllStoredGameResponse {
    const message = { ...baseQueryAllStoredGameResponse } as QueryAllStoredGameResponse
    message.StoredGame = []
    if (object.StoredGame !== undefined && object.StoredGame !== null) {
      for (const e of object.StoredGame) {
        message.StoredGame.push(StoredGame.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryGetNextGameRequest: object = {}

export const QueryGetNextGameRequest = {
  encode(_: QueryGetNextGameRequest, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNextGameRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetNextGameRequest } as QueryGetNextGameRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): QueryGetNextGameRequest {
    const message = { ...baseQueryGetNextGameRequest } as QueryGetNextGameRequest
    return message
  },

  toJSON(_: QueryGetNextGameRequest): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<QueryGetNextGameRequest>): QueryGetNextGameRequest {
    const message = { ...baseQueryGetNextGameRequest } as QueryGetNextGameRequest
    return message
  }
}

const baseQueryGetNextGameResponse: object = {}

export const QueryGetNextGameResponse = {
  encode(message: QueryGetNextGameResponse, writer: Writer = Writer.create()): Writer {
    if (message.NextGame !== undefined) {
      NextGame.encode(message.NextGame, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNextGameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetNextGameResponse } as QueryGetNextGameResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.NextGame = NextGame.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetNextGameResponse {
    const message = { ...baseQueryGetNextGameResponse } as QueryGetNextGameResponse
    if (object.NextGame !== undefined && object.NextGame !== null) {
      message.NextGame = NextGame.fromJSON(object.NextGame)
    } else {
      message.NextGame = undefined
    }
    return message
  },

  toJSON(message: QueryGetNextGameResponse): unknown {
    const obj: any = {}
    message.NextGame !== undefined && (obj.NextGame = message.NextGame ? NextGame.toJSON(message.NextGame) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetNextGameResponse>): QueryGetNextGameResponse {
    const message = { ...baseQueryGetNextGameResponse } as QueryGetNextGameResponse
    if (object.NextGame !== undefined && object.NextGame !== null) {
      message.NextGame = NextGame.fromPartial(object.NextGame)
    } else {
      message.NextGame = undefined
    }
    return message
  }
}

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a playerInfo by index. */
  PlayerInfo(request: QueryGetPlayerInfoRequest): Promise<QueryGetPlayerInfoResponse>
  /** Queries a list of playerInfo items. */
  PlayerInfoAll(request: QueryAllPlayerInfoRequest): Promise<QueryAllPlayerInfoResponse>
  /** Queries a list of canPlayMove items. */
  CanPlayMove(request: QueryCanPlayMoveRequest): Promise<QueryCanPlayMoveResponse>
  /** Queries a storedGame by index. */
  StoredGame(request: QueryGetStoredGameRequest): Promise<QueryGetStoredGameResponse>
  /** Queries a list of storedGame items. */
  StoredGameAll(request: QueryAllStoredGameRequest): Promise<QueryAllStoredGameResponse>
  /** Queries a nextGame by index. */
  NextGame(request: QueryGetNextGameRequest): Promise<QueryGetNextGameResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  PlayerInfo(request: QueryGetPlayerInfoRequest): Promise<QueryGetPlayerInfoResponse> {
    const data = QueryGetPlayerInfoRequest.encode(request).finish()
    const promise = this.rpc.request('xavierlepretre.checkers.checkers.Query', 'PlayerInfo', data)
    return promise.then((data) => QueryGetPlayerInfoResponse.decode(new Reader(data)))
  }

  PlayerInfoAll(request: QueryAllPlayerInfoRequest): Promise<QueryAllPlayerInfoResponse> {
    const data = QueryAllPlayerInfoRequest.encode(request).finish()
    const promise = this.rpc.request('xavierlepretre.checkers.checkers.Query', 'PlayerInfoAll', data)
    return promise.then((data) => QueryAllPlayerInfoResponse.decode(new Reader(data)))
  }

  CanPlayMove(request: QueryCanPlayMoveRequest): Promise<QueryCanPlayMoveResponse> {
    const data = QueryCanPlayMoveRequest.encode(request).finish()
    const promise = this.rpc.request('xavierlepretre.checkers.checkers.Query', 'CanPlayMove', data)
    return promise.then((data) => QueryCanPlayMoveResponse.decode(new Reader(data)))
  }

  StoredGame(request: QueryGetStoredGameRequest): Promise<QueryGetStoredGameResponse> {
    const data = QueryGetStoredGameRequest.encode(request).finish()
    const promise = this.rpc.request('xavierlepretre.checkers.checkers.Query', 'StoredGame', data)
    return promise.then((data) => QueryGetStoredGameResponse.decode(new Reader(data)))
  }

  StoredGameAll(request: QueryAllStoredGameRequest): Promise<QueryAllStoredGameResponse> {
    const data = QueryAllStoredGameRequest.encode(request).finish()
    const promise = this.rpc.request('xavierlepretre.checkers.checkers.Query', 'StoredGameAll', data)
    return promise.then((data) => QueryAllStoredGameResponse.decode(new Reader(data)))
  }

  NextGame(request: QueryGetNextGameRequest): Promise<QueryGetNextGameResponse> {
    const data = QueryGetNextGameRequest.encode(request).finish()
    const promise = this.rpc.request('xavierlepretre.checkers.checkers.Query', 'NextGame', data)
    return promise.then((data) => QueryGetNextGameResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
}

declare var self: any | undefined
declare var window: any | undefined
var globalThis: any = (() => {
  if (typeof globalThis !== 'undefined') return globalThis
  if (typeof self !== 'undefined') return self
  if (typeof window !== 'undefined') return window
  if (typeof global !== 'undefined') return global
  throw 'Unable to locate global object'
})()

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER')
  }
  return long.toNumber()
}

if (util.Long !== Long) {
  util.Long = Long as any
  configure()
}
