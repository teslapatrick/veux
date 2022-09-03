package main

import (
	"fmt"
	"os"

	"github.com/teslapatrick/veux/cli/chain"
	"github.com/teslapatrick/veux/cli/utils"
	"gopkg.in/urfave/cli.v1"
)

var (
	app = utils.NewApp("", "", "Blockchain toolkit.")
)

func init() {

	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend!", os.Args[0])
		return nil
	}

	app.Commands = []cli.Command{
		chain.GetBlockCommand,
	}

	app.Flags = []cli.Flag{
		utils.NodeURLFlag,
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Failed to running blockchain toolkit: %v\n", err)
		os.Exit(1)
	}
}
