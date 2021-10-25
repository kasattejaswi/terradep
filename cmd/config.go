package cmd

import (
	"log"

	"github.com/kasattejaswi/terradep/configs"
	"github.com/spf13/cobra"
)

// Root config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Setup configuration for terradep",
	Long: `Use this command to setup configuration which will be used by this
commandline in order to resolve modules as dependencies from different repositories.

It creates an encrypted file under ~/.terradep.
This file can be edited directly updating repositories, 
usernames and passwords. To edit directly, file must be decrypted first`,
}

// Init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize an empty configuration in user's home directory",
	Long:  "Initialize an empty configuration in user's home directory",
	Run: func(cmd *cobra.Command, args []string) {
		fStatus, err := cmd.Flags().GetBool("force")
		if err != nil {
			log.Fatal("Error occurred: ", err)
		}
		if fStatus {
			configs.ForceInitializeConfig()
		} else {
			configs.InitializeConfig()
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(initCmd)
	initCmd.Flags().BoolP("force", "f", false, "Force initialize empty config file")
	// initCmd.Flags().Bool("toggle", "t", false, "Help message for toggle")
}
