package cmd

import (
	"github.com/giovanniPeirone/proide/internal"
	"github.com/spf13/cobra"
)

var router string

var CrateInitProject = &cobra.Command{
	Use:   "init",
	Short: "Define th e rpject and create a file if it do not exist (in the router is void it will init in the current route)",
	Run: func(cmd *cobra.Command, args []string) {
		internal.InitProjectPOI(router)
	},
}

func init() {
	rootCmd.AddCommand(CrateInitProject)
	CrateInitProject.Flags().StringVarP(&router, "route", "r", "", "Define the route")
}
