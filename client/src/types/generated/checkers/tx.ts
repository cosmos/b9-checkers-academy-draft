/* eslint-disable */
import Long from "long"
import _m0 from "protobufjs/minimal"

export const protobufPackage = "xavierlepretre.checkers.checkers"

/** Msg defines the Msg service. */
export interface Msg {}

export class MsgClientImpl implements Msg {
    private readonly rpc: Rpc
    constructor(rpc: Rpc) {
        this.rpc = rpc
    }
}

interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
}

if (_m0.util.Long !== Long) {
    _m0.util.Long = Long as any
    _m0.configure()
}
