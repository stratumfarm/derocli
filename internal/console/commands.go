package console

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/muesli/termenv"
)

var requestTimeout = 5 * time.Second

var Cmds = []*Cmd{
	HelpCmd,
	QuitCmd,
	ClearCmd,
	InfoCmd,
	PeersCmd,
}

type Cmd struct {
	Name        string
	Aliases     []string
	Description string
	IgnorePipe  bool
	Matcher     func(cmd string) bool
	Handler     func(c *Console, cmd string) error
	Console     *Console
}

func (c *Cmd) defaultMatcher(cmd string) bool {
	if cmd == c.Name {
		return true
	}
	for _, alias := range c.Aliases {
		if cmd == alias {
			return true
		}
	}
	return false
}

func (c *Cmd) Match(cmd string) bool {
	if c.defaultMatcher(cmd) {
		return true
	}
	if c.Matcher != nil {
		return c.Matcher(cmd)
	}
	return false
}

func (c *Cmd) Handle(cmd string) error {
	return c.Handler(c.Console, cmd)
}

var HelpCmd = &Cmd{
	Name:        "help",
	Description: "Show the help",
	Handler: func(c *Console, cmd string) error {
		fmt.Println(helpView(c))
		return nil
	},
}

func helpView(c *Console) string {
	s := "Available commands:"
	for _, cmd := range c.cmds {
		if cmd.Name != "" && cmd.Description != "" {
			s += fmt.Sprintf("\n  %s - %s", cmd.Name, cmd.Description)
		}
	}
	return s
}

var QuitCmd = &Cmd{
	Name:        "quit",
	Aliases:     []string{"exit"},
	Description: "Quit the console",
	IgnorePipe:  true,
	Handler: func(c *Console, cmd string) error {
		c.Close()
		return nil
	},
}

var ClearCmd = &Cmd{
	Name:        "clear",
	Description: "Clear the screen",
	IgnorePipe:  true,
	Handler: func(c *Console, cmd string) error {
		termenv.ClearScreen()
		return nil
	},
}

var InfoCmd = &Cmd{
	Name:        "info",
	Aliases:     []string{"get_info"},
	Description: "Get info about the dero node",
	Handler:     handleInfoCmd,
}

func handleInfoCmd(c *Console, cmd string) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	info, err := c.client.GetInfo(ctx)
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

var PeersCmd = &Cmd{
	Name:        "peers",
	Aliases:     []string{"get_peers"},
	Description: "Get info about the peers",
	Handler:     handlePeersCmd,
}

func handlePeersCmd(c *Console, cmd string) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	peers, err := c.client.GetPeers(ctx)
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
