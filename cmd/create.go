package cmd

import (
	"fmt"
	"os"

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
		if projectName != "" {
			file, err := os.Create(projectName)
			if err != nil {
				return err
			}
			defer file.Close()

			fmt.Fprintf(cmd.OutOrStdout(), "Project %s created!\n", projectName)
		}

		if moduleName != "" {
			err := os.MkdirAll(moduleName, 0755)
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Module %s created!\n", moduleName)
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
