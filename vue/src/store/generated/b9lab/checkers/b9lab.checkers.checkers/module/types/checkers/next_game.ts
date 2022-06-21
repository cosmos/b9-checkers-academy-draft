/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'b9lab.checkers.checkers'

export interface NextGame {
  creator: string
  idValue: number
  fifoHead: string
  fifoTail: string
}

const baseNextGame: object = { creator: '', idValue: 0, fifoHead: '', fifoTail: '' }

export const NextGame = {
  encode(message: NextGame, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.idValue !== 0) {
      writer.uint32(16).uint64(message.idValue)
    }
    if (message.fifoHead !== '') {
      writer.uint32(26).string(message.fifoHead)
    }
    if (message.fifoTail !== '') {
      writer.uint32(34).string(message.fifoTail)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): NextGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseNextGame } as NextGame
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.idValue = longToNumber(reader.uint64() as Long)
          break
        case 3:
          message.fifoHead = reader.string()
          break
        case 4:
          message.fifoTail = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): NextGame {
    const message = { ...baseNextGame } as NextGame
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = Number(object.idValue)
    } else {
      message.idValue = 0
    }
    if (object.fifoHead !== undefined && object.fifoHead !== null) {
      message.fifoHead = String(object.fifoHead)
    } else {
      message.fifoHead = ''
    }
    if (object.fifoTail !== undefined && object.fifoTail !== null) {
      message.fifoTail = String(object.fifoTail)
    } else {
      message.fifoTail = ''
    }
    return message
  },

  toJSON(message: NextGame): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.idValue !== undefined && (obj.idValue = message.idValue)
    message.fifoHead !== undefined && (obj.fifoHead = message.fifoHead)
    message.fifoTail !== undefined && (obj.fifoTail = message.fifoTail)
    return obj
  },

  fromPartial(object: DeepPartial<NextGame>): NextGame {
    const message = { ...baseNextGame } as NextGame
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = object.idValue
    } else {
      message.idValue = 0
    }
    if (object.fifoHead !== undefined && object.fifoHead !== null) {
      message.fifoHead = object.fifoHead
    } else {
      message.fifoHead = ''
    }
    if (object.fifoTail !== undefined && object.fifoTail !== null) {
      message.fifoTail = object.fifoTail
    } else {
      message.fifoTail = ''
    }
    return message
  }
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
