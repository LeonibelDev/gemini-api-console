package utility

import (
	"flag"
	"fmt"
	"os"
)

var version string = "1.0.0-alpha"

func Flags() {
	showVersion := flag.Bool("version", false, "get app version")
	flag.Parse()

	if *showVersion {
		fmt.Printf("version %s", version)
		os.Exit(0)
	}
}
