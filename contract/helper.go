package contract

import (
	"github.com/0glabs/0g-storage-client/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
)

var (
	flowFilterer       *contract.FlowFilterer
	rewardFilterer     *OnePoolRewardFilterer
	mineFilterer       *PoraMineFilterer
	daSignersFilterer  *DASignersFilterer
	daEntranceFilterer *DAEntranceFilterer
)

func init() {
	flowFilterer, _ = contract.NewFlowFilterer(common.HexToAddress(""), nil)
	rewardFilterer, _ = NewOnePoolRewardFilterer(common.HexToAddress(""), nil)
	mineFilterer, _ = NewPoraMineFilterer(common.HexToAddress(""), nil)
	daSignersFilterer, _ = NewDASignersFilterer(common.HexToAddress(""), nil)
	daEntranceFilterer, _ = NewDAEntranceFilterer(common.HexToAddress(""), nil)
}

func DummyFlowFilterer() *contract.FlowFilterer {
	return flowFilterer
}

func DummyRewardFilterer() *OnePoolRewardFilterer {
	return rewardFilterer
}

func DummyMineFilterer() *PoraMineFilterer {
	return mineFilterer
}

func DummyDASignersFilterer() *DASignersFilterer {
	return daSignersFilterer
}

func DummyDAEntranceFilterer() *DAEntranceFilterer {
	return daEntranceFilterer
}

func TokenInfo(w3c *web3go.Client, address string) (string, string, uint8, error) {
	ethClient, _ := w3c.ToClientForContract()
	token, err := NewErc20Token(common.HexToAddress(address), ethClient)
	if err != nil {
		return "", "", 0, err
	}

	name, err := token.Name(nil)
	if err != nil {
		return "", "", 0, err
	}

	symbol, err := token.Symbol(nil)
	if err != nil {
		return "", "", 0, err
	}

	decimals, err := token.Decimals(nil)
	if err != nil {
		return "", "", 0, err
	}

	return name, symbol, decimals, nil
}
