package reader

import (
	"bufio"
	"os"
)

type Data struct {
	Text []string
}

func InitNewData() *Data {
	var response Data
	return &response
}

func (d *Data) StdioRead() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		d.Text = append(d.Text, input)
	}
}
