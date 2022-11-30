package main

import (
	"fmt"
	"os"

	"github.com/samilton/kapacity/cmd/kapacity/commands"
)

var (
	version   = "dev"
	gitCommit = "0000000000000"
	date      = "unknown"
)

func main() {
	os.Stderr.WriteString(fmt.Sprintf("Kapacity %s [%s] build on %s\n", version, gitCommit[0:6], date))

	commands.Execute()
}
