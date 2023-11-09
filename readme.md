# Checkers tutorial project

This is a companion project of the Cosmos SDK tutorials. Its object is to show various features of the Cosmos SDK and of Ignite, along with the progression of the code as elements and features are added.

The progression of the code is demonstrated via the help of branches and diffs.

## Build steps taken

All the build steps were run inside the Docker container.

```sh
$ ignite scaffold chain github.com/b9lab/checkers
```

## Progressive feature branches

The project is created with a clean list of commits in order to demonstrate the natural progression of the project. In this sense, there is no late commit that fixes an error introduced earlier. If there is a fix for an error introduced earlier, the fix should be squashed with the earlier commit that introduced the error. This may require some conflict resolution.

Having a clean list of commits makes it possible to do meaningful `compare`s.

To make it easier to link to the content at the different stages of the project's progression, a number of branches have been created at commits that have `Add branch the-branch-name` as message. Be careful with the commit message as it depends on it matching the `"Add branch [0-9a-zA-Z\-]*"` regular expression.

The script `push-branches.sh` is used to extract these commits and force push them to the appropriate branch in the repository. After having made changes, you should run this script, and also manually force push to `main`.

Versions used here are:

* Go: 1.18.3
* Ignite (formerly Starport): 0.22.1
* Cosmos SDK: v0.45.4

Branches:

* [`ignite-start`](../../tree/ignite-start)

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
