package cmd

import (
	"github.com/spf13/cobra"
)

// ServerFlags TODO
type ServerFlags struct {
	port string
}

var serverFlags ServerFlags

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "yada yada",
	Long:  "yada yada yada",
	RunE:  serverRun,
}

func init() {
	serverCmd.Flags().StringVar(&serverFlags.port, "port", "1380", "web server port")

	sdeCmd.AddCommand(serverCmd)
}

func serverRun(cmd *cobra.Command, args []string) error {
	cmd.Help()
	return nil
}
