package argparser

import (
	"flag"
	"fmt"
	"os"
)

type args struct {
	arg1 string
	arg2 string
}

var help = flag.Bool("help", false, "Show help")
var boolFlag = false
var stringFlag = "Hello There!"
var intFlag int
var arguments args

func ArgParser() {
	flag.BoolVar(&boolFlag, "boolFlag", false, "A boolean flag")
	flag.StringVar(&stringFlag, "stringFlag", "Hello There!", "A string flag")
	flag.IntVar(&intFlag, "intFlag", 4, "An integer flag")
	flag.StringVar(&arguments.arg1, "arg1", "", "structArg1")
	flag.StringVar(&arguments.arg2, "arg2", "", "structArg2")

	flag.Parse()

	// Usage Demo
	if *help {
		flag.Usage()
		os.Exit(0)
	}
	fmt.Print(arguments)
}
