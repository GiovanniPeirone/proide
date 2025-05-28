package main

import (
	"github.com/giovanniPeirone/proide/cmd"
)

// poi list --> list of pojects
// poi init route --> init the project with poi
// poi del id or name --> Delete porject from the list
// poi go id or name --> Go to the projec
// poi run "command"  -- Runs the run comand of the project

func main() {
	cmd.Execute()
}
