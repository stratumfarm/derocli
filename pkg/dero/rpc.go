package dero

import (
	"context"
	"fmt"
	"strings"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/channel"
	"github.com/gorilla/websocket"
	"github.com/stratumfarm/derohe/glue/rwc"
	derorpc "github.com/stratumfarm/derohe/rpc"
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

func (c *Client) GetPeers(ctx context.Context) (*derorpc.GetPeersResult, error) {
	res := new(derorpc.GetPeersResult)
	if err := c.rpc.CallResult(context.Background(), "DERO.GetPeers", nil, res); err != nil {
		return nil, fmt.Errorf("failed to call: %w", err)
	}
	return res, nil
}

func (c *Client) GetTxPool(ctx context.Context) (*derorpc.GetTxPool_Result, error) {
	res := new(derorpc.GetTxPool_Result)
	if err := c.rpc.CallResult(context.Background(), "DERO.GetTxPool", nil, res); err != nil {
		return nil, fmt.Errorf("failed to call: %w", err)
	}
	return res, nil
}

func (c *Client) GetConnections(ctx context.Context) (*derorpc.GetConnectionResult, error) {
	res := new(derorpc.GetConnectionResult)
	if err := c.rpc.CallResult(context.Background(), "DERO.GetConnections", nil, res); err != nil {
		return nil, fmt.Errorf("failed to call: %w", err)
	}
	return res, nil
}

func (c *Client) GetBlockHeaderByTopoHeight(ctx context.Context, height uint64) (*derorpc.GetBlockHeaderByHeight_Result, error) {
	res := new(derorpc.GetBlockHeaderByHeight_Result)
	req := derorpc.GetBlockHeaderByTopoHeight_Params{
		TopoHeight: height,
	}
	if err := c.rpc.CallResult(context.Background(), "DERO.GetBlockHeaderByTopoHeight", req, res); err != nil {
		return nil, fmt.Errorf("failed to call: %w", err)
	}
	return res, nil
}
