package cmd

import (
	"fmt"
	"os"

	"github.com/adamsanghera/trees/trees"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "api",
	Short: "cli tool for launching the trees api server",
	Long:  "cli tool for launching the trees api server",
	RunE: func(cmd *cobra.Command, args []string) error {
		tSrv, err := trees.New(os.Getenv("PG_PASS"))
		if err != nil {
			return err
		}

		return tSrv.Start()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
