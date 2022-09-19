package main

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/protolambda/clique/flags"
	"github.com/protolambda/clique/version"
	"github.com/urfave/cli"
	"os"
)

var (
	GitCommit = ""
	GitDate   = ""
)

// VersionWithMeta holds the textual version string including the metadata.
var VersionWithMeta = func() string {
	v := version.Version
	if GitCommit != "" {
		v += "-" + GitCommit[:8]
	}
	if GitDate != "" {
		v += "-" + GitDate
	}
	if version.Meta != "" {
		v += "-" + version.Meta
	}
	return v
}()

func main() {
	// Set up logger with a default INFO level in case we fail to parse flags,
	// otherwise the final critical log won't show what the parsing error was.
	log.Root().SetHandler(
		log.LvlFilterHandler(
			log.LvlInfo,
			log.StreamHandler(os.Stdout, log.TerminalFormat(true)),
		),
	)

	app := cli.NewApp()
	app.Version = VersionWithMeta
	app.Flags = flags.Flags
	app.Name = "clique"
	app.Usage = "Clique runs a PoA consensus node, coupled to an ethereum Engine"
	app.Description = "Clique is a Proof of Authority (PoA) consensus protocol for ethereum, defined in EIP-225. " +
		"This node implements clique, and inserts the ethereum blocks, also known as execution payloads, into an execution engine through an authenticated RPC."
	app.Action = CliqueMain

	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}

func CliqueMain(ctx *cli.Context) error {
	// TODO: init logger from cli flags

	// TODO: start clique node

	// TODO: wait for sys signal to close clique node
	return nil
}
