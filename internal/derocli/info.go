package derocli

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jon4hz/console"
)

func (c *Console) InfoCmd() *console.Cmd {
	return &console.Cmd{
		Name:        "info",
		Aliases:     []string{"get_info"},
		Description: "Get info about the dero node",
		Handler:     c.handleInfoCmd,
	}
}

func (c *Console) handleInfoCmd(con *console.Console, args []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	info, err := c.deroClient.GetInfo(ctx)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(info, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
