package manifest

import (
	"bufio"
	"log"
	"os"
)

type Manifest struct {
	Manifest []string
	fileName string
}

func New(fileName string) *Manifest {
	return &Manifest{
		Manifest: make([]string, 0),
		fileName: fileName,
	}
}

func (m *Manifest) Read() error {
	var scanner *bufio.Scanner
	if m.fileName == "" {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		file, err := os.Open(m.fileName)
		if err != nil {
			return err
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal("Can not close the file")
			}
		}(file)
		scanner = bufio.NewScanner(file)
	}
	for scanner.Scan() {
		input := scanner.Text()
		m.Manifest = append(m.Manifest, input)
	}
	return nil
}
