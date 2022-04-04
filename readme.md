# checkers
**checkers** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

## Get started

```
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.network).

### Launch

To launch your blockchain live on multiple nodes, use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

### Web Frontend

Starport has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Starport front-end development](https://github.com/tendermint/vue).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.starport.network/xavierlepretre/checkers@latest! | sudo bash
```
`xavierlepretre/checkers` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Progressive feature branches

* `starport-start`
* `rules-added`, [diff](../../compare/starport-start..rules-added)
* `stored-game`, [diff](../../compare/rules-added..stored-game)
* `full-game-object`, [diff](../../compare/stored-game..full-game-object)
* `create-game-msg`, [diff](../../compare/full-game-object..create-game-msg)
* `create-game-handler`, [diff](../../compare/create-game-msg..create-game-handler)
* `play-move-msg`, [diff](../../compare/create-game-handler..play-move-msg)
* `play-move-handler`, [diff](../../compare/play-move-msg..play-move-handler)
* `two-events`, [diff](../../compare/play-move-handler..two-events)
* `reject-game-msg`, [diff](../../compare/two-events..reject-game-msg)
* `reject-game-handler`, [diff](../../compare/reject-game-msg..reject-game-handler)
* `game-fifo`, [diff](../../compare/reject-game-handler..game-fifo)
* `game-deadline`, [diff](../../compare/game-fifo..game-deadline)
* `game-winner`, [diff](../../compare/game-deadline..game-winner)
* `forfeit-game`, [diff](../../compare/game-winner..forfeit-game)
* `game-wager`, [diff](../../compare/forfeit-game..game-wager)
* `payment-winning`, [diff](../../compare/game-wager..payment-winning)
* `gas-meter`, [diff](../../compare/payment-winning..gas-meter)
* `can-play-move-query`, [diff](../../compare/gas-meter..can-play-move-query)
* `can-play-move-handler`, [diff](../../compare/can-play-move-query..can-play-move-handler)
* `wager-denomination`, [diff](../../compare/can-play-move-handler..wager-denomination)
* `cosmjs-elements`, [diff](../../compare/wager-denomination..cosmjs-elements)
* `cosmjs-messages`, [diff](../../compare/cosmjs-elements..cosmjs-messages)
* `player-info-object`, [diff](../../compare/cosmjs-messages..player-info-object)
* `player-info-handling`, [diff](../../compare/player-info-object..player-info-handling)
* `leaderboard-object`, [diff](../../compare/player-info-handling..leaderboard-object)
* `leaderboard-handling`, [diff](../../compare/leaderboard-object..leaderboard-handling)
* `genesis-migration`, [diff](../../compare/leaderboard-handling..genesis-migration)
* `server-indexing`, [diff](../../compare/genesis-migration..server-indexing)

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Starport Docs](https://docs.starport.network)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/W8trcGV)
