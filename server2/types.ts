export interface LatestBlockStatus {
    height: number
}

export interface DbStatus {
    block: LatestBlockStatus
}

export interface PlayerInfo {
    gameIds: string[]
}

export interface PlayersInfo {
    [playerAddress: string]: PlayerInfo
}

export interface GameInfo {
    redAddress: string
    blackAddress: string
}

export interface GamesInfo {
    [gameId: string]: GameInfo
}

export interface DbType {
    status: DbStatus
    players: PlayersInfo
    games: GamesInfo
}
