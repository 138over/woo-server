package cmd

import (
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "yada yada",
	Long:  "yada yada yada",
	RunE:  serverRun,
}

func init() {
	sdeCmd.AddCommand(statusCmd)
}

func statusRun(cmd *cobra.Command, args []string) error {
	cmd.Help()
	return nil
}
