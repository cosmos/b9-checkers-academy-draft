/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'xavierlepretre.checkers.checkers';
const baseStoredGame = {
    creator: '',
    index: '',
    game: '',
    turn: '',
    red: '',
    black: '',
    moveCount: '',
    beforeId: '',
    afterId: '',
    deadline: '',
    winner: ''
};
export const StoredGame = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.index !== '') {
            writer.uint32(18).string(message.index);
        }
        if (message.game !== '') {
            writer.uint32(26).string(message.game);
        }
        if (message.turn !== '') {
            writer.uint32(34).string(message.turn);
        }
        if (message.red !== '') {
            writer.uint32(42).string(message.red);
        }
        if (message.black !== '') {
            writer.uint32(50).string(message.black);
        }
        if (message.moveCount !== '') {
            writer.uint32(58).string(message.moveCount);
        }
        if (message.beforeId !== '') {
            writer.uint32(66).string(message.beforeId);
        }
        if (message.afterId !== '') {
            writer.uint32(74).string(message.afterId);
        }
        if (message.deadline !== '') {
            writer.uint32(82).string(message.deadline);
        }
        if (message.winner !== '') {
            writer.uint32(90).string(message.winner);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseStoredGame };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.index = reader.string();
                    break;
                case 3:
                    message.game = reader.string();
                    break;
                case 4:
                    message.turn = reader.string();
                    break;
                case 5:
                    message.red = reader.string();
                    break;
                case 6:
                    message.black = reader.string();
                    break;
                case 7:
                    message.moveCount = reader.string();
                    break;
                case 8:
                    message.beforeId = reader.string();
                    break;
                case 9:
                    message.afterId = reader.string();
                    break;
                case 10:
                    message.deadline = reader.string();
                    break;
                case 11:
                    message.winner = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseStoredGame };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.index !== undefined && object.index !== null) {
            message.index = String(object.index);
        }
        else {
            message.index = '';
        }
        if (object.game !== undefined && object.game !== null) {
            message.game = String(object.game);
        }
        else {
            message.game = '';
        }
        if (object.turn !== undefined && object.turn !== null) {
            message.turn = String(object.turn);
        }
        else {
            message.turn = '';
        }
        if (object.red !== undefined && object.red !== null) {
            message.red = String(object.red);
        }
        else {
            message.red = '';
        }
        if (object.black !== undefined && object.black !== null) {
            message.black = String(object.black);
        }
        else {
            message.black = '';
        }
        if (object.moveCount !== undefined && object.moveCount !== null) {
            message.moveCount = String(object.moveCount);
        }
        else {
            message.moveCount = '';
        }
        if (object.beforeId !== undefined && object.beforeId !== null) {
            message.beforeId = String(object.beforeId);
        }
        else {
            message.beforeId = '';
        }
        if (object.afterId !== undefined && object.afterId !== null) {
            message.afterId = String(object.afterId);
        }
        else {
            message.afterId = '';
        }
        if (object.deadline !== undefined && object.deadline !== null) {
            message.deadline = String(object.deadline);
        }
        else {
            message.deadline = '';
        }
        if (object.winner !== undefined && object.winner !== null) {
            message.winner = String(object.winner);
        }
        else {
            message.winner = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.index !== undefined && (obj.index = message.index);
        message.game !== undefined && (obj.game = message.game);
        message.turn !== undefined && (obj.turn = message.turn);
        message.red !== undefined && (obj.red = message.red);
        message.black !== undefined && (obj.black = message.black);
        message.moveCount !== undefined && (obj.moveCount = message.moveCount);
        message.beforeId !== undefined && (obj.beforeId = message.beforeId);
        message.afterId !== undefined && (obj.afterId = message.afterId);
        message.deadline !== undefined && (obj.deadline = message.deadline);
        message.winner !== undefined && (obj.winner = message.winner);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseStoredGame };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = '';
        }
        if (object.game !== undefined && object.game !== null) {
            message.game = object.game;
        }
        else {
            message.game = '';
        }
        if (object.turn !== undefined && object.turn !== null) {
            message.turn = object.turn;
        }
        else {
            message.turn = '';
        }
        if (object.red !== undefined && object.red !== null) {
            message.red = object.red;
        }
        else {
            message.red = '';
        }
        if (object.black !== undefined && object.black !== null) {
            message.black = object.black;
        }
        else {
            message.black = '';
        }
        if (object.moveCount !== undefined && object.moveCount !== null) {
            message.moveCount = object.moveCount;
        }
        else {
            message.moveCount = '';
        }
        if (object.beforeId !== undefined && object.beforeId !== null) {
            message.beforeId = object.beforeId;
        }
        else {
            message.beforeId = '';
        }
        if (object.afterId !== undefined && object.afterId !== null) {
            message.afterId = object.afterId;
        }
        else {
            message.afterId = '';
        }
        if (object.deadline !== undefined && object.deadline !== null) {
            message.deadline = object.deadline;
        }
        else {
            message.deadline = '';
        }
        if (object.winner !== undefined && object.winner !== null) {
            message.winner = object.winner;
        }
        else {
            message.winner = '';
        }
        return message;
    }
};
