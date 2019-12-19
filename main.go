package main

import (
	"os"

	"github.com/138over/sde/cmd"
)

func main() {
	cmd.Execute(os.Args[1:])
}
