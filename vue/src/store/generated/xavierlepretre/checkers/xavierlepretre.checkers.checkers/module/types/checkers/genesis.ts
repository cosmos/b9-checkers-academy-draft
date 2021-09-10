/* eslint-disable */
import { StoredGame } from '../checkers/stored_game'
import { NextGame } from '../checkers/next_game'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'xavierlepretre.checkers.checkers'

/** GenesisState defines the checkers module's genesis state. */
export interface GenesisState {
  /** this line is used by starport scaffolding # genesis/proto/state */
  storedGameList: StoredGame[]
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  nextGame: NextGame | undefined
}

const baseGenesisState: object = {}

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.storedGameList) {
      StoredGame.encode(v!, writer.uint32(18).fork()).ldelim()
    }
    if (message.nextGame !== undefined) {
      NextGame.encode(message.nextGame, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    message.storedGameList = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 2:
          message.storedGameList.push(StoredGame.decode(reader, reader.uint32()))
          break
        case 1:
          message.nextGame = NextGame.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.storedGameList = []
    if (object.storedGameList !== undefined && object.storedGameList !== null) {
      for (const e of object.storedGameList) {
        message.storedGameList.push(StoredGame.fromJSON(e))
      }
    }
    if (object.nextGame !== undefined && object.nextGame !== null) {
      message.nextGame = NextGame.fromJSON(object.nextGame)
    } else {
      message.nextGame = undefined
    }
    return message
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {}
    if (message.storedGameList) {
      obj.storedGameList = message.storedGameList.map((e) => (e ? StoredGame.toJSON(e) : undefined))
    } else {
      obj.storedGameList = []
    }
    message.nextGame !== undefined && (obj.nextGame = message.nextGame ? NextGame.toJSON(message.nextGame) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.storedGameList = []
    if (object.storedGameList !== undefined && object.storedGameList !== null) {
      for (const e of object.storedGameList) {
        message.storedGameList.push(StoredGame.fromPartial(e))
      }
    }
    if (object.nextGame !== undefined && object.nextGame !== null) {
      message.nextGame = NextGame.fromPartial(object.nextGame)
    } else {
      message.nextGame = undefined
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
