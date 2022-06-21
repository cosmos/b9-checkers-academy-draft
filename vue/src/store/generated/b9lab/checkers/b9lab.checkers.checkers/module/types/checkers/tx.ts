/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'

export const protobufPackage = 'b9lab.checkers.checkers'

/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgPlayMove {
  creator: string
  idValue: string
  fromX: number
  fromY: number
  toX: number
  toY: number
}

export interface MsgPlayMoveResponse {
  idValue: string
  capturedX: number
  capturedY: number
  winner: string
}

export interface MsgCreateGame {
  creator: string
  red: string
  black: string
}

export interface MsgCreateGameResponse {
  idValue: string
}

const baseMsgPlayMove: object = { creator: '', idValue: '', fromX: 0, fromY: 0, toX: 0, toY: 0 }

export const MsgPlayMove = {
  encode(message: MsgPlayMove, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.idValue !== '') {
      writer.uint32(18).string(message.idValue)
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

  decode(input: Reader | Uint8Array, length?: number): MsgPlayMove {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgPlayMove } as MsgPlayMove
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.idValue = reader.string()
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

  fromJSON(object: any): MsgPlayMove {
    const message = { ...baseMsgPlayMove } as MsgPlayMove
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = String(object.idValue)
    } else {
      message.idValue = ''
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

  toJSON(message: MsgPlayMove): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.idValue !== undefined && (obj.idValue = message.idValue)
    message.fromX !== undefined && (obj.fromX = message.fromX)
    message.fromY !== undefined && (obj.fromY = message.fromY)
    message.toX !== undefined && (obj.toX = message.toX)
    message.toY !== undefined && (obj.toY = message.toY)
    return obj
  },

  fromPartial(object: DeepPartial<MsgPlayMove>): MsgPlayMove {
    const message = { ...baseMsgPlayMove } as MsgPlayMove
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = object.idValue
    } else {
      message.idValue = ''
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

const baseMsgPlayMoveResponse: object = { idValue: '', capturedX: 0, capturedY: 0, winner: '' }

export const MsgPlayMoveResponse = {
  encode(message: MsgPlayMoveResponse, writer: Writer = Writer.create()): Writer {
    if (message.idValue !== '') {
      writer.uint32(10).string(message.idValue)
    }
    if (message.capturedX !== 0) {
      writer.uint32(16).int64(message.capturedX)
    }
    if (message.capturedY !== 0) {
      writer.uint32(24).int64(message.capturedY)
    }
    if (message.winner !== '') {
      writer.uint32(34).string(message.winner)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPlayMoveResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgPlayMoveResponse } as MsgPlayMoveResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.idValue = reader.string()
          break
        case 2:
          message.capturedX = longToNumber(reader.int64() as Long)
          break
        case 3:
          message.capturedY = longToNumber(reader.int64() as Long)
          break
        case 4:
          message.winner = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgPlayMoveResponse {
    const message = { ...baseMsgPlayMoveResponse } as MsgPlayMoveResponse
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = String(object.idValue)
    } else {
      message.idValue = ''
    }
    if (object.capturedX !== undefined && object.capturedX !== null) {
      message.capturedX = Number(object.capturedX)
    } else {
      message.capturedX = 0
    }
    if (object.capturedY !== undefined && object.capturedY !== null) {
      message.capturedY = Number(object.capturedY)
    } else {
      message.capturedY = 0
    }
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = String(object.winner)
    } else {
      message.winner = ''
    }
    return message
  },

  toJSON(message: MsgPlayMoveResponse): unknown {
    const obj: any = {}
    message.idValue !== undefined && (obj.idValue = message.idValue)
    message.capturedX !== undefined && (obj.capturedX = message.capturedX)
    message.capturedY !== undefined && (obj.capturedY = message.capturedY)
    message.winner !== undefined && (obj.winner = message.winner)
    return obj
  },

  fromPartial(object: DeepPartial<MsgPlayMoveResponse>): MsgPlayMoveResponse {
    const message = { ...baseMsgPlayMoveResponse } as MsgPlayMoveResponse
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = object.idValue
    } else {
      message.idValue = ''
    }
    if (object.capturedX !== undefined && object.capturedX !== null) {
      message.capturedX = object.capturedX
    } else {
      message.capturedX = 0
    }
    if (object.capturedY !== undefined && object.capturedY !== null) {
      message.capturedY = object.capturedY
    } else {
      message.capturedY = 0
    }
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = object.winner
    } else {
      message.winner = ''
    }
    return message
  }
}

const baseMsgCreateGame: object = { creator: '', red: '', black: '' }

export const MsgCreateGame = {
  encode(message: MsgCreateGame, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.red !== '') {
      writer.uint32(18).string(message.red)
    }
    if (message.black !== '') {
      writer.uint32(26).string(message.black)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateGame } as MsgCreateGame
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.red = reader.string()
          break
        case 3:
          message.black = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgCreateGame {
    const message = { ...baseMsgCreateGame } as MsgCreateGame
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.red !== undefined && object.red !== null) {
      message.red = String(object.red)
    } else {
      message.red = ''
    }
    if (object.black !== undefined && object.black !== null) {
      message.black = String(object.black)
    } else {
      message.black = ''
    }
    return message
  },

  toJSON(message: MsgCreateGame): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.red !== undefined && (obj.red = message.red)
    message.black !== undefined && (obj.black = message.black)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateGame>): MsgCreateGame {
    const message = { ...baseMsgCreateGame } as MsgCreateGame
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.red !== undefined && object.red !== null) {
      message.red = object.red
    } else {
      message.red = ''
    }
    if (object.black !== undefined && object.black !== null) {
      message.black = object.black
    } else {
      message.black = ''
    }
    return message
  }
}

const baseMsgCreateGameResponse: object = { idValue: '' }

export const MsgCreateGameResponse = {
  encode(message: MsgCreateGameResponse, writer: Writer = Writer.create()): Writer {
    if (message.idValue !== '') {
      writer.uint32(10).string(message.idValue)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateGameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateGameResponse } as MsgCreateGameResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.idValue = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgCreateGameResponse {
    const message = { ...baseMsgCreateGameResponse } as MsgCreateGameResponse
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = String(object.idValue)
    } else {
      message.idValue = ''
    }
    return message
  },

  toJSON(message: MsgCreateGameResponse): unknown {
    const obj: any = {}
    message.idValue !== undefined && (obj.idValue = message.idValue)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateGameResponse>): MsgCreateGameResponse {
    const message = { ...baseMsgCreateGameResponse } as MsgCreateGameResponse
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = object.idValue
    } else {
      message.idValue = ''
    }
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  PlayMove(request: MsgPlayMove): Promise<MsgPlayMoveResponse>
  CreateGame(request: MsgCreateGame): Promise<MsgCreateGameResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  PlayMove(request: MsgPlayMove): Promise<MsgPlayMoveResponse> {
    const data = MsgPlayMove.encode(request).finish()
    const promise = this.rpc.request('b9lab.checkers.checkers.Msg', 'PlayMove', data)
    return promise.then((data) => MsgPlayMoveResponse.decode(new Reader(data)))
  }

  CreateGame(request: MsgCreateGame): Promise<MsgCreateGameResponse> {
    const data = MsgCreateGame.encode(request).finish()
    const promise = this.rpc.request('b9lab.checkers.checkers.Msg', 'CreateGame', data)
    return promise.then((data) => MsgCreateGameResponse.decode(new Reader(data)))
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
