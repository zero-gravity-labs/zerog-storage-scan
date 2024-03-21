package stat

// LogSyncInfo model info
// @Description Submit log sync information
type LogSyncInfo struct {
	Layer1LogSyncHeight uint64 `json:"layer1-logSyncHeight"` // Synchronization height of submit log on blockchain
	LogSyncHeight       uint64 `json:"logSyncHeight"`        // Synchronization height of submit log on storage node
}
