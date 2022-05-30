package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var allCmd = &cobra.Command{
	Use:     "all",
	Short:   "Get all information from the node",
	PreRunE: connectNode,
	PostRun: func(cmd *cobra.Command, args []string) { client.Close() },
	RunE:    all,
}

func all(cmd *cobra.Command, args []string) error {
	height, err := client.GetHeight(context.Background())
	if err != nil {
		return err
	}
	prettyPrint(height)

	info, err := client.GetInfo(context.Background())
	if err != nil {
		return err
	}
	prettyPrint(info)

	peers, err := client.GetPeers(context.Background())
	if err != nil {
		return err
	}
	prettyPrint(peers)
	return nil
}
