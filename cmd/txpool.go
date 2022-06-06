package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var txPoolCmd = &cobra.Command{
	Use:     "txpool",
	Short:   "Get the transaction pool",
	PreRunE: connectNode,
	PostRun: func(cmd *cobra.Command, args []string) { client.Close() },
	RunE:    txPool,
}

func txPool(cmd *cobra.Command, args []string) error {
	txpool, err := client.GetTxPool(cmd.Context())
	if err != nil {
		log.Fatalln(err)
	}
	prettyPrint(txpool)
	return nil
}
