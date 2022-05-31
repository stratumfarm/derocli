package cmd

import (
	"context"
	"time"

	"github.com/spf13/cobra"
)

var timeout = time.Second * 5

var allCmd = &cobra.Command{
	Use:     "all",
	Short:   "Get all information from the node",
	PreRunE: connectNode,
	PostRun: func(cmd *cobra.Command, args []string) { client.Close() },
	RunE:    all,
}

func all(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	height, err := client.GetHeight(ctx)
	if err != nil {
		return err
	}
	prettyPrint(height)

	info, err := client.GetInfo(ctx)
	if err != nil {
		return err
	}
	prettyPrint(info)

	/* peers, err := client.GetPeers(ctx)
	if err != nil {
		return err
	}
	prettyPrint(peers) */

	txPool, err := client.GetTxPool(ctx)
	if err != nil {
		return err
	}
	prettyPrint(txPool)
	return nil
}
