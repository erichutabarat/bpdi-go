package cmd

import (
	"fmt"

	"github.com/erichutabarat/bpdi-go/internal/createmodule"
	"github.com/erichutabarat/bpdi-go/internal/createproject"
	"github.com/spf13/cobra"
)

var (
	projectName string
	moduleName  string
)

var createCMD = &cobra.Command{
	Use:   "create",
	Short: "Create Project Folder or Module",
	RunE: func(cmd *cobra.Command, args []string) error {
		out := cmd.OutOrStdout()

		if projectName != "" {
			if err := createproject.Run(projectName); err != nil {
				return err
			}
			fmt.Fprintf(out, "Project %s created!\n", projectName)
		}

		if moduleName != "" {
			if err := createmodule.Run(moduleName); err != nil {
				return err
			}
			fmt.Fprintf(out, "Module %s created!\n", moduleName)
		}

		return nil
	},
}

func init() {
	createCMD.Flags().StringVarP(
		&projectName,
		"project",
		"p",
		"",
		"Name of project",
	)

	createCMD.Flags().StringVarP(
		&moduleName,
		"module",
		"m",
		"",
		"Name of module",
	)

	createCMD.MarkFlagsMutuallyExclusive("project", "module")
	createCMD.MarkFlagsOneRequired("project", "module")

	rootCmd.AddCommand(createCMD)
}
