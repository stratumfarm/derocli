package console

import (
	"bufio"
	"fmt"
	"os"

	"github.com/stratumfarm/derocli/pkg/dero"
)

type Opts func(*Console)

func WithClient(client *dero.Client) Opts {
	return func(c *Console) {
		c.client = client
	}
}

type Console struct {
	stdin  *os.File
	stdout *os.File
	stderr *os.File

	client *dero.Client
}

func New(opts ...Opts) *Console {
	c := &Console{
		stdin:  os.Stdin,
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Console) Read() error {
	scanner := bufio.NewScanner(c.stdin)
	for scanner.Scan() {
		c.handleInput(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(c.stderr, "reading standard input:", err)
	}
	return nil
}

func (c *Console) handleInput(input string) error {
	return nil
}
