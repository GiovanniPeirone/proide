package internal

import (
	"fmt"
	"os"
)

type path struct {
	namePath string `json:namePath`
	path     string `json:path`
}

func RunCommand(command string) {

	JsonPath := "../path-projects/projects.json"

	file, err := os.Open(JsonPath)
	if err != nil {
		fmt.Println("Error while reading the file", err)
		return
	}
	defer file.Close()

	fmt.Println(":", command)
}
