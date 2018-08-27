package main

import (
	"os"
	"io/ioutil"
	"encoding/json"
	
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/cli"
	tm "github.com/tendermint/tendermint/types"
)

func todoTest(_ *cobra.Command, _ []string) error {
	
	//
	registerGenesis := "/root/myspace/src/github.com/kingblockio/lcv/kinglcv/genesis.json"
	
	
	//
	genesisBytes, err := ioutil.ReadFile(registerGenesis)
	if err != nil {
		return errors.Errorf("Error reading genesis file %v: %v\n", registerGenesis, err)
	}
	
	//
	chainGenDoc := new(tm.GenesisDoc)
	err = json.Unmarshal(genesisBytes, chainGenDoc)
	if err != nil {
		return errors.Errorf("Genesis doc couldn't be parsed: %v\n", err)
	}
	
	//
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
