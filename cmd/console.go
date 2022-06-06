package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/stratumfarm/derocli/internal/derocli"
)

var consoleCmd = &cobra.Command{
	Use:     "console",
	Short:   "Start an interactive console",
	PreRunE: connectNode,
	PostRun: func(cmd *cobra.Command, args []string) { client.Close() },
	RunE:    runConsole,
}

func runConsole(cmd *cobra.Command, args []string) error {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(cmd.Context())
	defer cancel()

	c, err := derocli.New(ctx, client)
	if err != nil {
		log.Fatalln(err)
	}

	errc := make(chan error, 1)
	go func() {
		defer close(errc)
		defer c.Close()
		if err := c.Start(); err != nil {
			errc <- err
		}
	}()

	select {
	case <-done:
		cancel()
	case err, ok := <-errc:
		if !ok {
			break
		}
		log.Fatalln(err)
	}

	return nil
}
