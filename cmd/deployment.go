package cmd

import (
	"client-go/api"
	"github.com/spf13/cobra"
)

var name, image string
var replica int32

var getDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "A brief description",
	Long: "A longer description",
	Run: func(cmd *cobra.Command, args []string) {
		api.GetDeployments()
	},
}

var createDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "A brief description",
	Long: "A longer description",
	Run: func(cmd *cobra.Command, args []string) {
		api.CreateDeployment()
	},
}

var deleteDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "A brief description",
	Long: "A longer description",
	Run: func(cmd *cobra.Command, args []string) {
		api.DeleteDeployment(args)
	},
}

var updateDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "A brief description",
	Long: "A longer description",
	Run: func(cmd *cobra.Command, args []string) {
		api.UpdateDeployment(name, replica, image)
	},
}

func init() {
	getCmd.AddCommand(getDeploymentCmd)
	createCmd.AddCommand(createDeploymentCmd)
	deleteCmd.AddCommand(deleteDeploymentCmd)
	updateCmd.AddCommand(updateDeploymentCmd)

	updateCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Deployment Name")
	updateCmd.PersistentFlags().Int32VarP(&replica, "replica", "r", 2, "Replica Count")
	updateCmd.PersistentFlags().StringVarP(&image, "image", "i", "", "Deployment Image")

}
