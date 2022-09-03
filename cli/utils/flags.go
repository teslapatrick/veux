package utils

import (
	"gopkg.in/urfave/cli.v1"
	"os"
	"path/filepath"
)

var NodeURLFlag = cli.StringFlag{
	Name:  "rpcList",
	Usage: "The rpc endpoint list of servial local or remote geth nodes(separator ',')",
	Value: "https://api.aspark.space/v1/eth/ropsten",
	//Value: "http://127.0.0.1:8540, http://localhost:8540",
}

var PathFlag = cli.StringFlag{
	Name:  "path",
	Usage: "The path of the file",
	Value: "./result.csv",
}

var FromFlag = cli.Int64Flag{
	Name:  "from",
	Usage: "From which block height.",
	Value: 0,
}

var ToFlag = cli.Int64Flag{
	Name:  "to",
	Usage: "To which block height.",
	Value: 0,
}

// NewApp creates an app with sane defaults.
func NewApp(gitCommit, gitDate, usage string) *cli.App {
	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Author = "teslapatrick"
	app.Email = "teslapatrick@gmail.com"
	app.Version = gitCommit + " - " + gitDate
	app.Usage = usage
	return app
}
