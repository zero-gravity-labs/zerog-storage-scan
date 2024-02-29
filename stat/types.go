package stat

type LogSyncInfo struct {
	LogSyncHeight   uint64 `json:"logSyncHeight"`
	L2LogSyncHeight uint64 `json:"l2LogSyncHeight"`
}
