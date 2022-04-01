import { createProtobufRpcClient, QueryClient } from "@cosmjs/stargate"
import { assert } from "@cosmjs/utils"
import { NextGame } from "../../types/generated/checkers/next_game"
import {
    QueryAllStoredGameResponse,
    QueryClientImpl,
    QueryGetStoredGameResponse,
} from "../../types/generated/checkers/query"
import { StoredGame } from "../../types/generated/checkers/stored_game"
import { PageResponse } from "../../types/generated/cosmos/base/query/v1beta1/pagination"

export interface AllStoredGameResponse {
    storedGames: StoredGame[]
    pagination?: PageResponse
}

export interface CheckersExtension {
    readonly checkers: {
        readonly getNextGame: () => Promise<NextGame>
        readonly getStoredGame: (index: string) => Promise<StoredGame | undefined>
        readonly getAllStoredGames: (
            key: Uint8Array,
            offset: Long,
            limit: Long,
            countTotal: boolean,
        ) => Promise<AllStoredGameResponse>
    }
}

export function setupCheckersExtension(base: QueryClient): CheckersExtension {
    const rpc = createProtobufRpcClient(base)
    // Use this service to get easy typed access to query methods
    // This cannot be used for proof verification
    const queryService = new QueryClientImpl(rpc)

    return {
        checkers: {
            getNextGame: async (): Promise<NextGame> => {
                const { NextGame } = await queryService.NextGame({})
                assert(NextGame)
                return NextGame
            },
            getStoredGame: async (index: string): Promise<StoredGame | undefined> => {
                const response: QueryGetStoredGameResponse = await queryService.StoredGame({
                    index: index,
                })
                return response.StoredGame
            },
            getAllStoredGames: async (
                key: Uint8Array,
                offset: Long,
                limit: Long,
                countTotal: boolean,
            ): Promise<AllStoredGameResponse> => {
                const response: QueryAllStoredGameResponse = await queryService.StoredGameAll({
                    pagination: {
                        key: key,
                        offset: offset,
                        limit: limit,
                        countTotal: countTotal,
                    },
                })
                return {
                    storedGames: response.StoredGame,
                    pagination: response.pagination,
                }
            },
        },
    }
}
