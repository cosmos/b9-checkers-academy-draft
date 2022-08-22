import { txClient, queryClient, MissingWalletError , registry} from './module'

import { Leaderboard } from "./module/types/checkers/leaderboard"
import { Params } from "./module/types/checkers/params"
import { PlayerInfo } from "./module/types/checkers/player_info"
import { StoredGame } from "./module/types/checkers/stored_game"
import { SystemInfo } from "./module/types/checkers/system_info"
import { WinningPlayer } from "./module/types/checkers/winning_player"


export { Leaderboard, Params, PlayerInfo, StoredGame, SystemInfo, WinningPlayer };

async function initTxClient(vuexGetters) {
	return await txClient(vuexGetters['common/wallet/signer'], {
		addr: vuexGetters['common/env/apiTendermint']
	})
}

async function initQueryClient(vuexGetters) {
	return await queryClient({
		addr: vuexGetters['common/env/apiCosmos']
	})
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

function getStructure(template) {
	let structure = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field: any = {}
		field.name = key
		field.type = typeof value
		structure.fields.push(field)
	}
	return structure
}

const getDefaultState = () => {
	return {
				Params: {},
				SystemInfo: {},
				StoredGame: {},
				StoredGameAll: {},
				CanPlayMove: {},
				PlayerInfo: {},
				PlayerInfoAll: {},
				Leaderboard: {},
				
				_Structure: {
						Leaderboard: getStructure(Leaderboard.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						PlayerInfo: getStructure(PlayerInfo.fromPartial({})),
						StoredGame: getStructure(StoredGame.fromPartial({})),
						SystemInfo: getStructure(SystemInfo.fromPartial({})),
						WinningPlayer: getStructure(WinningPlayer.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getSystemInfo: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.SystemInfo[JSON.stringify(params)] ?? {}
		},
				getStoredGame: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.StoredGame[JSON.stringify(params)] ?? {}
		},
				getStoredGameAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.StoredGameAll[JSON.stringify(params)] ?? {}
		},
				getCanPlayMove: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CanPlayMove[JSON.stringify(params)] ?? {}
		},
				getPlayerInfo: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PlayerInfo[JSON.stringify(params)] ?? {}
		},
				getPlayerInfoAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PlayerInfoAll[JSON.stringify(params)] ?? {}
		},
				getLeaderboard: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Leaderboard[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: b9lab.checkers.checkers initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QuerySystemInfo({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.querySystemInfo()).data
				
					
				commit('QUERY', { query: 'SystemInfo', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QuerySystemInfo', payload: { options: { all }, params: {...key},query }})
				return getters['getSystemInfo']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QuerySystemInfo API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryStoredGame({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryStoredGame( key.index)).data
				
					
				commit('QUERY', { query: 'StoredGame', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryStoredGame', payload: { options: { all }, params: {...key},query }})
				return getters['getStoredGame']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryStoredGame API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryStoredGameAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryStoredGameAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryStoredGameAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'StoredGameAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryStoredGameAll', payload: { options: { all }, params: {...key},query }})
				return getters['getStoredGameAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryStoredGameAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCanPlayMove({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryCanPlayMove( key.gameIndex,  key.player,  key.fromX,  key.fromY,  key.toX,  key.toY)).data
				
					
				commit('QUERY', { query: 'CanPlayMove', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCanPlayMove', payload: { options: { all }, params: {...key},query }})
				return getters['getCanPlayMove']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCanPlayMove API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPlayerInfo({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryPlayerInfo( key.index)).data
				
					
				commit('QUERY', { query: 'PlayerInfo', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPlayerInfo', payload: { options: { all }, params: {...key},query }})
				return getters['getPlayerInfo']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPlayerInfo API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPlayerInfoAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryPlayerInfoAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryPlayerInfoAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'PlayerInfoAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPlayerInfoAll', payload: { options: { all }, params: {...key},query }})
				return getters['getPlayerInfoAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPlayerInfoAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLeaderboard({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryLeaderboard()).data
				
					
				commit('QUERY', { query: 'Leaderboard', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLeaderboard', payload: { options: { all }, params: {...key},query }})
				return getters['getLeaderboard']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLeaderboard API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgPlayMove({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgPlayMove(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPlayMove:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgPlayMove:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateGame({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateGame(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateGame:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateGame:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRejectGame({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgRejectGame(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRejectGame:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRejectGame:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgPlayMove({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgPlayMove(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPlayMove:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgPlayMove:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateGame({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateGame(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateGame:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateGame:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRejectGame({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgRejectGame(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRejectGame:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRejectGame:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
