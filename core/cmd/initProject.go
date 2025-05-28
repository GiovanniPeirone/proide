package cmd

import (
	"github.com/giovanniPeirone/proide/internal"
	"github.com/spf13/cobra"
)

var route string

var CreateInitProject = &cobra.Command{
	Use:   "init",
	Short: "Initializes the project and creates a file if it doesn't exist (uses current directory if no route is provided)",
	Run: func(cmd *cobra.Command, args []string) {
		internal.InitProjectPOI(route)
	},
}

func init() {
	rootCmd.AddCommand(CreateInitProject)
	CreateInitProject.Flags().StringVarP(&route, "route", "r", "", "Define the route where to initialize the project")
}
