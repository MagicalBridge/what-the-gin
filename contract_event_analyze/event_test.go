package contract_event_analyze

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

/*
 *  txReceipt, logs
 *  getLogs    logs
 */

const ConfirmDataStoreEventABI = "ConfirmDataStore(uint32,bytes32)"

var ConfirmDataStoreEventABIHash = crypto.Keccak256Hash([]byte(ConfirmDataStoreEventABI))

const DataLayrServiceManagerAddr = "0x5BD63a7ECc13b955C4F57e3F12A64c10263C14c1"

func TestEthClient_GetTxReceiptByHash(t *testing.T) {
	fmt.Println("test start for tx receipt")
	clint, err := NewEthClient("https://rpc.mevblocker.io")
	if err != nil {
		fmt.Println("New eth client fail", err)
	}
	txReceipt, err := clint.GetTxReceiptByHash("0xfd26d40e17213bcafcf94bab9af92343302df9df970f20e1c9d515525e86e23e")
	if err != nil {
		fmt.Println("get tx receipt fail", err)
	}

	abiUint32, err := abi.NewType("uint32", "uint32", nil)
	if err != nil {
		fmt.Println("new uint32 abi type fail", err)
	}

	abiBytes32, err := abi.NewType("bytes32", "bytes32", nil)
	if err != nil {
		fmt.Println("new uint32 abi type fail", err)
	}
	confirmDataStoreArgs := abi.Arguments{
		{
			Name:    "dataStoreId",
			Type:    abiUint32,
			Indexed: false,
		}, {
			Name:    "headerHash",
			Type:    abiBytes32,
			Indexed: false,
		},
	}
	var dataStoreData = make(map[string]interface{})
	for _, rLog := range txReceipt.Logs {
		fmt.Println("address====", rLog.Address.String())
		if strings.ToLower(rLog.Address.String()) != strings.ToLower(DataLayrServiceManagerAddr) {
			continue
		}
		if rLog.Topics[0] != ConfirmDataStoreEventABIHash {
			continue
		}
		if len(rLog.Data) > 0 {
			err = confirmDataStoreArgs.UnpackIntoMap(dataStoreData, rLog.Data)
			if err != nil {
				fmt.Println("unpack data into mapping fail", err)
				continue
			}

			if dataStoreData != nil {
				fmt.Println("dataStoreId====", dataStoreData["dataStoreId"].(uint32))
				fmt.Println("dataStoreId====", dataStoreData["headerHash"])
			}
		}
	}
}

func TestEthClient_GetLogs(t *testing.T) {
	startBlock := big.NewInt(20483831)
	endBlock := big.NewInt(20483833)
	var contractList []common.Address
	addressCm := common.HexToAddress(DataLayrServiceManagerAddr)
	contractList = append(contractList, addressCm)
	clint, err := NewEthClient("https://rpc.payload.de")
	if err != nil {
		fmt.Println("connect ethereum fail", "err", err)
		return
	}
	logList, err := clint.GetLogs(startBlock, endBlock, contractList)
	if err != nil {
		fmt.Println("get log fail")
		return
	}
	abiUint32, err := abi.NewType("uint32", "uint32", nil)
	if err != nil {
		fmt.Println("Abi new uint32 type error", "err", err)
		return
	}
	abiBytes32, err := abi.NewType("bytes32", "bytes32", nil)
	if err != nil {
		fmt.Println("Abi new bytes32 type error", "err", err)
		return
	}
	confirmDataStoreArgs := abi.Arguments{
		{
			Name:    "dataStoreId",
			Type:    abiUint32,
			Indexed: false,
		}, {
			Name:    "headerHash",
			Type:    abiBytes32,
			Indexed: false,
		},
	}
	var dataStoreData = make(map[string]interface{})
	for _, rLog := range logList {
		fmt.Println(rLog.Address.String())
		if strings.ToLower(rLog.Address.String()) != strings.ToLower(DataLayrServiceManagerAddr) {
			continue
		}
		if rLog.Topics[0] != ConfirmDataStoreEventABIHash {
			continue
		}
		if len(rLog.Data) > 0 {
			err := confirmDataStoreArgs.UnpackIntoMap(dataStoreData, rLog.Data)
			if err != nil {
				fmt.Println("Unpack data into map fail", "err", err)
				continue
			}
			if dataStoreData != nil {
				dataStoreId := dataStoreData["dataStoreId"].(uint32)
				headerHash := dataStoreData["headerHash"]
				fmt.Println(dataStoreId)
				fmt.Println(headerHash)
			}
			return
		}
	}
}
