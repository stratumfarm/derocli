package derocli

import (
	"context"
	"fmt"
	"time"

	"github.com/jon4hz/console"
	"github.com/stratumfarm/derocli/pkg/dero"
)

var requestTimeout = 5 * time.Second

type Console struct {
	deroClient *dero.Client
	console    *console.Console
}

func New(ctx context.Context, client *dero.Client) (*Console, error) {
	if client == nil {
		return nil, fmt.Errorf("client is nil")
	}
	con, err := console.New(
		console.WithWelcomeMsg("Welcome to derocli!"),
		console.WithHandleCtrlC(true),
		console.WithContext(ctx),
	)
	if err != nil {
		return nil, err
	}
	c := &Console{
		console:    con,
		deroClient: client,
	}
	if err := con.RegisterCommands(c.getCmds()...); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Console) Start() error {
	return c.console.Start()
}

func (c *Console) Close() error {
	if err := c.deroClient.Close(); err != nil {
		return err
	}
	return c.console.Close()
}

func (c *Console) getCmds() []*console.Cmd {
	return []*console.Cmd{
		c.InfoCmd(),
		c.PeersCmd(),
		c.ConnectionsCmd(),
	}
}
