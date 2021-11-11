/* eslint-disable */
import { WinningPlayer } from '../checkers/winning_player';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'xavierlepretre.checkers.checkers';
const baseLeaderboard = {};
export const Leaderboard = {
    encode(message, writer = Writer.create()) {
        for (const v of message.winners) {
            WinningPlayer.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseLeaderboard };
        message.winners = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.winners.push(WinningPlayer.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseLeaderboard };
        message.winners = [];
        if (object.winners !== undefined && object.winners !== null) {
            for (const e of object.winners) {
                message.winners.push(WinningPlayer.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.winners) {
            obj.winners = message.winners.map((e) => (e ? WinningPlayer.toJSON(e) : undefined));
        }
        else {
            obj.winners = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseLeaderboard };
        message.winners = [];
        if (object.winners !== undefined && object.winners !== null) {
            for (const e of object.winners) {
                message.winners.push(WinningPlayer.fromPartial(e));
            }
        }
        return message;
    }
};
