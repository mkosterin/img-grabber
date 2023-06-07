package main

import (
	"img-grubber/internal/argparser"
	"os"
)

func main() {

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		argparser.StdioConsumer()
	} else {
		argparser.ArgParser()
	}
	/*
		scanner := bufio.NewScanner(os.Stdin)
		writer := bufio.NewWriter(os.Stdout)
		img := *counter.NewImg()

		for scanner.Scan() {
			input := scanner.Text()
			err, key, value := counter.Parse(input)
			if err == nil {
				img[key] = value
			}
		}
		writer.Flush()
		for key, _ := range img {
			fmt.Printf("%v\n", key)
		}
	*/
}
