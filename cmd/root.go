package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/stratumfarm/derocli/pkg/dero"
)

var client *dero.Client

var rootCmdFlags struct {
	rpc string
}

var rootCmd = &cobra.Command{
	Use: "derocli",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(allCmd, heightCmd, infoCmd, peersCmd, txPoolCmd)

	rootCmd.PersistentFlags().StringVarP(&rootCmdFlags.rpc, "rpc", "r", "localhost:10102", "address of the node")
}

func Execute() error {
	return rootCmd.Execute()
}

func prettyPrint(data any) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to marshal: %w", err))
	}
	fmt.Println(string(b))
}

func connectNode(cmd *cobra.Command, args []string) error {
	c, err := dero.New("ws://" + rootCmdFlags.rpc + "/ws")
	if err != nil {
		return err
	}
	client = c
	return nil
}
