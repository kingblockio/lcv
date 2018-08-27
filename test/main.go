package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/version"
	ibccmd "github.com/cosmos/cosmos-sdk/x/ibc/client/cli"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/cli"
)

// rootCmd is the entry point for this binary
var (
	rootCmd = &cobra.Command{
		Use:   "lcv",
		Short: "lcv test",
	}
)

func todoTest(_ *cobra.Command, _ []string) error {
	return errors.New("todo: Command not yet implemented")
}

func initTestCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start test light client",
		RunE:  todoTest,
	}
	return cmd
}

func main() {
	// disable sorting
	cobra.EnableCommandSorting = false

	// get the codec
	cdc := app.MakeCodec()

	// add proxy, version and key info
	rootCmd.AddCommand(
		initTestCommand(),
	)

	// prepare and add flags
	executor := cli.PrepareMainCmd(rootCmd, "BC", os.ExpandEnv("$HOME/.kinglcv"))
	err := executor.Execute()
	if err != nil {
		// Note: Handle with #870
		panic(err)
	}
}
