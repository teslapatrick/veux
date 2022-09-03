package chain

import (
	"fmt"
	"github.com/nknorg/nnet/log"
	chain "github.com/teslapatrick/veux/chain/fetcher"
	"github.com/teslapatrick/veux/cli/utils"
	"github.com/teslapatrick/veux/common"
	"gopkg.in/urfave/cli.v1"
	"time"
)

var GetBlockCommand = cli.Command{
	Name:   "getBlock",
	Usage:  "Get Blocks from the blockchain",
	Action: getBlock,
	Flags: []cli.Flag{
		utils.FromFlag,
		utils.ToFlag,
		utils.PathFlag,
	},
}

func getBlock(ctx *cli.Context) error {
	var (
		from = ctx.Int64(utils.FromFlag.Name)
		to   = ctx.Int64(utils.ToFlag.Name)
		path = ctx.String(utils.PathFlag.Name)
	)

	if to < from || (to == from && from == utils.FromFlag.Value) {
		err := fmt.Errorf("Invalid range: from %d to %d", from, to)
		return err
	}
	fmt.Println("Getting blocks: ", from, ":", to)

	// new clients
	// get urls from params
	clients := common.NewClientsFromURL(ctx)
	if len(clients) == 0 {
		log.Error("Getting clients Failed.", "len:", 0)
	}
	fmt.Println("Getting clients successfully,", "len:", len(clients))

	// new fetcher
	//var fetcher chain.Fetcher
	fetcher := chain.NewFetcher(clients)
	// results
	results := make([]interface{}, 0)

	// timer
	start := time.Now()
	// waitgroup
	//var wg sync.WaitGroup
	for number := from; number <= to; number++ {
		//wg.Add(1)
		// get blocks
		//go fetcher.FetchBlock(number, &results, &wg)
		fetcher.FetchBlock(number, &results)
	}
	//wg.Wait()
	fmt.Println("Getting Blocks Done...", "results len:", len(results), "Elapsed time:", time.Since(start))

	// write out to filepath
	err := common.WriteBlocksCSV(path, &results)
	if err == nil {
		log.Info("Successfully written to csv file")
	}

	return nil
}
