package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var sdeCmd = &cobra.Command{
	Use:   "sde",
	Short: "Run sde services",
	Run:   sdeRun,
}

func sdeRun(cmd *cobra.Command, args []string) {
	// cmd.Help()
}

// Execute TODO
func Execute(args []string) {
	sdeCmd.SetArgs(args)
	if err := sdeCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
