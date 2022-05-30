package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info",
	Short:   "Get information about the node",
	PreRunE: connectNode,
	PostRun: func(cmd *cobra.Command, args []string) { client.Close() },
	RunE:    info,
}

func info(cmd *cobra.Command, args []string) error {
	info, err := client.GetInfo(context.Background())
	if err != nil {
		return err
	}
	prettyPrint(info)
	return nil
}
