package sync

import (
	"context"
	"github.com/zero-gravity-labs/zerog-storage-scan/store"
	set "github.com/deckarep/golang-set"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/go-rpc-provider"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrChainReorged = errors.New("chain re-orged")
)

func isTxExecutedInBlock(tx *types.TransactionDetail, receipt *types.Receipt) bool {
	return tx != nil && receipt.Status != nil && *receipt.Status < 2
}

func queryEthData(w3c *web3go.Client, blockNumber uint64) (*store.EthData, error) {
	// get block
	block, err := w3c.Eth.BlockByNumber(types.BlockNumber(blockNumber), true)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to get block by number %v", blockNumber)
	}
	if block == nil {
		return nil, nil
	}

	// batch get receipts
	blockNumOrHash := types.BlockNumberOrHashWithNumber(types.BlockNumber(blockNumber))
	blockReceipts, err := w3c.Parity.BlockReceipts(&blockNumOrHash)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get block receipts")
	}
	if blockReceipts == nil {
		return nil, errors.WithMessage(ErrChainReorged, "batch retrieved block receipts nil")
	}

	// get receipt
	txReceipts := map[common.Hash]*types.Receipt{}
	blockTxs := block.Transactions.Transactions()
	if len(blockTxs) != len(blockReceipts) {
		return nil, errors.Errorf("block receipts number mismatch, rcpts %v, txs %v", len(blockReceipts), len(blockTxs))
	}
	for i := 0; i < len(blockTxs); i++ {
		tx := blockTxs[i]
		var receipt *types.Receipt
		receipt = &blockReceipts[i]

		// check re-org
		switch {
		case receipt == nil: // receipt shouldn't be nil unless chain re-org
			return nil, errors.WithMessage(ErrChainReorged, "tx receipt nil")
		case receipt.BlockHash != block.Hash:
			return nil, errors.WithMessagef(ErrChainReorged, "receipt block hash mismatch, rcptBlkHash %v, blkHash %v",
				receipt.BlockHash, block.Hash)
		case receipt.BlockNumber != blockNumber:
			return nil, errors.WithMessagef(ErrChainReorged, "receipt block num mismatch, rcptBlkNum %v, blkNum %v",
				receipt.BlockNumber, blockNumber)
		case receipt.TransactionHash != tx.Hash:
			return nil, errors.WithMessagef(ErrChainReorged, "receipt tx hash mismatch, rcptTxHash %v, TxHash %v",
				receipt.TransactionHash, tx.Hash)
		}
		txReceipts[tx.Hash] = receipt
	}

	return &store.EthData{blockNumber, block, txReceipts}, nil
}

func queryFlowSubmits(w3c *web3go.Client, blockFrom, blockTo uint64, flowAddr common.Address,
	flowSubmitSig common.Hash) ([]types.Log, error) {
	bnFrom := types.NewBlockNumber(int64(blockFrom))
	bnTo := types.NewBlockNumber(int64(blockTo))
	logFilter := types.FilterQuery{
		FromBlock: &bnFrom,
		ToBlock:   &bnTo,
		Addresses: []common.Address{flowAddr},
		Topics:    [][]common.Hash{{flowSubmitSig}},
	}
	return w3c.Eth.Logs(logFilter)
}

func queryErc20Transfers(w3c *web3go.Client, blockFrom, blockTo uint64, erc20Addr common.Address, erc20TransferSig,
	flowAddrTopic common.Hash) ([]types.Log, error) {
	bnFrom := types.NewBlockNumber(int64(blockFrom))
	bnTo := types.NewBlockNumber(int64(blockTo))

	logFilter := types.FilterQuery{
		FromBlock: &bnFrom,
		ToBlock:   &bnTo,
		Addresses: []common.Address{erc20Addr},
		Topics:    [][]common.Hash{{erc20TransferSig}, {}, {flowAddrTopic}},
	}
	return w3c.Eth.Logs(logFilter)
}

func mapBlockNum2Time(ctx context.Context, w3c *web3go.Client, blkNums []types.BlockNumber,
	batchSize uint64) (map[uint64]uint64, error) {
	if len(blkNums) == 0 {
		return nil, errors.New("no block numbers")
	}

	blkNumSet := set.NewSet()
	for _, num := range blkNums {
		blkNumSet.Add(num)
	}

	blockNum2Time := make(map[uint64]uint64)
	blkNumSlice := blkNumSet.ToSlice()
	blkNumSize := len(blkNumSlice)
	for i := 0; i < blkNumSize; i += int(batchSize) {
		end := i + int(batchSize)
		if end > blkNumSize {
			end = blkNumSize
		}
		blockNums := blkNumSlice[i:end]

		batch := make([]rpc.BatchElem, 0)
		for _, blkNum := range blockNums {
			elem := rpc.BatchElem{
				Method: "eth_getBlockByNumber",
				Args:   []interface{}{blkNum, false},
				Result: new(types.Block),
			}
			batch = append(batch, elem)
		}

		err := w3c.Eth.BatchCallContext(ctx, batch)
		if err != nil {
			return nil, err
		}

		for _, elem := range batch {
			block := elem.Result.(*types.Block)
			blockNum2Time[block.Number.Uint64()] = block.Timestamp
		}
	}

	return blockNum2Time, nil
}

type transaction struct {
	tx   *types.TransactionDetail
	rcpt *types.Receipt
}

func queryTxsByHashes(ctx context.Context, w3c *web3go.Client, hashes []common.Hash, batchSize uint64) (
	[]*transaction, error) {
	if len(hashes) == 0 {
		return nil, errors.New("no tx hashes")
	}

	hashSet := set.NewSet()
	for _, hash := range hashes {
		hashSet.Add(hash)
	}
	hashSlice := hashSet.ToSlice()
	size := len(hashSlice)

	hashTxMap := make(map[common.Hash]*types.TransactionDetail)
	hashRcptMap := make(map[common.Hash]*types.Receipt)
	for i := 0; i < size; i += int(batchSize) {
		end := i + int(batchSize)
		if end > size {
			end = size
		}
		hashes := hashSlice[i:end]

		batch := make([]rpc.BatchElem, 0)
		for _, hash := range hashes {
			elem := rpc.BatchElem{
				Method: "eth_getTransactionByHash",
				Args:   []interface{}{hash},
				Result: new(types.TransactionDetail),
			}
			elemR := rpc.BatchElem{
				Method: "eth_getTransactionReceipt",
				Args:   []interface{}{hash},
				Result: new(types.Receipt),
			}
			batch = append(batch, elem)
			batch = append(batch, elemR)
		}

		err := w3c.Eth.BatchCallContext(ctx, batch)
		if err != nil {
			return nil, err
		}

		for i, elem := range batch {
			if i%2 == 0 {
				txn := elem.Result.(*types.TransactionDetail)
				hashTxMap[txn.Hash] = txn
			} else {
				rcpt := elem.Result.(*types.Receipt)
				hashRcptMap[rcpt.TransactionHash] = rcpt
			}
		}
	}

	txns := make([]*transaction, 0)
	for _, hash := range hashes { // for returning txs in block number's asc order
		tx := hashTxMap[hash]
		rcpt := hashRcptMap[hash]
		if tx == nil || rcpt == nil {
			continue
		}
		delete(hashTxMap, hash)
		delete(hashRcptMap, hash)
		txns = append(txns, &transaction{tx: tx, rcpt: rcpt})
	}

	return txns, nil
}
