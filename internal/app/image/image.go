package image

import (
	"errors"
	"regexp"
	"strings"
)

type Image struct {
	manifest []string
}

func New(m []string) *Image {
	return &Image{
		manifest: m,
	}
}

func (img *Image) Count() (response map[string]bool) {
	response = make(map[string]bool, 10)
	for i := 0; i < len(img.manifest); i++ {
		key, err := findImage(img.manifest[i])
		if err == nil {
			response[key] = true
		}
	}
	return response
}

func findImage(str string) (string, error) {
	matched, _ := regexp.MatchString(`(?m)\b(image):\s*([^\s]+)`, str)
	if matched {
		step1 := strings.ReplaceAll(strings.TrimSpace(str), "image: ", "")
		step2 := strings.ReplaceAll(strings.TrimSpace(step1), "\"", "")
		return step2, nil
	}
	return "", errors.New("not image string")
}
