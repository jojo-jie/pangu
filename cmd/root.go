package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(cleanCmd)
	rootCmd.AddCommand(installCmd)
}

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

func Execute() error {
	return rootCmd.Execute()
}
