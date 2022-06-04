package derocli

import (
	"context"
	"encoding/json"
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

func (c *Console) InfoCmd() *console.Cmd {
	return &console.Cmd{
		Name:        "info",
		Aliases:     []string{"get_info"},
		Description: "Get info about the dero node",
		Handler:     c.handleInfoCmd,
	}
}

func (c *Console) handleInfoCmd(con *console.Console, cmd string) error {
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

func (c *Console) PeersCmd() *console.Cmd {
	return &console.Cmd{
		Name:        "peers",
		Aliases:     []string{"get_peers"},
		Description: "Get all peers connected to the dero node",
		Handler:     c.handlePeersCmd,
	}
}

func (c *Console) handlePeersCmd(con *console.Console, cmd string) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
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

func (c *Console) ConnectionsCmd() *console.Cmd {
	return &console.Cmd{
		Name:        "connections",
		Aliases:     []string{"get_connections"},
		Description: "Get all connections from the dero node",
		Handler:     c.handleConnectionsCmd,
	}
}

func (c *Console) handleConnectionsCmd(con *console.Console, cmd string) error {
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
