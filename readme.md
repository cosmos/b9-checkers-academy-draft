# checkers

**checkers** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

Versions used are:

* Go: 1.16.15
* Ignite (formerly Starport): 0.17.3
* Cosmos SDK: v0.42.6
* NodeJs: 16.x

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
curl https://get.starport.network/b9lab/checkers@latest! | sudo bash
```
`b9lab/checkers` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Submodules

From the `cosmjs-elements` branch onward, there is a submodule linking the `client` folder to the [GUI repository](https://github.com/cosmos/academy-checkers-ui).

* To clone this repository and checkout the submodule at the same time (this is optional), run
    
    ```sh
    $ git clone --recurse-submodules THIS_REPO
    ```

    Or `--recursive` if your Git version is `< 2.13`.

* If you want to checkout the submodule at a later date, run:

    ```sh
    $ git submodule update --init
    ```

### Make updates

The submodule is checked at a specific commit. If you are a maintainer of this repository and want to change the submodule's commit, then you have to commit this new information. To do this:

```sh
# Go to client
$ cd client
# Checkout the version you want
$ git checkout xxyy
# Return to the main repository
$ cd ..
# Add the submodule to git's list
$ git add client
```

From there, proceed as you usually proceed to `commit` and `push`.

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
* `player-info-object`, [diff](../../compare/cosmjs-elements..player-info-object)
* `player-info-handling`, [diff](../../compare/player-info-object..player-info-handling)
* `leaderboard-object`, [diff](../../compare/player-info-handling..leaderboard-object)
* `leaderboard-handling`, [diff](../../compare/leaderboard-object..leaderboard-handling)
* `genesis-migration`, [diff](../../compare/leaderboard-handling..genesis-migration)

## Version 1 progressive tags

Versions used in version 1 are:

* Go: 1.16.15
* Ignite (formerly Starport): 0.17.3
* Cosmos SDK: v0.42.6
* NodeJs: 16.x

Tags:

* `v1-starport-start`
* `v1-rules-added`, [diff](../../compare/v1-starport-start..v1-rules-added)
* `v1-stored-game`, [diff](../../compare/v1-rules-added..v1-stored-game)
* `v1-full-game-object`, [diff](../../compare/v1-stored-game..v1-full-game-object)
* `v1-create-game-msg`, [diff](../../compare/v1-full-game-object..v1-create-game-msg)
* `v1-create-game-handler`, [diff](../../compare/v1-create-game-msg..v1-create-game-handler)
* `v1-play-move-msg`, [diff](../../compare/v1-create-game-handler..v1-play-move-msg)
* `v1-play-move-handler`, [diff](../../compare/v1-play-move-msg..v1-play-move-handler)
* `v1-two-events`, [diff](../../compare/v1-play-move-handler..v1-two-events)
* `v1-reject-game-msg`, [diff](../../compare/v1-two-events..v1-reject-game-msg)
* `v1-reject-game-handler`, [diff](../../compare/v1-reject-game-msg..v1-reject-game-handler)
* `v1-game-fifo`, [diff](../../compare/v1-reject-game-handler..v1-game-fifo)
* `v1-game-deadline`, [diff](../../compare/v1-game-fifo..v1-game-deadline)
* `v1-game-winner`, [diff](../../compare/v1-game-deadline..v1-game-winner)
* `v1-forfeit-game`, [diff](../../compare/v1-game-winner..v1-forfeit-game)
* `v1-game-wager`, [diff](../../compare/v1-forfeit-game..v1-game-wager)
* `v1-payment-winning`, [diff](../../compare/v1-game-wager..v1-payment-winning)
* `v1-gas-meter`, [diff](../../compare/v1-payment-winning..v1-gas-meter)
* `v1-can-play-move-query`, [diff](../../compare/v1-gas-meter..v1-can-play-move-query)
* `v1-can-play-move-handler`, [diff](../../compare/v1-can-play-move-query..v1-can-play-move-handler)
* `v1-wager-denomination`, [diff](../../compare/v1-can-play-move-handler..v1-wager-denomination)
* `v1-cosmjs-elements`, [diff](../../compare/v1-wager-denomination..v1-cosmjs-elements)
* `v1-player-info-object`, [diff](../../compare/v1-cosmjs-elements..v1-player-info-object)
* `v1-player-info-handling`, [diff](../../compare/v1-player-info-object..v1-player-info-handling)
* `v1-leaderboard-object`, [diff](../../compare/v1-player-info-handling..v1-leaderboard-object)
* `v1-leaderboard-handling`, [diff](../../compare/v1-leaderboard-object..v1-leaderboard-handling)
* `v1-genesis-migration`, [diff](../../compare/v1-leaderboard-handling..v1-genesis-migration)

## Use Docker

The Docker file has been prepared so that it allows you to run this project on any machine that supports Docker.

* Create the image:
  
    ```sh
    $ docker build -f Dockerfile-ubuntu . -t checkers_i
    ```

* Run a scaffold command in a throwaway container:

    ```sh
    $ docker run --rm -it -v $(pwd):/home/checkers checkers_i ignite scaffold ...
    ```

* Build a reusable container:

    ```sh
    $ docker create --name checkers -i -v $(pwd):/home/checkers -p 1317:1317 -p 4500:4500 -p 26657:26657 checkers_i
    $ docker start checkers
    ```

* Serve a chain and do actions on the reusable container:

    ```sh
    $ docker exec -it checkers ignite chain serve --reset-once
    ```

    And in another shell:

    ```sh
    $ docker exec -it checkers checkersd query checkers list-stored-game
    ```

    Or if you prefer to connect to the container and stay in it:

    ```sh
    $ docker exec -it checkers bash
    ```

    From where you can do:

    ```sh
    $ checkersd query checkers list-stored-game
    ```

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Starport Docs](https://docs.starport.network)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/W8trcGV)
