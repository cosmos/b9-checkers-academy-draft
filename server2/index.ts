require("./indexer")
    .createIndexer({
        listenPort: "3001",
        dbFile: "./db.json",
        pollIntervalMs: 5_000, // 5 seconds
        rpcPoint: "http://localhost:26657",
    })
    .then(console.log)
    .catch(console.error)
