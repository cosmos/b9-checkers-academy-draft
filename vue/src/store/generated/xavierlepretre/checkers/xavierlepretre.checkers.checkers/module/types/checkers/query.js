/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { StoredGame } from '../checkers/stored_game';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { NextGame } from '../checkers/next_game';
export const protobufPackage = 'xavierlepretre.checkers.checkers';
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
    StoredGame(request) {
        const data = QueryGetStoredGameRequest.encode(request).finish();
        const promise = this.rpc.request('xavierlepretre.checkers.checkers.Query', 'StoredGame', data);
        return promise.then((data) => QueryGetStoredGameResponse.decode(new Reader(data)));
    }
    StoredGameAll(request) {
        const data = QueryAllStoredGameRequest.encode(request).finish();
        const promise = this.rpc.request('xavierlepretre.checkers.checkers.Query', 'StoredGameAll', data);
        return promise.then((data) => QueryAllStoredGameResponse.decode(new Reader(data)));
    }
    NextGame(request) {
        const data = QueryGetNextGameRequest.encode(request).finish();
        const promise = this.rpc.request('xavierlepretre.checkers.checkers.Query', 'NextGame', data);
        return promise.then((data) => QueryGetNextGameResponse.decode(new Reader(data)));
    }
}
