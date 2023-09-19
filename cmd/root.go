package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sukha-id/bee/cmd/http"
	"os"
)

func Start() {
	var rootCmd = &cobra.Command{Use: "4dw command"}

	var allCmd = &cobra.Command{
		Use:   "http",
		Short: "Run HTTP",
		Run: func(cmd *cobra.Command, args []string) {
			http.Start()
		},
	}

	rootCmd.AddCommand(allCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
