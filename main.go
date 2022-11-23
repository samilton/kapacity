package main

import (
	"fmt"
	"os"
)

var (
	version = "dev"
	commit  = "0000000000000"
	date    = "unknown"
	builtBy = "unknown"
)

func main() {
	os.Stderr.WriteString(fmt.Sprintf("Kapacity %s [%s] build on %s", version, commit[0:6], date))
}
