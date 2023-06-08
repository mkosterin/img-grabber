package argparser

import (
	"flag"
	"os"
)

type Args struct {
	mode string
	pipe bool
}

func ArgParser(pipe bool) Args {
	var help = flag.Bool("help", false, "Show help")
	var arguments Args
	flag.StringVar(&arguments.mode, "mode", "images", "[images|stats]")

	flag.Parse()

	if pipe {
		arguments.pipe = true
	} else {
		arguments.pipe = false
	}
	// Usage
	if *help {
		flag.Usage()
		os.Exit(0)
	}
	return arguments
}
