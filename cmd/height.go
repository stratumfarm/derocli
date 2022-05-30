package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var heightCmd = &cobra.Command{
	Use:     "height",
	Short:   "Get the current height of the blockchain",
	PreRunE: connectNode,
	PostRun: func(cmd *cobra.Command, args []string) { client.Close() },
	RunE:    height,
}

func height(cmd *cobra.Command, args []string) error {
	height, err := client.GetHeight(context.Background())
	if err != nil {
		return err
	}
	prettyPrint(height)
	return nil
}
