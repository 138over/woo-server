package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var sdeCmd = &cobra.Command{
	Use:   "sde [command]",
	Short: "yada yada",
	Long:  "yada yada yada",
	Run:   sdeRun,
}

func sdeRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}

// Execute TODO
func Execute(args []string) {
	sdeCmd.SetArgs(args)
	if err := sdeCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
