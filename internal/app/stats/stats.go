package stat

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

type Stat struct {
	manifest []string
}

func New(m []string) *Stat {
	return &Stat{
		manifest: m,
	}
}

func (s *Stat) ObjectCounter() (keys []string, response map[string]uint) {
	response = make(map[string]uint)
	for i := 0; i < len(s.manifest); i++ {
		buffer, err := findKind(s.manifest[i])
		if err == nil {
			response[buffer] = response[buffer] + 1
		}
	}
	keys = make([]string, 0, len(response))
	for key := range response {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return
}

func findKind(str string) (string, error) {
	matched, _ := regexp.MatchString(`^.*kind: [A-Z].*`, str)
	substr := "kind: "
	if matched {
		index := strings.Index(str, substr)
		step1 := str[index+len(substr):]
		//step1 := strings.ReplaceAll(strings.TrimSpace(str), "image: ", "")
		//step2 := strings.ReplaceAll(strings.TrimSpace(step1), "\"", "")
		return step1, nil
	}
	return "", errors.New("not Kind string")
}
