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
	Description string
	IgnorePipe  bool
	Matcher     func(cmd string) bool
	Handler     func(c *Console, cmd string) error
	Console     *Console
}

func (c *Cmd) Match(cmd string) bool {
	return c.Matcher(cmd)
}

func (c *Cmd) Handle(cmd string) error {
	return c.Handler(c.Console, cmd)
}

var HelpCmd = &Cmd{
	Name:        "help",
	Description: "Show the help",
	Matcher:     func(cmd string) bool { return cmd == "help" },
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
	Description: "Quit the console",
	IgnorePipe:  true,
	Matcher:     func(cmd string) bool { return cmd == "quit" || cmd == "exit" },
	Handler: func(c *Console, cmd string) error {
		c.Close()
		return nil
	},
}

var ClearCmd = &Cmd{
	Name:        "clear",
	Description: "Clear the screen",
	IgnorePipe:  true,
	Matcher:     func(cmd string) bool { return cmd == "clear" },
	Handler: func(c *Console, cmd string) error {
		termenv.ClearScreen()
		return nil
	},
}

var InfoCmd = &Cmd{
	Name:        "info",
	Description: "Get info about the dero node",
	Matcher:     func(cmd string) bool { return cmd == "info" },
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
	Description: "Get info about the peers",
	Matcher:     func(cmd string) bool { return cmd == "peers" },
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
