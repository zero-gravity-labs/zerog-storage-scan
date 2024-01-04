package contract

import (
	"github.com/zero-gravity-labs/zerog-storage-client/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
)

var (
	flowFilterer       *contract.FlowFilterer
	erc20TokenFilterer *Erc20TokenFilterer
)

func init() {
	flowFilterer, _ = contract.NewFlowFilterer(common.HexToAddress(""), nil)
	erc20TokenFilterer, _ = NewErc20TokenFilterer(common.HexToAddress(""), nil)
}

func DummyFlowFilterer() *contract.FlowFilterer {
	return flowFilterer
}

func DummyErc20TokenFilterer() *Erc20TokenFilterer {
	return erc20TokenFilterer
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
