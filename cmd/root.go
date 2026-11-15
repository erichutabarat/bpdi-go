package cmd

import (
	"os"

	"github.com/erichutabarat/bpdi-go/internal/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "bpdi-go",
	Short:        "bpdi-go is cli tools to make boilerplate and dependency injections for golang project",
	Version:      version.Version,
	SilenceUsage: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate("bpdi-go version {{.Version}}\n")
	rootCmd.Flags().BoolP("version", "v", false, "show cli version information")
}
