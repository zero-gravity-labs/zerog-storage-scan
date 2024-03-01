package api

import (
	"math/big"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

type PageParam struct {
	Skip  int `form:"skip,default=0" binding:"omitempty,gte=0"`
	Limit int `form:"limit,default=10" binding:"omitempty,lte=2000"`
}

type statParam struct {
	PageParam
	MinTimestamp int    `form:"minTimestamp,default=0" binding:"omitempty,number"`
	MaxTimestamp int    `form:"maxTimestamp,default=0" binding:"omitempty,number"`
	IntervalType string `form:"intervalType,default=day" binding:"omitempty,oneof=hour day"`
	Sort         string `form:"sort,default=desc" binding:"omitempty,oneof=asc desc"`
}

func (sp *statParam) isDesc() bool {
	return strings.EqualFold(sp.Sort, "desc")
}

type listTxParam struct {
	PageParam
	Address  string `form:"address" binding:"omitempty"`
	RootHash string `form:"rootHash" binding:"omitempty"`
}

type queryTxParam struct {
	TxSeq *uint64 `form:"txSeq" binding:"required,number,gte=0"`
}

type StorageTx struct {
	TxSeq     uint64          `json:"txSeq"`
	BlockNum  uint64          `json:"blockNum"`
	TxHash    string          `json:"txHash"`
	RootHash  string          `json:"rootHash"`
	Address   string          `json:"address"`
	Method    string          `json:"method"`
	Status    uint64          `json:"status"`
	Timestamp int64           `json:"timestamp"`
	DataSize  uint64          `json:"dataSize"`
	BaseFee   decimal.Decimal `json:"baseFee"`
}

type TokenInfo struct {
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals uint8  `json:"decimals"`
}

type CostInfo struct {
	TokenInfo `json:"tokenInfo"`
	BasicCost decimal.Decimal `json:"basicCost"`
}

type SubmissionNode struct {
	Root   string   `json:"root"`
	Height *big.Int `json:"height"`
}

type TxList struct {
	Total int64       `json:"total"`
	List  []StorageTx `json:"list"`
}

type TxBrief struct {
	TxSeq  string `json:"txSeq"`
	From   string `json:"from"`
	Method string `json:"method"`

	RootHash   string    `json:"rootHash"`
	DataSize   uint64    `json:"dataSize"`
	Expiration uint64    `json:"expiration"`
	CostInfo   *CostInfo `json:"costInfo"`

	BlockNumber uint64 `json:"blockNumber"`
	TxHash      string `json:"txHash"`
	Timestamp   uint64 `json:"timestamp"`
	Status      uint64 `json:"status"`
	GasFee      uint64 `json:"gasFee"`
	GasUsed     uint64 `json:"gasUsed"`
	GasLimit    uint64 `json:"gasLimit"`
}

type TxDetail struct {
	TxSeq    string `json:"txSeq"`
	RootHash string `json:"rootHash"`

	StartPos    uint64           `json:"startPos"`
	EndPos      uint64           `json:"endPos"`
	PieceCounts uint64           `json:"pieceCounts"`
	Pieces      []SubmissionNode `json:"pieces"`
}

type StorageBasicCost struct {
	TokenInfo
	BasicCostTotal decimal.Decimal `json:"basicCostTotal"`
}

type Dashboard struct {
	StorageBasicCost `json:"storageBasicCost"`
}

type DataStatList struct {
	Total int64      `json:"total"`
	List  []DataStat `json:"list"`
}

type TxStatList struct {
	Total int64    `json:"total"`
	List  []TxStat `json:"list"`
}

type FeeStatList struct {
	Total int64     `json:"total"`
	List  []FeeStat `json:"list"`
}

type DataStat struct {
	StatTime  time.Time `json:"statTime"`
	FileCount uint64    `json:"fileCount"`
	FileTotal uint64    `json:"fileTotal"`
	DataSize  uint64    `json:"dataSize"`
	DataTotal uint64    `json:"dataTotal"`
}

type TxStat struct {
	StatTime time.Time `json:"statTime"`
	TxCount  uint64    `json:"txCount"`
	TxTotal  uint64    `json:"txTotal"`
}

type FeeStat struct {
	StatTime     time.Time       `json:"statTime"`
	BaseFee      decimal.Decimal `json:"baseFee"`
	BaseFeeTotal decimal.Decimal `json:"baseFeeTotal"`
}
