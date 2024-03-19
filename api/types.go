package api

import (
	"strings"
	"time"

	"github.com/0glabs/0g-storage-scan/stat"
	"github.com/shopspring/decimal"
)

type PageParam struct {
	Skip  int `form:"skip,default=0" binding:"omitempty,gte=0"`
	Limit int `form:"limit,default=10" binding:"omitempty,lte=2000"`
}

type statParam struct {
	PageParam
	MinTimestamp *int   `form:"minTimestamp" binding:"omitempty,number"`
	MaxTimestamp *int   `form:"maxTimestamp" binding:"omitempty,number"`
	IntervalType string `form:"intervalType,default=day" binding:"omitempty,oneof=hour day"`
	Sort         string `form:"sort,default=desc" binding:"omitempty,oneof=asc desc"`
}

func (sp *statParam) isDesc() bool {
	return strings.EqualFold(sp.Sort, "desc")
}

type listTxParam struct {
	PageParam
	Address  *string `form:"address" binding:"omitempty"`
	RootHash *string `form:"rootHash" binding:"omitempty"`
	Sort     string  `form:"sort,default=desc" binding:"omitempty,oneof=asc desc"`
}

func (sp *listTxParam) isDesc() bool {
	return strings.EqualFold(sp.Sort, "desc")
}

type queryTxParam struct {
	TxSeq *uint64 `form:"txSeq" binding:"required,number,gte=0"`
}

// StorageTx model info
// @Description Submission information
type StorageTx struct {
	TxSeq          uint64          `json:"txSeq"`          // Submission index in submit event
	BlockNum       uint64          `json:"blockNum"`       // The block where the submit event is emitted
	TxHash         string          `json:"txHash"`         // The transaction where the submit event is emitted
	RootHash       string          `json:"rootHash"`       // Merkle root of the file to upload
	Address        string          `json:"address"`        // File uploader address
	Method         string          `json:"method"`         // The name of the submit event is always `submit`
	Status         uint8           `json:"status"`         // File upload status, 0-not uploaded,1-uploading,2-uploaded
	TotalSegNum    uint64          `json:"totalSegNum"`    // The total number of segments the file is split into
	UploadedSegNum uint64          `json:"uploadedSegNum"` // The number of segments the file has been uploaded
	Timestamp      int64           `json:"timestamp"`      // The block time when submit event emits
	DataSize       uint64          `json:"dataSize"`       // File size in bytes
	BaseFee        decimal.Decimal `json:"baseFee"`        // The token fee required to upload the file
}

// TokenInfo model info
// @Description Charge token information
type TokenInfo struct {
	Address  string `json:"address"`  // The address of the token contract
	Name     string `json:"name"`     // Token name
	Symbol   string `json:"symbol"`   // Token symbol
	Decimals uint8  `json:"decimals"` // Token decimals
}

// CostInfo model info
// @Description Charge fee information
type CostInfo struct {
	TokenInfo `json:"tokenInfo"` // Charge token info
	BasicCost decimal.Decimal    `json:"basicCost"` // Charge fee
}

// TxList model info
// @Description Submission information list
type TxList struct {
	Total int64       `json:"total"` // The total number of submission returned
	List  []StorageTx `json:"list"`  // Submission list
}

// TxBrief model info
// @Description Submission brief information
type TxBrief struct {
	TxSeq  string `json:"txSeq"`  // Submission index in submit event
	From   string `json:"from"`   // File uploader address
	Method string `json:"method"` // The name of the submit event is always `submit`

	RootHash   string    `json:"rootHash"`   // Merkle root of the file to upload
	DataSize   uint64    `json:"dataSize"`   // File size in bytes
	Expiration uint64    `json:"expiration"` // Expiration date of the uploaded file
	CostInfo   *CostInfo `json:"costInfo"`   // Charge fee information

	BlockNumber uint64 `json:"blockNumber"` // The block where the submit event is emitted
	TxHash      string `json:"txHash"`      // The transaction where the submit event is emitted
	Timestamp   uint64 `json:"timestamp"`   // The block time when submit event emits
	Status      uint8  `json:"status"`      // The status of the transaction on layer1
	GasFee      uint64 `json:"gasFee"`      // The gas fee of the transaction on layer1
	GasUsed     uint64 `json:"gasUsed"`     // The gas used of the transaction on layer1
	GasLimit    uint64 `json:"gasLimit"`    // The gas limit of the transaction on layer1
}

// TxDetail model info
// @Description Submission detail information
type TxDetail struct {
	TxSeq    string `json:"txSeq"`    // Submission index in submit event
	RootHash string `json:"rootHash"` // Merkle root of the file to upload

	StartPos    uint64 `json:"startPos"`    // The starting position of the file stored in the storage node
	EndPos      uint64 `json:"endPos"`      // The ending position of the file stored in the storage node
	PieceCounts uint64 `json:"pieceCounts"` // The total number of segments the file is split into
}

// StorageBasicCost model info
// @Description Storage fee information
type StorageBasicCost struct {
	TokenInfo                      // Charge token info
	BasicCostTotal decimal.Decimal `json:"basicCostTotal"` // Total storage fee
}

// Dashboard model info
// @Description Storage status information
type Dashboard struct {
	StorageBasicCost `json:"storageBasicCost"` // Storage fee information
	stat.LogSyncInfo `json:"logSyncInfo"`      // Synchronization information of submit event
}

// DataStatList model info
// @Description Storage data list
type DataStatList struct {
	Total int64      `json:"total"` // The total number of stat returned
	List  []DataStat `json:"list"`  // Stat list
}

// TxStatList model info
// @Description Storage transaction list
type TxStatList struct {
	Total int64    `json:"total"` // The total number of stat returned
	List  []TxStat `json:"list"`  // Stat list
}

// FeeStatList model info
// @Description Storage fee list
type FeeStatList struct {
	Total int64     `json:"total"` // The total number of stat returned
	List  []FeeStat `json:"list"`  // Stat list
}

// DataStat model info
// @Description Storage data information
type DataStat struct {
	StatTime  time.Time `json:"statTime"`  // Statistics time
	FileCount uint64    `json:"fileCount"` // Number of files in a specific time interval
	FileTotal uint64    `json:"fileTotal"` // Total number of files by a certain time
	DataSize  uint64    `json:"dataSize"`  // Size of storage data in a specific time interval
	DataTotal uint64    `json:"dataTotal"` // Total Size of storage data by a certain time
}

// TxStat model info
// @Description Storage transaction information
type TxStat struct {
	StatTime time.Time `json:"statTime"` // Statistics time
	TxCount  uint64    `json:"txCount"`  // Number of layer1 transaction in a specific time interval
	TxTotal  uint64    `json:"txTotal"`  // Total number of layer1 transaction by a certain time
}

// FeeStat model info
// @Description Storage fee information
type FeeStat struct {
	StatTime     time.Time       `json:"statTime"`     // Statistics time
	BaseFee      decimal.Decimal `json:"baseFee"`      // The base fee for storage in a specific time interval
	BaseFeeTotal decimal.Decimal `json:"baseFeeTotal"` // The total base fee for storage by a certain time
}
