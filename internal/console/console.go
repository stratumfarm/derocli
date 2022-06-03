package console

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/peterh/liner"
	"github.com/stratumfarm/derocli/pkg/dero"
)

var historyFn = filepath.Join(os.TempDir(), ".derocli_history")

type Opts func(*Console)

func WithClient(client *dero.Client) Opts {
	return func(c *Console) {
		c.client = client
	}
}

func WithContext(ctx context.Context) Opts {
	return func(c *Console) {
		c.parentCtx = ctx
	}
}

type Console struct {
	parentCtx context.Context
	ctx       context.Context
	cancel    context.CancelFunc
	liner     *liner.State

	cmds   []*Cmd
	client *dero.Client
}

func New(opts ...Opts) (*Console, error) {
	c := &Console{
		parentCtx: context.Background(),
		liner:     liner.NewLiner(),
	}
	for _, opt := range opts {
		opt(c)
	}

	ctx, cancel := context.WithCancel(c.parentCtx)
	c.ctx = ctx
	c.cancel = cancel

	c.registerCommands(
		HelpCmd,
		QuitCmd,
		ClearCmd,
		InfoCmd,
	)
	c.setCompleter()

	c.liner.SetCtrlCAborts(true)

	return c, nil
}

func (c *Console) registerCommands(cmds ...*Cmd) {
	for _, cmd := range cmds {
		cmd.Console = c
	}
	c.cmds = cmds
}

func (c *Console) Start() error {
	c.readHistory()
	c.welcomeMsg()
	return c.read()
}

func (c *Console) Close() error {
	c.cancel()
	c.writeHistory()
	c.liner.Close()
	return nil
}

func (c *Console) setCompleter() {
	c.liner.SetCompleter(func(line string) (s []string) {
		for _, n := range c.cmds {
			if strings.HasPrefix(n.Name, strings.ToLower(line)) {
				s = append(s, n.Name)
			}
		}
		return
	})
}

func (c *Console) welcomeMsg() {
	fmt.Println("Welcome to derocli!\n> ")
}

func (c *Console) read() error {
	doneC := make(chan struct{})
	go func() {
		defer close(doneC)
		for {
			if in, err := c.liner.Prompt("> "); err == nil {
				c.liner.AppendHistory(in)
				if err := c.handleInput(in); err != nil {
					fmt.Println(styleError.Render(err.Error()))
				}
				if QuitCmd.Matcher(in) { // special case to prevent an unnecessary newline
					break
				}
			} else if err == liner.ErrPromptAborted {
				fmt.Println("Aborted")
				break
			} else if err == io.EOF {
				break
			} else {
				fmt.Println(styleError.Render(fmt.Sprintf("Error reading line: %s", err)))
				break
			}
		}
	}()

	select {
	case <-doneC:
		return nil
	case <-c.ctx.Done():
		return nil
	}
}

func (c *Console) readHistory() {
	f, err := os.Open(historyFn)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println(styleError.Render(fmt.Sprintf("Error opening history file: %s", err)))
	}
	defer f.Close()
	if _, err := c.liner.ReadHistory(f); err != nil {
		fmt.Println(styleError.Render(fmt.Sprintf("Error reading history file: %s", err)))
	}
}

func (c *Console) writeHistory() {
	f, err := os.Create(historyFn)
	if err != nil {
		fmt.Println(styleError.Render(fmt.Sprintf("Error creating history file: %s", err)))
	}
	defer f.Close()
	if _, err := c.liner.WriteHistory(f); err != nil {
		fmt.Println(styleError.Render(fmt.Sprintf("Error writing history file: %s", err)))
	}
}

func (c *Console) handleInput(input string) error {
	input = strings.TrimSpace(input)
	for _, cmd := range c.cmds {
		if cmd.Match(input) {
			if err := cmd.Handle(input); err != nil {
				fmt.Println(styleError.Render(fmt.Sprintf("error running command %s: %s\n", cmd.Name, err)))
			}
			return nil
		}
	}
	return nil
}
