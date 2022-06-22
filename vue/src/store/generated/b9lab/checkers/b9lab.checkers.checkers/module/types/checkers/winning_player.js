/* eslint-disable */
import * as Long from 'long';
import { util, configure, Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'b9lab.checkers.checkers';
const baseWinningPlayer = { playerAddress: '', wonCount: 0, dateAdded: '' };
export const WinningPlayer = {
    encode(message, writer = Writer.create()) {
        if (message.playerAddress !== '') {
            writer.uint32(10).string(message.playerAddress);
        }
        if (message.wonCount !== 0) {
            writer.uint32(16).uint64(message.wonCount);
        }
        if (message.dateAdded !== '') {
            writer.uint32(26).string(message.dateAdded);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseWinningPlayer };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.playerAddress = reader.string();
                    break;
                case 2:
                    message.wonCount = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.dateAdded = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseWinningPlayer };
        if (object.playerAddress !== undefined && object.playerAddress !== null) {
            message.playerAddress = String(object.playerAddress);
        }
        else {
            message.playerAddress = '';
        }
        if (object.wonCount !== undefined && object.wonCount !== null) {
            message.wonCount = Number(object.wonCount);
        }
        else {
            message.wonCount = 0;
        }
        if (object.dateAdded !== undefined && object.dateAdded !== null) {
            message.dateAdded = String(object.dateAdded);
        }
        else {
            message.dateAdded = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.playerAddress !== undefined && (obj.playerAddress = message.playerAddress);
        message.wonCount !== undefined && (obj.wonCount = message.wonCount);
        message.dateAdded !== undefined && (obj.dateAdded = message.dateAdded);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseWinningPlayer };
        if (object.playerAddress !== undefined && object.playerAddress !== null) {
            message.playerAddress = object.playerAddress;
        }
        else {
            message.playerAddress = '';
        }
        if (object.wonCount !== undefined && object.wonCount !== null) {
            message.wonCount = object.wonCount;
        }
        else {
            message.wonCount = 0;
        }
        if (object.dateAdded !== undefined && object.dateAdded !== null) {
            message.dateAdded = object.dateAdded;
        }
        else {
            message.dateAdded = '';
        }
        return message;
    }
};
var globalThis = (() => {
    if (typeof globalThis !== 'undefined')
        return globalThis;
    if (typeof self !== 'undefined')
        return self;
    if (typeof window !== 'undefined')
        return window;
    if (typeof global !== 'undefined')
        return global;
    throw 'Unable to locate global object';
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER');
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
