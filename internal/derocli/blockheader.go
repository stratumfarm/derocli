package derocli

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/jon4hz/console"
)

func (c *Console) BlockHeaderByTopoHeight() *console.Cmd {
	return &console.Cmd{
		Name:        "block_header_by_topo_height",
		Aliases:     []string{"get_block_header_by_topo_height"},
		Description: "Get the height of the dero node",
		Handler:     c.blockHeaderByTopoHeightCmd,
	}
}

func (c *Console) blockHeaderByTopoHeightCmd(con *console.Console, args []string) error {
	if len(args) != 1 {
		return errors.New("exactly one arg is required")
	}
	height := uint64(0)
	if len(args) > 0 {
		height, _ = strconv.ParseUint(args[0], 10, 64)
	}
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	blockHeader, err := c.deroClient.GetBlockHeaderByTopoHeight(ctx, height)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(blockHeader, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
