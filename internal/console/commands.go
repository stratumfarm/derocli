package console

import (
	"fmt"

	"github.com/muesli/termenv"
)

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
	if c.Console.isPipe && c.IgnorePipe {
		return nil
	}
	return c.Handler(c.Console, cmd)
}

var HelpCmd = &Cmd{
	Name:        "help",
	Description: "Show the help",
	Matcher:     func(cmd string) bool { return cmd == "help" },
	Handler: func(c *Console, cmd string) error {
		fmt.Fprintln(c.stdout, helpView(c))
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
		c.cancel()
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
