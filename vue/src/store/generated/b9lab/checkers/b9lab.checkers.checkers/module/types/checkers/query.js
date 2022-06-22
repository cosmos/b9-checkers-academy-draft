/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal';
import * as Long from 'long';
import { Leaderboard } from '../checkers/leaderboard';
import { PlayerInfo } from '../checkers/player_info';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { StoredGame } from '../checkers/stored_game';
import { NextGame } from '../checkers/next_game';
export const protobufPackage = 'b9lab.checkers.checkers';
const baseQueryGetLeaderboardRequest = {};
export const QueryGetLeaderboardRequest = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetLeaderboardRequest };
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
    fromJSON(_) {
        const message = { ...baseQueryGetLeaderboardRequest };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseQueryGetLeaderboardRequest };
        return message;
    }
};
const baseQueryGetLeaderboardResponse = {};
export const QueryGetLeaderboardResponse = {
    encode(message, writer = Writer.create()) {
        if (message.Leaderboard !== undefined) {
            Leaderboard.encode(message.Leaderboard, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetLeaderboardResponse };
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
    fromJSON(object) {
        const message = { ...baseQueryGetLeaderboardResponse };
        if (object.Leaderboard !== undefined && object.Leaderboard !== null) {
            message.Leaderboard = Leaderboard.fromJSON(object.Leaderboard);
        }
        else {
            message.Leaderboard = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.Leaderboard !== undefined && (obj.Leaderboard = message.Leaderboard ? Leaderboard.toJSON(message.Leaderboard) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetLeaderboardResponse };
        if (object.Leaderboard !== undefined && object.Leaderboard !== null) {
            message.Leaderboard = Leaderboard.fromPartial(object.Leaderboard);
        }
        else {
            message.Leaderboard = undefined;
        }
        return message;
    }
};
const baseQueryGetPlayerInfoRequest = { index: '' };
export const QueryGetPlayerInfoRequest = {
    encode(message, writer = Writer.create()) {
        if (message.index !== '') {
            writer.uint32(10).string(message.index);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetPlayerInfoRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryGetPlayerInfoRequest };
        if (object.index !== undefined && object.index !== null) {
            message.index = String(object.index);
        }
        else {
            message.index = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.index !== undefined && (obj.index = message.index);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetPlayerInfoRequest };
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = '';
        }
        return message;
    }
};
const baseQueryGetPlayerInfoResponse = {};
export const QueryGetPlayerInfoResponse = {
    encode(message, writer = Writer.create()) {
        if (message.PlayerInfo !== undefined) {
            PlayerInfo.encode(message.PlayerInfo, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetPlayerInfoResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.PlayerInfo = PlayerInfo.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetPlayerInfoResponse };
        if (object.PlayerInfo !== undefined && object.PlayerInfo !== null) {
            message.PlayerInfo = PlayerInfo.fromJSON(object.PlayerInfo);
        }
        else {
            message.PlayerInfo = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.PlayerInfo !== undefined && (obj.PlayerInfo = message.PlayerInfo ? PlayerInfo.toJSON(message.PlayerInfo) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetPlayerInfoResponse };
        if (object.PlayerInfo !== undefined && object.PlayerInfo !== null) {
            message.PlayerInfo = PlayerInfo.fromPartial(object.PlayerInfo);
        }
        else {
            message.PlayerInfo = undefined;
        }
        return message;
    }
};
const baseQueryAllPlayerInfoRequest = {};
export const QueryAllPlayerInfoRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllPlayerInfoRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryAllPlayerInfoRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllPlayerInfoRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllPlayerInfoResponse = {};
export const QueryAllPlayerInfoResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.PlayerInfo) {
            PlayerInfo.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllPlayerInfoResponse };
        message.PlayerInfo = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.PlayerInfo.push(PlayerInfo.decode(reader, reader.uint32()));
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
    fromJSON(object) {
        const message = { ...baseQueryAllPlayerInfoResponse };
        message.PlayerInfo = [];
        if (object.PlayerInfo !== undefined && object.PlayerInfo !== null) {
            for (const e of object.PlayerInfo) {
                message.PlayerInfo.push(PlayerInfo.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.PlayerInfo) {
            obj.PlayerInfo = message.PlayerInfo.map((e) => (e ? PlayerInfo.toJSON(e) : undefined));
        }
        else {
            obj.PlayerInfo = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllPlayerInfoResponse };
        message.PlayerInfo = [];
        if (object.PlayerInfo !== undefined && object.PlayerInfo !== null) {
            for (const e of object.PlayerInfo) {
                message.PlayerInfo.push(PlayerInfo.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryCanPlayMoveRequest = { idValue: '', player: '', fromX: 0, fromY: 0, toX: 0, toY: 0 };
export const QueryCanPlayMoveRequest = {
    encode(message, writer = Writer.create()) {
        if (message.idValue !== '') {
            writer.uint32(10).string(message.idValue);
        }
        if (message.player !== '') {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryCanPlayMoveRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.idValue = reader.string();
                    break;
                case 2:
                    message.player = reader.string();
                    break;
                case 3:
                    message.fromX = longToNumber(reader.uint64());
                    break;
                case 4:
                    message.fromY = longToNumber(reader.uint64());
                    break;
                case 5:
                    message.toX = longToNumber(reader.uint64());
                    break;
                case 6:
                    message.toY = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryCanPlayMoveRequest };
        if (object.idValue !== undefined && object.idValue !== null) {
            message.idValue = String(object.idValue);
        }
        else {
            message.idValue = '';
        }
        if (object.player !== undefined && object.player !== null) {
            message.player = String(object.player);
        }
        else {
            message.player = '';
        }
        if (object.fromX !== undefined && object.fromX !== null) {
            message.fromX = Number(object.fromX);
        }
        else {
            message.fromX = 0;
        }
        if (object.fromY !== undefined && object.fromY !== null) {
            message.fromY = Number(object.fromY);
        }
        else {
            message.fromY = 0;
        }
        if (object.toX !== undefined && object.toX !== null) {
            message.toX = Number(object.toX);
        }
        else {
            message.toX = 0;
        }
        if (object.toY !== undefined && object.toY !== null) {
            message.toY = Number(object.toY);
        }
        else {
            message.toY = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.idValue !== undefined && (obj.idValue = message.idValue);
        message.player !== undefined && (obj.player = message.player);
        message.fromX !== undefined && (obj.fromX = message.fromX);
        message.fromY !== undefined && (obj.fromY = message.fromY);
        message.toX !== undefined && (obj.toX = message.toX);
        message.toY !== undefined && (obj.toY = message.toY);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryCanPlayMoveRequest };
        if (object.idValue !== undefined && object.idValue !== null) {
            message.idValue = object.idValue;
        }
        else {
            message.idValue = '';
        }
        if (object.player !== undefined && object.player !== null) {
            message.player = object.player;
        }
        else {
            message.player = '';
        }
        if (object.fromX !== undefined && object.fromX !== null) {
            message.fromX = object.fromX;
        }
        else {
            message.fromX = 0;
        }
        if (object.fromY !== undefined && object.fromY !== null) {
            message.fromY = object.fromY;
        }
        else {
            message.fromY = 0;
        }
        if (object.toX !== undefined && object.toX !== null) {
            message.toX = object.toX;
        }
        else {
            message.toX = 0;
        }
        if (object.toY !== undefined && object.toY !== null) {
            message.toY = object.toY;
        }
        else {
            message.toY = 0;
        }
        return message;
    }
};
const baseQueryCanPlayMoveResponse = { possible: false, reason: '' };
export const QueryCanPlayMoveResponse = {
    encode(message, writer = Writer.create()) {
        if (message.possible === true) {
            writer.uint32(8).bool(message.possible);
        }
        if (message.reason !== '') {
            writer.uint32(18).string(message.reason);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryCanPlayMoveResponse };
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
    fromJSON(object) {
        const message = { ...baseQueryCanPlayMoveResponse };
        if (object.possible !== undefined && object.possible !== null) {
            message.possible = Boolean(object.possible);
        }
        else {
            message.possible = false;
        }
        if (object.reason !== undefined && object.reason !== null) {
            message.reason = String(object.reason);
        }
        else {
            message.reason = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.possible !== undefined && (obj.possible = message.possible);
        message.reason !== undefined && (obj.reason = message.reason);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryCanPlayMoveResponse };
        if (object.possible !== undefined && object.possible !== null) {
            message.possible = object.possible;
        }
        else {
            message.possible = false;
        }
        if (object.reason !== undefined && object.reason !== null) {
            message.reason = object.reason;
        }
        else {
            message.reason = '';
        }
        return message;
    }
};
const baseQueryGetStoredGameRequest = { index: '' };
export const QueryGetStoredGameRequest = {
    encode(message, writer = Writer.create()) {
        if (message.index !== '') {
            writer.uint32(10).string(message.index);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetStoredGameRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryGetStoredGameRequest };
        if (object.index !== undefined && object.index !== null) {
            message.index = String(object.index);
        }
        else {
            message.index = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.index !== undefined && (obj.index = message.index);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetStoredGameRequest };
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = '';
        }
        return message;
    }
};
const baseQueryGetStoredGameResponse = {};
export const QueryGetStoredGameResponse = {
    encode(message, writer = Writer.create()) {
        if (message.StoredGame !== undefined) {
            StoredGame.encode(message.StoredGame, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetStoredGameResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.StoredGame = StoredGame.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetStoredGameResponse };
        if (object.StoredGame !== undefined && object.StoredGame !== null) {
            message.StoredGame = StoredGame.fromJSON(object.StoredGame);
        }
        else {
            message.StoredGame = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.StoredGame !== undefined && (obj.StoredGame = message.StoredGame ? StoredGame.toJSON(message.StoredGame) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetStoredGameResponse };
        if (object.StoredGame !== undefined && object.StoredGame !== null) {
            message.StoredGame = StoredGame.fromPartial(object.StoredGame);
        }
        else {
            message.StoredGame = undefined;
        }
        return message;
    }
};
const baseQueryAllStoredGameRequest = {};
export const QueryAllStoredGameRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllStoredGameRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryAllStoredGameRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllStoredGameRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllStoredGameResponse = {};
export const QueryAllStoredGameResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.StoredGame) {
            StoredGame.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllStoredGameResponse };
        message.StoredGame = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.StoredGame.push(StoredGame.decode(reader, reader.uint32()));
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
    fromJSON(object) {
        const message = { ...baseQueryAllStoredGameResponse };
        message.StoredGame = [];
        if (object.StoredGame !== undefined && object.StoredGame !== null) {
            for (const e of object.StoredGame) {
                message.StoredGame.push(StoredGame.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.StoredGame) {
            obj.StoredGame = message.StoredGame.map((e) => (e ? StoredGame.toJSON(e) : undefined));
        }
        else {
            obj.StoredGame = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllStoredGameResponse };
        message.StoredGame = [];
        if (object.StoredGame !== undefined && object.StoredGame !== null) {
            for (const e of object.StoredGame) {
                message.StoredGame.push(StoredGame.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryGetNextGameRequest = {};
export const QueryGetNextGameRequest = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetNextGameRequest };
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
    fromJSON(_) {
        const message = { ...baseQueryGetNextGameRequest };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseQueryGetNextGameRequest };
        return message;
    }
};
const baseQueryGetNextGameResponse = {};
export const QueryGetNextGameResponse = {
    encode(message, writer = Writer.create()) {
        if (message.NextGame !== undefined) {
            NextGame.encode(message.NextGame, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetNextGameResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.NextGame = NextGame.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetNextGameResponse };
        if (object.NextGame !== undefined && object.NextGame !== null) {
            message.NextGame = NextGame.fromJSON(object.NextGame);
        }
        else {
            message.NextGame = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.NextGame !== undefined && (obj.NextGame = message.NextGame ? NextGame.toJSON(message.NextGame) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetNextGameResponse };
        if (object.NextGame !== undefined && object.NextGame !== null) {
            message.NextGame = NextGame.fromPartial(object.NextGame);
        }
        else {
            message.NextGame = undefined;
        }
        return message;
    }
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    Leaderboard(request) {
        const data = QueryGetLeaderboardRequest.encode(request).finish();
        const promise = this.rpc.request('b9lab.checkers.checkers.Query', 'Leaderboard', data);
        return promise.then((data) => QueryGetLeaderboardResponse.decode(new Reader(data)));
    }
    PlayerInfo(request) {
        const data = QueryGetPlayerInfoRequest.encode(request).finish();
        const promise = this.rpc.request('b9lab.checkers.checkers.Query', 'PlayerInfo', data);
        return promise.then((data) => QueryGetPlayerInfoResponse.decode(new Reader(data)));
    }
    PlayerInfoAll(request) {
        const data = QueryAllPlayerInfoRequest.encode(request).finish();
        const promise = this.rpc.request('b9lab.checkers.checkers.Query', 'PlayerInfoAll', data);
        return promise.then((data) => QueryAllPlayerInfoResponse.decode(new Reader(data)));
    }
    CanPlayMove(request) {
        const data = QueryCanPlayMoveRequest.encode(request).finish();
        const promise = this.rpc.request('b9lab.checkers.checkers.Query', 'CanPlayMove', data);
        return promise.then((data) => QueryCanPlayMoveResponse.decode(new Reader(data)));
    }
    StoredGame(request) {
        const data = QueryGetStoredGameRequest.encode(request).finish();
        const promise = this.rpc.request('b9lab.checkers.checkers.Query', 'StoredGame', data);
        return promise.then((data) => QueryGetStoredGameResponse.decode(new Reader(data)));
    }
    StoredGameAll(request) {
        const data = QueryAllStoredGameRequest.encode(request).finish();
        const promise = this.rpc.request('b9lab.checkers.checkers.Query', 'StoredGameAll', data);
        return promise.then((data) => QueryAllStoredGameResponse.decode(new Reader(data)));
    }
    NextGame(request) {
        const data = QueryGetNextGameRequest.encode(request).finish();
        const promise = this.rpc.request('b9lab.checkers.checkers.Query', 'NextGame', data);
        return promise.then((data) => QueryGetNextGameResponse.decode(new Reader(data)));
    }
}
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
