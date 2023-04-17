package counter

import (
	"errors"
	"regexp"
	"strings"
)

type Img map[string]bool

func NewImg() *Img {
	resp := make(Img)
	return &resp
}

func Parse(str string) (error, string, bool) {
	matched, _ := regexp.MatchString(`(?m)\b(image):\s*([^\s]+)`, str)
	if matched {
		step1 := strings.ReplaceAll(strings.TrimSpace(str), "image: ", "")
		step2 := strings.ReplaceAll(strings.TrimSpace(step1), "\"", "")
		return nil, step2, false
	}
	return errors.New("not image string"), "", false
}
