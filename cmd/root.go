package cmd

import (
	"os"

	"github.com/chihiros/monaco/asciiart"
	"github.com/harakeishi/curver"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "monaco",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	RunE: func(cmd *cobra.Command, args []string) error {
		version, err := cmd.Flags().GetBool("version")
		if err != nil {
			return err
		}
		if version {
			monacoVersion()
			return nil
		}

		return nil
	},
}

func monacoVersion() {
	aa := asciiart.NewAsciiArt([]string{curver.GetVersion()})
	aa.Name = "mona_back"
	aa.Print()
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "echo version")
}
