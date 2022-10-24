package cmd

import (
	"github.com/spf13/cobra"
)

var installCMD = &cobra.Command{
	Use:   "install",
	Short: "installs the murum daemon",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

func init() {
	RootCMD.AddCommand(installCMD)
}
