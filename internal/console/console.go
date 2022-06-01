package console

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/muesli/termenv"
	"github.com/stratumfarm/derocli/pkg/dero"
)

type Opts func(*Console)

func WithClient(client *dero.Client) Opts {
	return func(c *Console) {
		c.client = client
	}
}

type Console struct {
	ctx    context.Context
	cancel context.CancelFunc

	stdin  *os.File
	isPipe bool
	stdout *os.File
	stderr *os.File

	cmds   []*Cmd
	client *dero.Client
}

func New(opts ...Opts) (*Console, error) {
	ctx, cancel := context.WithCancel(context.Background())
	c := &Console{
		ctx:    ctx,
		cancel: cancel,
		stdin:  os.Stdin,
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
	for _, opt := range opts {
		opt(c)
	}
	ok, err := fileIsPipe(c.stdin)
	if err != nil {
		return nil, err
	}
	c.isPipe = ok

	c.registerCommands(
		HelpCmd,
		QuitCmd,
		ClearCmd,
	)

	return c, nil
}

func fileIsPipe(in *os.File) (bool, error) {
	if fi, _ := in.Stat(); (fi.Mode() & os.ModeNamedPipe) != 0 {
		return true, nil
	}
	return false, nil
}

func (c *Console) registerCommands(cmds ...*Cmd) {
	for _, cmd := range cmds {
		cmd.Console = c
	}
	c.cmds = cmds
}

func (c *Console) Start(ctx context.Context) error {
	if !c.isPipe {
		termenv.AltScreen()
		termenv.ClearScreen()
		defer termenv.ExitAltScreen()
	}
	c.welcomeMsg()
	return c.read(ctx)
}

func (c *Console) Close() error {
	c.cancel()
	return nil
}

func (c *Console) welcomeMsg() {
	fmt.Fprintf(c.stdout, "Welcome to dero-cli!\n> ")
}

func (c *Console) read(ctx context.Context) error {
	if c.isPipe {
		msg, err := c.readFromPipe()
		if err != nil {
			fmt.Fprintln(c.stderr, "reading pipe input:", err)
		}
		c.stdout.Write([]byte(msg))
		return c.handleInput(msg)
	}
	scanner := bufio.NewScanner(c.stdin)
	inch := make(chan string)
	go func() {
		defer close(inch)
		for scanner.Scan() {
			inch <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			if err == io.EOF {
				return
			}
			fmt.Fprintln(c.stderr, "reading standard input:", err)
		}
	}()
	for {
		select {
		case in, ok := <-inch:
			if !ok {
				return nil
			}
			if err := c.handleInput(in); err != nil {
				fmt.Fprintln(c.stderr, err)
			}
		case <-ctx.Done():
			return nil
		case <-c.ctx.Done():
			return nil
		}
	}
}

func (c *Console) readFromPipe() (string, error) {
	defer c.stdin.Close()
	data, err := io.ReadAll(c.stdin)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (c *Console) handleInput(input string) error {
	input = strings.TrimSpace(input)
	if !c.isPipe {
		defer fmt.Fprintf(c.stdout, "\n> ")
	}
	for _, cmd := range c.cmds {
		if cmd.Match(input) {
			cmd.Handle(input)
		}
	}
	return nil
}
