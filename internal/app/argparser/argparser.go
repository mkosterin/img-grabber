package argparser

import (
	"flag"
	"os"
)

type Argument struct {
	Filename string
	Mode     string
	Pipe     bool
}

func New() (response *Argument) {
	response = &Argument{}
	var help = flag.Bool("help", false, "Show help")
	flag.StringVar(&response.Mode, "mode", "images", "[images|stats]")
	flag.StringVar(&response.Filename, "file", "", "<manifest filename>")

	flag.Parse()

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// we run in pipeline, but parse arguments anyway
		response.Pipe = true
	} else {
		// we run as a standalone application and don't wait anything from stdio
		response.Pipe = false
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	return
}
