/* eslint-disable */
import { WinningPlayer } from '../checkers/winning_player'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'b9lab.checkers.checkers'

export interface Leaderboard {
  winners: WinningPlayer[]
}

const baseLeaderboard: object = {}

export const Leaderboard = {
  encode(message: Leaderboard, writer: Writer = Writer.create()): Writer {
    for (const v of message.winners) {
      WinningPlayer.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): Leaderboard {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseLeaderboard } as Leaderboard
    message.winners = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.winners.push(WinningPlayer.decode(reader, reader.uint32()))
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): Leaderboard {
    const message = { ...baseLeaderboard } as Leaderboard
    message.winners = []
    if (object.winners !== undefined && object.winners !== null) {
      for (const e of object.winners) {
        message.winners.push(WinningPlayer.fromJSON(e))
      }
    }
    return message
  },

  toJSON(message: Leaderboard): unknown {
    const obj: any = {}
    if (message.winners) {
      obj.winners = message.winners.map((e) => (e ? WinningPlayer.toJSON(e) : undefined))
    } else {
      obj.winners = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<Leaderboard>): Leaderboard {
    const message = { ...baseLeaderboard } as Leaderboard
    message.winners = []
    if (object.winners !== undefined && object.winners !== null) {
      for (const e of object.winners) {
        message.winners.push(WinningPlayer.fromPartial(e))
      }
    }
    return message
  }
}

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
