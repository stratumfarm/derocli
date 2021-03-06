package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var peersCmd = &cobra.Command{
	Use:     "peers",
	Short:   "Get the connected peers from the node",
	PreRunE: connectNode,
	PostRun: func(cmd *cobra.Command, args []string) { client.Close() },
	RunE:    peers,
}

func peers(cmd *cobra.Command, args []string) error {
	info, err := client.GetPeers(cmd.Context())
	if err != nil {
		log.Fatalln(err)
	}
	prettyPrint(info)
	return nil
}
