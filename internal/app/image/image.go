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

func (img *Image) Count() (response []string) {
	images := make(map[string]bool, 1)
	for i := 0; i < len(img.manifest); i++ {
		key, err := findImage(img.manifest[i])
		if err == nil {
			images[key] = true
		}
	}
	for key := range images {
		response = append(response, key)
	}
	return response
}

func findImage(str string) (string, error) {
	matched, _ := regexp.MatchString(`(?m)\b(image):\s*([^\s]+)`, str)
	if matched {
		step0 := strings.Replace(strings.TrimSpace(str), "- ", "", 1)
		step1 := strings.ReplaceAll(strings.TrimSpace(step0), "image: ", "")
		step2 := strings.ReplaceAll(strings.TrimSpace(step1), "\"", "")
		return step2, nil
	}
	return "", errors.New("not image string")
}
