package cmd

import (
	"github.com/giovanniPeirone/proide/internal"
	"github.com/spf13/cobra"
)

var command string

var RunCommand = &cobra.Command{
	Use:   "run",
	Short: "run a comand in the poi file",
	Run: func(cmd *cobra.Command, args []string) {
		internal.RunCommand(command)
	},
}

func init() {
	rootCmd.AddCommand(RunCommand)
	CreateInitProject.Flags().StringVarP(&command, "route", "r", "", "Define the route where to initialize the project")
}
