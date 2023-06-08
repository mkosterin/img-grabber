package main

import (
	"fmt"
	"img-grubber/internal/argparser"
	"img-grubber/internal/logic"
	"img-grubber/internal/reader"
	"os"
)

func main() {

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// we run in pipeline, but parse arguments anyway
		fmt.Print(argparser.ArgParser(true))
	} else {
		// we run as a standalone application and don't wait anything from stdio
		fmt.Print(argparser.ArgParser(false))
	}

	//storage for manifest
	arr := reader.InitNewData()
	arr.StdioRead()
	//find images and normalize the list
	logic.ImageCounter(arr)

}
