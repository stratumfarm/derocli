package derocli

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jon4hz/console"
)

func (c *Console) HeightCmd() *console.Cmd {
	return &console.Cmd{
		Name:        "height",
		Aliases:     []string{"get_height"},
		Description: "Get the height of the dero node",
		Handler:     c.heightCmd,
	}
}

func (c *Console) heightCmd(con *console.Console, args []string) error {
	ctx, cancel := context.WithTimeout(con.Ctx(), requestTimeout)
	defer cancel()
	height, err := c.deroClient.GetHeight(ctx)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(height, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
