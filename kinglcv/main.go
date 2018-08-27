package main

import (
	"os"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/cli"
)

func todoTest(_ *cobra.Command, _ []string) error {
	return errors.New("todo: Command not yet implemented")
}


func main() {
	
	rootCmd := &cobra.Command{
		Use:   "start",
		Short: "start test light client",
		RunE:  todoTest,
	}
		
	// disable sorting
	cobra.EnableCommandSorting = false

	// prepare and add flags
	executor := cli.PrepareMainCmd(rootCmd, "BC", os.ExpandEnv("$HOME/.kinglcv"))
	err := executor.Execute()
	if err != nil {
		// Note: Handle with #870
		panic(err)
	}
}
