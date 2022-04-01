/* eslint-disable */
import Long from "long"
import _m0 from "protobufjs/minimal"

export const protobufPackage = "xavierlepretre.checkers.checkers"

export interface NextGame {
    creator: string
    idValue: Long
    fifoHead: string
    fifoTail: string
}

function createBaseNextGame(): NextGame {
    return { creator: "", idValue: Long.UZERO, fifoHead: "", fifoTail: "" }
}

export const NextGame = {
    encode(message: NextGame, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
        if (message.creator !== "") {
            writer.uint32(10).string(message.creator)
        }
        if (!message.idValue.isZero()) {
            writer.uint32(16).uint64(message.idValue)
        }
        if (message.fifoHead !== "") {
            writer.uint32(26).string(message.fifoHead)
        }
        if (message.fifoTail !== "") {
            writer.uint32(34).string(message.fifoTail)
        }
        return writer
    },

    decode(input: _m0.Reader | Uint8Array, length?: number): NextGame {
        const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input)
        let end = length === undefined ? reader.len : reader.pos + length
        const message = createBaseNextGame()
        while (reader.pos < end) {
            const tag = reader.uint32()
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string()
                    break
                case 2:
                    message.idValue = reader.uint64() as Long
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
        return {
            creator: isSet(object.creator) ? String(object.creator) : "",
            idValue: isSet(object.idValue) ? Long.fromString(object.idValue) : Long.UZERO,
            fifoHead: isSet(object.fifoHead) ? String(object.fifoHead) : "",
            fifoTail: isSet(object.fifoTail) ? String(object.fifoTail) : "",
        }
    },

    toJSON(message: NextGame): unknown {
        const obj: any = {}
        message.creator !== undefined && (obj.creator = message.creator)
        message.idValue !== undefined && (obj.idValue = (message.idValue || Long.UZERO).toString())
        message.fifoHead !== undefined && (obj.fifoHead = message.fifoHead)
        message.fifoTail !== undefined && (obj.fifoTail = message.fifoTail)
        return obj
    },

    fromPartial<I extends Exact<DeepPartial<NextGame>, I>>(object: I): NextGame {
        const message = createBaseNextGame()
        message.creator = object.creator ?? ""
        message.idValue =
            object.idValue !== undefined && object.idValue !== null
                ? Long.fromValue(object.idValue)
                : Long.UZERO
        message.fifoHead = object.fifoHead ?? ""
        message.fifoTail = object.fifoTail ?? ""
        return message
    },
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined

export type DeepPartial<T> = T extends Builtin
    ? T
    : T extends Long
    ? string | number | Long
    : T extends Array<infer U>
    ? Array<DeepPartial<U>>
    : T extends ReadonlyArray<infer U>
    ? ReadonlyArray<DeepPartial<U>>
    : T extends {}
    ? { [K in keyof T]?: DeepPartial<T[K]> }
    : Partial<T>

type KeysOfUnion<T> = T extends T ? keyof T : never
export type Exact<P, I extends P> = P extends Builtin
    ? P
    : P & { [K in keyof P]: Exact<P[K], I[K]> } & Record<Exclude<keyof I, KeysOfUnion<P>>, never>

if (_m0.util.Long !== Long) {
    _m0.util.Long = Long as any
    _m0.configure()
}

function isSet(value: any): boolean {
    return value !== null && value !== undefined
}
