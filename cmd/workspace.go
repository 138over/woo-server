package cmd

import (
	"github.com/spf13/cobra"
)

// WorkspaceFlags TODO
type WorkspaceFlags struct {
	name string
}

var workspaceFlags WorkspaceFlags

var workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "yada yada",
	Long:  "yada yada yada",
	RunE:  workspaceRun,
}

func init() {
	workspaceCmd.Flags().StringVar(&workspaceFlags.name, "name", ".", "name of workspace")

	sdeCmd.AddCommand(workspaceCmd)
}

func workspaceRun(cmd *cobra.Command, args []string) error {
	cmd.Help()
	return nil
}
