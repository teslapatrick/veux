package common

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"

	eu "github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/teslapatrick/veux/cli/utils"
	"gopkg.in/urfave/cli.v1"
)

func newClient(url string) *ethclient.Client {
	client, err := ethclient.Dial(url)
	if err != nil {
		eu.Fatalf("Failed to connect to ethereum node: %v", err)
		return nil
	}

	return client
}

func NewClients(urls []string) []*ethclient.Client {
	clients := make([]*ethclient.Client, 0)

	for _, url := range urls {
		client := newClient(url)
		if client != nil {
			clients = append(clients, client)
		}
	}

	return clients
}

func NewClientsFromURL(ctx *cli.Context) []*ethclient.Client {
	urls := getURLs(ctx)
	return NewClients(urls)
}

func getURLs(ctx *cli.Context) []string {
	urlStr := ctx.GlobalString(utils.NodeURLFlag.Name)
	list := make([]string, 0)

	for _, url := range strings.Split(urlStr, ",") {
		if url = strings.TrimSpace(url); url != "" {
			list = append(list, url)
		}
	}

	return list
}

func WriteBlocksCSV(path string, results *[]interface{}) error {
	// init csv file
	f, err := os.Create(path)
	if err != nil {
		log.Error("Failed to create file", "path", path, "err", err)
		return err
	}
	defer f.Close()

	// Insert UTF-8 BOM
	//_, _ = f.WriteString("\xEF\xBB\xBF")
	// NewWriter
	w := csv.NewWriter(f)
	// init data
	data := make([][]string, 0)
	//for i := range data {
	//	data[i] = make([]string, 0)
	//}
	data = append(data, []string{ // init csv header
		"BlockNumber",
		"BlockHash",
		"TransactionsLength",
	})
	// loop
	for _, ele := range *results {
		block := ele.(*types.Block)
		data = append(data, []string{
			block.Number().String(),
			block.Hash().String(),
			strconv.Itoa(block.Transactions().Len()),
		})
	}
	// write out
	err = w.WriteAll(data)
	if err != nil {
		log.Error("Failed to write file", "path", path, "err", err)
		return err
	}
	w.Flush()

	return nil
}
