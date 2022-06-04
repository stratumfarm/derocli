package derocli

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jon4hz/console"
)

func (c *Console) TxPoolCmd() *console.Cmd {
	return &console.Cmd{
		Name:        "txpool",
		Aliases:     []string{"get_txpool"},
		Description: "Get the transaction pool",
		Handler:     c.txPoolCmd,
	}
}

func (c *Console) txPoolCmd(con *console.Console, args []string) error {
	ctx, cancel := context.WithTimeout(con.Ctx(), requestTimeout)
	defer cancel()
	txpool, err := c.deroClient.GetTxPool(ctx)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(txpool, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
