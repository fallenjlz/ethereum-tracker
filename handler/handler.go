package handler

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"sync"

	"ethereum-tracker/monitor"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

type LogDetails struct {
	Name        string      `json:"name"`
	BlockNumber uint64      `json:"block_number"`
	Index       uint        `json:"index"`
	Data        interface{} `json:"data"`
}

var (
	LogDataMutex sync.Mutex
	LogEntries   []LogDetails
)

func HandleLog(vLog types.Log, contractAbi abi.ABI, logTransferSigHash common.Hash, logApprovalSigHash common.Hash) {
	var logDetail LogDetails
	logDetail.BlockNumber = vLog.BlockNumber
	logDetail.Index = vLog.Index

	switch vLog.Topics[0].Hex() {
	case logTransferSigHash.Hex():
		logDetail.Name = "Transfer"
		monitor.TxCounter.Inc()

		var transferEvent LogTransfer
		err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
		transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
		transferEvent.Tokens = new(big.Int).SetBytes(vLog.Data)

		tokensTransferValue, _ := new(big.Float).SetInt(transferEvent.Tokens).Float64()
		monitor.TokensTransferred.Set(tokensTransferValue)

		logDetail.Data = transferEvent

	case logApprovalSigHash.Hex():
		logDetail.Name = "Approval"

		var approvalEvent LogApproval
		err := contractAbi.UnpackIntoInterface(&approvalEvent, "Approval", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
		approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())
		approvalEvent.Tokens = new(big.Int).SetBytes(vLog.Data)

		logDetail.Data = approvalEvent
	}

	updateLogEntries(logDetail)
}

func updateLogEntries(logDetail LogDetails) {
	LogDataMutex.Lock()
	defer LogDataMutex.Unlock()

	if len(LogEntries) >= 5 {
		LogEntries = LogEntries[1:] // Keep only the latest five entries
	}
	LogEntries = append(LogEntries, logDetail)
}

func LatestLogsHandler(w http.ResponseWriter, r *http.Request) {
	LogDataMutex.Lock()
	defer LogDataMutex.Unlock()

	json.NewEncoder(w).Encode(LogEntries)
}
