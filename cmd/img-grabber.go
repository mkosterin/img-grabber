package main

import (
	"bufio"
	"fmt"
	"img-grubber/internal/counter"
	"os"
)

func main() {
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
}
