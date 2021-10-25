package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "terradep",
	Short: "A terraform module dependency management solution",
	Long:  "Manages terraform dependencies and allow modules to be version control in GIT, SVN, MERCURY or some other repositories of your choice.",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
