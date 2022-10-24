package cmd

import (
	"github.com/spf13/cobra"
)

var generateCMD = &cobra.Command{
	Use:   "generate",
	Short: "generates a murum wall",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

func init() {
	RootCMD.AddCommand(generateCMD)
}
