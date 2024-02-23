package cmd

import (
	"fmt"
	"os"

	"github.com/actanonvebra/gopix/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "My CLI Application",
	Run: func(cmd *cobra.Command, args []string) {
		// Ana komutun çalıştırılacak kodu burada olacak
		fmt.Println("Welcome to my CLI Application!")
	},
}

var createConfigCmd = &cobra.Command{
	Use:   "create-config",
	Short: "Create a config file, type:yaml",
	Run: func(cmd *cobra.Command, args []string) {
		clientID, _ := cmd.Flags().GetString("client-id")
		clientSecret, _ := cmd.Flags().GetString("client-secret")
		config.CreateConfigFile(clientID, clientSecret)
	},
}

func init() {
	rootCmd.AddCommand(createConfigCmd)
	createConfigCmd.Flags().StringP("client-id", "i", "", "Client ID")
	createConfigCmd.Flags().StringP("client-secret", "s", "", "Client Secret")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
