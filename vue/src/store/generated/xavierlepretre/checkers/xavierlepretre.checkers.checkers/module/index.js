// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgPlayMove } from "./types/checkers/tx";
import { MsgRejectGame } from "./types/checkers/tx";
import { MsgCreateGame } from "./types/checkers/tx";
const types = [
    ["/xavierlepretre.checkers.checkers.MsgPlayMove", MsgPlayMove],
    ["/xavierlepretre.checkers.checkers.MsgRejectGame", MsgRejectGame],
    ["/xavierlepretre.checkers.checkers.MsgCreateGame", MsgCreateGame],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgPlayMove: (data) => ({ typeUrl: "/xavierlepretre.checkers.checkers.MsgPlayMove", value: data }),
        msgRejectGame: (data) => ({ typeUrl: "/xavierlepretre.checkers.checkers.MsgRejectGame", value: data }),
        msgCreateGame: (data) => ({ typeUrl: "/xavierlepretre.checkers.checkers.MsgCreateGame", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
