# Checkers tutorial project

This is a companion project of the Cosmos SDK tutorials. Its object is to show various features of the Cosmos SDK and of Ignite, along with the progression of the code as elements and features are added.

The progression of the code is demonstrated via the help of branches and diffs.

## Build steps taken

All the build steps were run inside the Docker container.

```sh
$ ignite scaffold chain github.com/b9lab/checkers
```

## Progressive feature branches

Versions used here are:

* Go: 1.18.3
* Ignite (formerly Starport): 0.22.1
* Cosmos SDK: v0.45.4
* NodeJs: 18.x
* Mockgen: 1.6.0
* Protoc: 21.7

Branches:

* [`v2-ignite-start`](../../tree/v2-ignite-start)
* [`v2-rules-added`](../../tree/v2-rules-added), [diff](../../compare/v2-ignite-start..v2-rules-added)
* [`v2-stored-game`](../../tree/v2-stored-game), [diff](../../compare/v2-rules-added..v2-stored-game)
* [`v2-full-game-object`](../../tree/v2-full-game-object), [diff](../../compare/v2-stored-game..v2-full-game-object)
* [`v2-create-game-msg`](../../tree/v2-create-game-msg), [diff](../../compare/v2-full-game-object..v2-create-game-msg)
* [`v2-create-game-handler`](../../tree/v2-create-game-handler), [diff](../../compare/v2-create-game-msg..v2-create-game-handler)
* [`v2-play-move-msg`](../../tree/v2-play-move-msg), [diff](../../compare/v2-create-game-handler..v2-play-move-msg)
* [`v2-play-move-handler`](../../tree/v2-play-move-handler), [diff](../../compare/v2-play-move-msg..v2-play-move-handler)
* [`v2-two-events`](../../tree/v2-two-events), [diff](../../compare/v2-play-move-handler..v2-two-events)
* [`v2-game-winner`](../../tree/v2-game-winner), [diff](../../compare/v2-two-events..v2-game-winner)
* [`v2-game-deadline`](../../tree/v2-game-deadline), [diff](../../compare/v2-game-winner..v2-game-deadline)
* [`v2-move-count`](../../tree/v2-move-count), [diff](../../compare/v2-game-deadline..v2-move-count)
* [`v2-fifo-fields`](../../tree/v2-fifo-fields), [diff](../../compare/v2-move-count..v2-fifo-fields)
* [`v2-game-fifo`](../../tree/v2-game-fifo), [diff](../../compare/v2-fifo-fields..v2-game-fifo)
* [`v2-forfeit-game`](../../tree/v2-forfeit-game), [diff](../../compare/v2-game-fifo..v2-forfeit-game)
* [`v2-game-wager`](../../tree/v2-game-wager), [diff](../../compare/v2-forfeit-game..v2-game-wager)
* [`v2-payment-winning`](../../tree/v2-payment-winning), [diff](../../compare/v2-game-wager..v2-payment-winning)
* [`v2-integration-tests`](../../tree/v2-integration-tests), [diff](../../compare/v2-payment-winning..v2-integration-tests)
* [`v2-gas-meter`](../../tree/v2-gas-meter), [diff](../../compare/v2-integration-tests..v2-gas-meter)
* [`v2-can-play-move-query`](../../tree/v2-can-play-move-query), [diff](../../compare/v2-gas-meter..v2-can-play-move-query)
* [`v2-can-play-move-handler`](../../tree/v2-can-play-move-handler), [diff](../../compare/v2-can-play-move-query..v2-can-play-move-handler)
* [`v2-wager-denomination`](../../tree/v2-wager-denomination), [diff](../../compare/v2-can-play-move-handler..v2-wager-denomination)
* [`v2-cosmjs-elements`](../../tree/v2-cosmjs-elements), [diff](../../compare/v2-wager-denomination..v2-cosmjs-elements)
* [`v2-run-prod`](../../tree/v2-run-prod), [diff](../../compare/v2-cosmjs-elements..v2-run-prod)
* [`v2-player-info-object`](../../tree/v2-player-info-object), [diff](../../compare/v2-run-prod..v2-player-info-object)
* [`v2-player-info-handling`](../../tree/v2-player-info-handling), [diff](../../compare/v2-player-info-object..v2-player-info-handling)
* [`v2-player-info-migration`](../../tree/v2-player-info-migration), [diff](../../compare/v2-player-info-handling..v2-player-info-migration)

## Version 1 progressive tags

Versions used in version 1 are:

* Go: 1.16.15
* Ignite (formerly Starport): 0.17.3
* Cosmos SDK: v0.42.6
* NodeJs: 16.x

Tags:

* [`v1-starport-start`](../../tree/v1-starport-start)
* [`v1-rules-added`](../../tree/v1-rules-added), [diff](../../compare/v1-starport-start..v1-rules-added)
* [`v1-stored-game`](../../tree/v1-stored-game), [diff](../../compare/v1-rules-added..v1-stored-game)
* [`v1-full-game-object`](../../tree/v1-full-game-object), [diff](../../compare/v1-stored-game..v1-full-game-object)
* [`v1-create-game-msg`](../../tree/v1-create-game-msg), [diff](../../compare/v1-full-game-object..v1-create-game-msg)
* [`v1-create-game-handler`](../../tree/v1-create-game-handler), [diff](../../compare/v1-create-game-msg..v1-create-game-handler)
* [`v1-play-move-msg`](../../tree/v1-play-move-msg), [diff](../../compare/v1-create-game-handler..v1-play-move-msg)
* [`v1-play-move-handler`](../../tree/v1-play-move-handler), [diff](../../compare/v1-play-move-msg..v1-play-move-handler)
* [`v1-two-events`](../../tree/v1-two-events), [diff](../../compare/v1-play-move-handler..v1-two-events)
* [`v1-reject-game-msg`](../../tree/v1-reject-game-msg), [diff](../../compare/v1-two-events..v1-reject-game-msg)
* [`v1-reject-game-handler`](../../tree/v1-reject-game-handler), [diff](../../compare/v1-reject-game-msg..v1-reject-game-handler)
* [`v1-game-fifo`](../../tree/v1-game-fifo), [diff](../../compare/v1-reject-game-handler..v1-game-fifo)
* [`v1-game-deadline`](../../tree/v1-game-deadline), [diff](../../compare/v1-game-fifo..v1-game-deadline)
* [`v1-game-winner`](../../tree/v1-game-winner), [diff](../../compare/v1-game-deadline..v1-game-winner)
* [`v1-forfeit-game`](../../tree/v1-forfeit-game), [diff](../../compare/v1-game-winner..v1-forfeit-game)
* [`v1-game-wager`](../../tree/v1-game-wager), [diff](../../compare/v1-forfeit-game..v1-game-wager)
* [`v1-payment-winning`](../../tree/v1-payment-winning), [diff](../../compare/v1-game-wager..v1-payment-winning)
* [`v1-gas-meter`](../../tree/v1-gas-meter), [diff](../../compare/v1-payment-winning..v1-gas-meter)
* [`v1-can-play-move-query`](../../tree/v1-can-play-move-query), [diff](../../compare/v1-gas-meter..v1-can-play-move-query)
* [`v1-can-play-move-handler`](../../tree/v1-can-play-move-handler), [diff](../../compare/v1-can-play-move-query..v1-can-play-move-handler)
* [`v1-wager-denomination`](../../tree/v1-wager-denomination), [diff](../../compare/v1-can-play-move-handler..v1-wager-denomination)
* [`v1-cosmjs-elements`](../../tree/v1-cosmjs-elements), [diff](../../compare/v1-wager-denomination..v1-cosmjs-elements)
* [`v1-player-info-object`](../../tree/v1-player-info-object), [diff](../../compare/v1-cosmjs-elements..v1-player-info-object)
* [`v1-player-info-handling`](../../tree/v1-player-info-handling), [diff](../../compare/v1-player-info-object..v1-player-info-handling)
* [`v1-leaderboard-object`](../../tree/v1-leaderboard-object), [diff](../../compare/v1-player-info-handling..v1-leaderboard-object)
* [`v1-leaderboard-handling`](../../tree/v1-leaderboard-handling), [diff](../../compare/v1-leaderboard-object..v1-leaderboard-handling)
* [`v1-genesis-migration`](../../tree/v1-genesis-migration), [diff](../../compare/v1-leaderboard-handling..v1-genesis-migration)
