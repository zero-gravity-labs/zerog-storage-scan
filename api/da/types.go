package da

import (
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

type PageParam struct {
	Skip  int    `form:"skip,default=0" binding:"omitempty,gte=0"`
	Limit int    `form:"limit,default=10" binding:"omitempty,lte=2000"`
	Sort  string `form:"sort,default=desc" binding:"omitempty,oneof=asc desc"`
}

func (sp *PageParam) isDesc() bool {
	return strings.EqualFold(sp.Sort, "desc")
}

type statParam struct {
	PageParam
	MinTimestamp *int   `form:"minTimestamp" binding:"omitempty,number"`
	MaxTimestamp *int   `form:"maxTimestamp" binding:"omitempty,number"`
	IntervalType string `form:"intervalType,default=day" binding:"omitempty,oneof=hour day"`
}

// TokenInfo model info
// @Description Charge token information
type TokenInfo struct {
	Address  string `json:"address"`  // The address of the token contract
	Name     string `json:"name"`     // Token name
	Symbol   string `json:"symbol"`   // Token symbol
	Decimals uint8  `json:"decimals"` // Token decimals
	Native   bool   `json:"native"`   // True is native token, otherwise is not
}

type listDATxParam struct {
	PageParam
	RootHash *string `form:"rootHash" binding:"omitempty"`
	TxHash   *string `form:"txHash" binding:"omitempty"`
}

// DATxList model info
// @Description DA submission information list
type DATxList struct {
	Total int64      `json:"total"` // The total number of da submission returned
	List  []DATxInfo `json:"list"`  // DA submission list
}

// DATxInfo model info
// @Description DA submission transaction information
type DATxInfo struct {
	BlockNumber uint64 `json:"blockNumber"` // The block where the submit event is emitted
	TxHash      string `json:"txHash"`      // The transaction where the submit event is emitted
	Timestamp   int64  `json:"timestamp"`   // The block time when submit event emits
	From        string `json:"from"`        // File uploader address
	Method      string `json:"method"`      // The name of the submit event

	Epoch      uint64          `json:"epoch"`      // Epoch index in DataUpload event
	QuorumID   uint64          `json:"quorumID"`   // QuorumID in DataUpload event
	RootHash   string          `json:"rootHash"`   // Merkle root of the data to upload
	StorageFee decimal.Decimal `json:"storageFee"` // The storage fee required to upload the file
	Status     uint8           `json:"status"`     // Data upload status, 0-not verified,1-verified
}

// DADataStatList model info
// @Description DA storage data list
type DADataStatList struct {
	Total int64        `json:"total"` // The total number of stat returned
	List  []DADataStat `json:"list"`  // Stat list
}

// DAClientStatList model info
// @Description DA client stat list
type DAClientStatList struct {
	Total int64          `json:"total"` // The total number of stat returned
	List  []DAClientStat `json:"list"`  // Stat list
}

// DASignerStatList model info
// @Description DA signer stat list
type DASignerStatList struct {
	Total int64          `json:"total"` // The total number of stat returned
	List  []DASignerStat `json:"list"`  // Stat list
}

// DADataStat model info
// @Description DA storage data information
type DADataStat struct {
	StatTime        time.Time       `json:"statTime"`        // Statistics time
	BlobNew         uint64          `json:"blobNew"`         // Number of blobs in a specific time interval
	BlobTotal       uint64          `json:"blobTotal"`       // Total number of blobs by a certain time
	DataSizeNew     uint64          `json:"dataSizeNew"`     // Size of storage data in a specific time interval
	DataSizeTotal   uint64          `json:"dataSizeTotal"`   // Total size of storage data by a certain time
	StorageFeeNew   decimal.Decimal `json:"storageFeeNew"`   // Storage fee in a specific time interval
	StorageFeeTotal decimal.Decimal `json:"storageFeeTotal"` // Total storage fee by a certain time
}

// DAClientStat model info
// @Description DA client information
type DAClientStat struct {
	StatTime     time.Time `json:"statTime"`     // Statistics time
	ClientNew    uint64    `json:"clientNew"`    // Number of da client in a specific time interval
	ClientActive uint64    `json:"clientActive"` // Number of active da client in a specific time interval
	ClientTotal  uint64    `json:"clientTotal"`  // Total number of da client by a certain time
}

// DASignerStat model info
// @Description DA signer information
type DASignerStat struct {
	StatTime     time.Time `json:"statTime"`     // Statistics time
	SignerNew    uint64    `json:"signerNew"`    // Number of da signer in a specific time interval
	SignerActive uint64    `json:"signerActive"` // Number of active da signer in a specific time interval
	SignerTotal  uint64    `json:"signerTotal"`  // Total number of da signer by a certain time
}

// RewardList model info
// @Description Miner reward list
type RewardList struct {
	Total int64    `json:"total"` // The total number of miner reward returned
	List  []Reward `json:"list"`  // Miner reward list
}

// Reward model info
// @Description Reward information
type Reward struct {
	Miner       string          `json:"miner"`       // Miner address
	Amount      decimal.Decimal `json:"amount"`      // The reward amount
	BlockNumber uint64          `json:"blockNumber"` // The block where the reward event is emitted
	TxHash      string          `json:"txHash"`      // The transaction where the reward event is emitted
	Timestamp   int64           `json:"timestamp"`   // The block time when reward event emits
	SampleRound uint64          `json:"sampleRound"` // DA Sample round
	Epoch       uint64          `json:"epoch"`       // The consecutive blocks in 0g chain is divided into groups of EpochBlocks and each group is an epoch.
	QuorumID    uint64          `json:"quorumID"`    // The i-th quorum in an epoch
	RootHash    string          `json:"rootHash"`    // The data root
}

type AddressInfo struct {
	address   string
	addressId uint64
}
