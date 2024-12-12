package main

import (
	"os"

	"github.com/incu6us/weather-cli/cmd"
)

func main() {
	cmd.Execute(os.Args)
}
