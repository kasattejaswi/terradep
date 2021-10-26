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

// Proxy config
// Command variables
var interactive bool
var proxyType string
var hostname string
var port string
var username string
var password string

var proxyCmd = &cobra.Command{
	Use:   "setProxy",
	Short: "Set proxy related configuration",
	Long:  "Set proxy related configurations",
	Run: func(cmd *cobra.Command, args []string) {
		if interactive {
			configs.StartInteractiveProxyConfigurationSession()
		} else {
			if proxyType == "" {
				printHelpWithError(cmd, "proxyType is required")
			}
			if hostname == "" {
				printHelpWithError(cmd, "hostname is required")
			}
			if port == "" {
				printHelpWithError(cmd, "port is required")
			}
			if username == "" {
				printHelpWithError(cmd, "username is required")
			}
			if password == "" {
				printHelpWithError(cmd, "password is required")
			}
			configs.SetupProxyInConfiguration(proxyType, hostname, port, username, password)
		}
	},
}

func printHelpWithError(cmd *cobra.Command, errorMessage string) {
	cmd.Help()
	log.Fatal("Error:", errorMessage)
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.AddCommand(initCmd)
	initCmd.Flags().BoolP("force", "f", false, "Force initialize empty config file")

	configCmd.AddCommand(proxyCmd)
	proxyCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Use this flag to interactively enter values")
	proxyCmd.Flags().StringVar(&proxyType, "type", "", "Type of proxy: httpProxy or httpsProxy")
	proxyCmd.Flags().StringVar(&hostname, "hostname", "", "Proxy hostname")
	proxyCmd.Flags().StringVar(&port, "port", "", "Port number for proxy")
	proxyCmd.Flags().StringVar(&username, "username", "", "Username for proxy authentication. Leave it empty in case of no auth")
	proxyCmd.Flags().StringVar(&password, "password", "", "Username for proxy authentication. Leave it empty in case of no auth")
}
