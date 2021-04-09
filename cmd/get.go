package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description",
	Long: "A longer description",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}
func init() {
	rootCmd.AddCommand(getCmd)
}
