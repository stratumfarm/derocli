package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	mcobra "github.com/muesli/mango-cobra"
	"github.com/muesli/roff"
	"github.com/spf13/cobra"
	"github.com/stratumfarm/derocli/internal/version"
	"github.com/stratumfarm/derocli/pkg/dero"
)

var client *dero.Client

var rootCmdFlags struct {
	rpc string
}

var rootCmd = &cobra.Command{
	Use:     "derocli",
	Short:   "A cli tool to fetch information from a dero rpc node",
	Version: version.Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd, manCmd, allCmd, heightCmd, infoCmd, peersCmd, txPoolCmd)

	rootCmd.PersistentFlags().StringVarP(&rootCmdFlags.rpc, "rpc", "r", "localhost:10102", "address of the node")
}

func Execute() error {
	return rootCmd.Execute()
}

func prettyPrint(data any) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to marshal: %w", err))
	}
	fmt.Println(string(b))
}

func connectNode(cmd *cobra.Command, args []string) error {
	c, err := dero.New("ws://" + rootCmdFlags.rpc + "/ws")
	if err != nil {
		return err
	}
	client = c
	return nil
}

var manCmd = &cobra.Command{
	Use:                   "man",
	Short:                 "generates the manpages",
	SilenceUsage:          true,
	DisableFlagsInUseLine: true,
	Hidden:                true,
	Args:                  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		manPage, err := mcobra.NewManPage(1, rootCmd)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(os.Stdout, manPage.Build(roff.NewDocument()))
		return err
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\n", version.Version)
		fmt.Printf("Commit: %s\n", version.Commit)
		fmt.Printf("Date: %s\n", version.Date)
		fmt.Printf("Build by: %s\n", version.BuiltBy)
	},
}
