package derocli

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jon4hz/console"
)

func (c *Console) PeersCmd() *console.Cmd {
	return &console.Cmd{
		Name:        "peers",
		Aliases:     []string{"get_peers"},
		Description: "Get all peers connected to the dero node",
		Handler:     c.handlePeersCmd,
	}
}

func (c *Console) handlePeersCmd(con *console.Console, args []string) error {
	ctx, cancel := context.WithTimeout(con.Ctx(), requestTimeout)
	defer cancel()
	peers, err := c.deroClient.GetPeers(ctx)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(peers, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
