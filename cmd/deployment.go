package cmd

import (
	"client-go/api"
	"github.com/spf13/cobra"
)

var getDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "A brief description",
	Long: "A longer description",
	Run: func(cmd *cobra.Command, args []string) {
		api.GetDeployments()
	},
}

func init() {
	getCmd.AddCommand(getDeploymentCmd)
}
