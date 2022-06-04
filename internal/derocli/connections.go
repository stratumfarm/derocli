package derocli

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jon4hz/console"
)

func (c *Console) ConnectionsCmd() *console.Cmd {
	return &console.Cmd{
		Name:        "connections",
		Aliases:     []string{"get_connections"},
		Description: "Get all connections from the dero node",
		Handler:     c.handleConnectionsCmd,
	}
}

func (c *Console) handleConnectionsCmd(con *console.Console, args []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	connections, err := c.deroClient.GetConnections(ctx)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(connections.Connections, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
