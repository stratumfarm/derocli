package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
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

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := c.Start(); err != nil {
			log.Fatalln(err)
		}
	}()
	defer c.Close()
	go func() {
		<-done
		cancel()
	}()
	wg.Wait()

	return nil
}
