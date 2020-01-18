package cmd

import (
	"github.com/138over/sde/pkg/service"

	"github.com/spf13/cobra"
)

// ServiceFlags TODO
type ServiceFlags struct {
	ipaddress string
	lifecycle string
	name      string
	port      string
}

var serviceFlags ServiceFlags

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "yada yada",
	Long:  "yada yada service",
	RunE:  serviceRun,
}

func init() {
	serviceCmd.Flags().StringVar(&serviceFlags.name, "name", "web", "service name")
	serviceCmd.Flags().StringVar(&serviceFlags.port, "port", "3000", "service port")
	serviceCmd.Flags().StringVar(&serviceFlags.ipaddress, "ipaddress", "127.0.0.1", "service ipaddress")
	serviceCmd.Flags().StringVar(&serviceFlags.lifecycle, "lifecycle", "", "lifecycle structure file (json)")

	sdeCmd.AddCommand(serviceCmd)
}

func serviceRun(cmd *cobra.Command, args []string) error {
	service := service.Config{
		Lifecycle: serviceFlags.lifecycle,
		Name:      serviceFlags.name,
		Port:      serviceFlags.port,
		IPAddress: serviceFlags.ipaddress,
	}
	service.Start()
	return nil
}
