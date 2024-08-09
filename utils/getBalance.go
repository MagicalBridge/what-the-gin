package utils

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func GetBalance(targetAddress string) (string, *big.Float, error) {
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}

	// 要查询的钱包地址
	address := common.HexToAddress(targetAddress)

	// 获取余额
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 将余额从Wei转换为Ether
	floatBalance := new(big.Float)
	floatBalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(floatBalance, big.NewFloat(1e18))

	fmt.Printf("账户 %s 的余额: %f ETH\n", address.Hex(), ethValue)

	return address.Hex(), ethValue, err
}
