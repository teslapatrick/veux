package chain

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"sync"
)

type Fetcher interface {
	// FetchBlock fetches the block at the given block number.
	//FetchBlock(numnber int64, results *[]interface{}, wg *sync.WaitGroup)
	FetchBlock(numnber int64, results *[]interface{})
	// FetchTxs	 fetches the transactions from the world states.
	FetchTxs(txHash string, results []interface{}) (types.Transaction, error)
}

type EthFetcher struct {
	clients []*ethclient.Client
	current int
	Lock    sync.Mutex
}

func NewFetcher(clients []*ethclient.Client) *EthFetcher {
	ef := &EthFetcher{
		clients: clients,
		current: 0,
	}
	return ef
}

func (ef *EthFetcher) SetClients(clients []*ethclient.Client) {
	ef.clients = clients
}

func (ef *EthFetcher) NextClient() *ethclient.Client {
	l := len(ef.clients)
	//fmt.Println("client: ", ef.current)
	client := ef.clients[ef.current]
	if ef.current++; ef.current >= l {
		ef.current = 0
	}
	return client
}

//func (ef *EthFetcher) FetchBlock(number int64, results *[]interface{}, wg *sync.WaitGroup) {
func (ef *EthFetcher) FetchBlock(number int64, results *[]interface{}) {
	client := ef.NextClient()

	nb := new(big.Int).SetInt64(number)
	block, err := client.BlockByNumber(context.Background(), nb)
	if err != nil {
		log.Error("Failed to receive block: number=%v err=%v", nb.String(), err)
	}
	if block != nil {
		*results = append(*results, block)
	}
	//wg.Done()
}

func (ef *EthFetcher) FetchTxs(txHash string, results []interface{}) (types.Transaction, error) {
	tx := types.Transaction{}
	return tx, nil
}
