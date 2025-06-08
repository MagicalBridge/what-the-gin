package contract_event_analyze

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

type EthClient struct {
	client *ethclient.Client
}

func NewEthClient(rpcUrl string) (*EthClient, error) {
	ethClient, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		log.Error("new eth client fail", "err", err)
		return nil, err
	}
	return &EthClient{client: ethClient}, nil
}

// GetTxReceiptByHash eth_getTransactionReceipt
func (eth *EthClient) GetTxReceiptByHash(txHash string) (*types.Receipt, error) {
	return eth.client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
}

// GetLogs eth_getLogs
func (eth *EthClient) GetLogs(startBlock, endBlock *big.Int, contractAddressList []common.Address) ([]types.Log, error) {
	filterQueryParams := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: contractAddressList,
	}
	return eth.client.FilterLogs(context.Background(), filterQueryParams)
}
