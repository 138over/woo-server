package cmd

import (
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test the ecosystem",
	Long:  "yada yada yada",
	RunE:  serverRun,
}

func init() {
	sdeCmd.AddCommand(testCmd)
}

func testRun(cmd *cobra.Command, args []string) error {
	cmd.Help()
	return nil
}
