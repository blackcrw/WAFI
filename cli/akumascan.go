package cli

import (
	"log"
	"os"
	"strings"

	"github.com/blackcrw/akumascan/cli/cmd"
	"github.com/blackcrw/akumascan/internal"
	"github.com/blackcrw/akumascan/pkg/nettools"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:     "akumascan",
	Short:   "A K U M A  S C A N",
	Long:    internal.TextBanner() + `Akuma Scan (Web Application Firewall / Intrusion)`,
	Run:     cmd.RootCMDRun,
	PostRun: cmd.RootCMDPostRun,
}

func init() {
	cobra.OnInitialize(checks_lists)

	root.PersistentFlags().StringP("url", "u", "", "Target URL (Ex: http(s)://example.com/). ")

	root.MarkPersistentFlagRequired("url")
}

func checks_lists() {
	var target, err = root.Flags().GetString("url")

	if err != nil { log.Fatalln("OOOOOO") }

	internal.SimpleBanner()

	if !strings.HasSuffix(target, "/") { target = target + "/" }
	if !nettools.URLValidate(target) { log.Fatalln("This is URL not validate") }
}

func Execute() {
	if err := root.Execute(); err != nil { os.Exit(0) }
}
