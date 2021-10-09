package cmd

import (
	"github.com/blackcrw/wafi/tools"
	"github.com/spf13/cobra"
)

func RootCMDRun(cmd *cobra.Command, args []string) {
	var target, _ = cmd.Flags().GetString("url")

	tools.NewDetection().SetURL(target)
}

func RootCMDPostRun(cmd *cobra.Command, args []string) {}
