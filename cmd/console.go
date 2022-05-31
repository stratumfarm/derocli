package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stratumfarm/derocli/internal/console"
)

var consoleCmd = &cobra.Command{
	Use:     "console",
	Short:   "Start an interactive console",
	PreRunE: connectNode,
	PostRun: func(cmd *cobra.Command, args []string) { client.Close() },
	RunE:    runConsole,
}

func runConsole(cmd *cobra.Command, args []string) error {
	c := console.New(console.WithClient(client))
	return c.Read()
	// return nil
}
