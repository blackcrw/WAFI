package cmd

import (
	"github.com/blackcrw/akumascan/tools"
	"github.com/spf13/cobra"
)

func RootCMDRun(cmd *cobra.Command, args []string) {
	var target, _ = cmd.Flags().GetString("url")

	var ddd = tools.NewDetection()
	ddd.SetURL(target)
	ddd.RunnerPassive()
}

func RootCMDPostRun(cmd *cobra.Command, args []string) {}
