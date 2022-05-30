package dero

import (
	"context"
	"fmt"
	"strings"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/channel"
	"github.com/deroproject/derohe/glue/rwc"
	derop2p "github.com/deroproject/derohe/p2p"
	derorpc "github.com/deroproject/derohe/rpc"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	rpc  *jrpc2.Client
	io   *rwc.ReadWriteCloser
}

func New(addr string) (*Client, error) {
	conn, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}
	inputOutput := rwc.New(conn)
	return &Client{
		conn: conn,
		io:   inputOutput,
		rpc:  jrpc2.NewClient(channel.RawJSON(inputOutput, inputOutput), nil),
	}, nil
}

func (c *Client) Close() error {
	c.io.Close()
	return c.conn.Close()
}

func (c *Client) Echo(msg string) (string, error) {
	var res string
	if err := c.rpc.CallResult(context.Background(), "DERO.Echo", strings.Split(msg, " "), &res); err != nil {
		return "", fmt.Errorf("failed to call: %w", err)
	}
	return res, nil
}

func (c *Client) GetHeight(ctx context.Context) (*derorpc.Daemon_GetHeight_Result, error) {
	res := new(derorpc.Daemon_GetHeight_Result)
	if err := c.rpc.CallResult(context.Background(), "DERO.GetHeight", nil, res); err != nil {
		return nil, fmt.Errorf("failed to call: %w", err)
	}
	return res, nil
}

func (c *Client) GetInfo(ctx context.Context) (*derorpc.GetInfo_Result, error) {
	res := new(derorpc.GetInfo_Result)
	if err := c.rpc.CallResult(context.Background(), "DERO.GetInfo", nil, res); err != nil {
		return nil, fmt.Errorf("failed to call: %w", err)
	}
	return res, nil
}

func (c *Client) GetPeers(ctx context.Context) (*derop2p.PeersInfo, error) {
	res := new(derop2p.PeersInfo)
	if err := c.rpc.CallResult(context.Background(), "DERO.GetPeers", nil, res); err != nil {
		return nil, fmt.Errorf("failed to call: %w", err)
	}
	return res, nil
}
