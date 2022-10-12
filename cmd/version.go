/*
Copyright © 2022 chihiros
*/
package cmd

import (
	"github.com/chihiros/monaco/asciiart"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of monaco",
	Long:  "All software has versions. This is monaco's",
	Run: func(cmd *cobra.Command, args []string) {
		aa := asciiart.NewAsciiArt([]string{"version", "0.0.1"})
		aa.Print()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
