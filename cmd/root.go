package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sukha-id/bee/internal/app_rest"
	"os"
)

func Start() {
	var rootCmd = &cobra.Command{Use: "sukha command"}

	var allCmd = &cobra.Command{
		Use:   "http",
		Short: "Run HTTP",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				err := errors.New("missing args")
				//TODO:add handler
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			configPath := args[0]
			app_rest.Run(configPath)
		},
	}

	rootCmd.AddCommand(allCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
