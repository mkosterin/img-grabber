package yamlparser

import (
	"gopkg.in/yaml.v2"
)

type YamlParser struct {
	manifest     []string
	listManifest []map[interface{}]interface{}
}

func New(m []string) *YamlParser {
	return &YamlParser{
		manifest:     m,
		listManifest: make([]map[interface{}]interface{}, 0),
	}
}

func (y *YamlParser) ParseYamlManifest() {
	// full y.listManifest list of maps. Each map contain one manifest
	manifest := y.splitManifestToBytes()
	for i := range manifest {
		var m map[interface{}]interface{}
		err := yaml.Unmarshal(manifest[i], &m)
		if err != nil {
			panic(err)
		}
		y.listManifest = append(y.listManifest, m)
	}
	return
}

func (y *YamlParser) splitManifestToBytes() (response [][]byte) {
	// this function get []string contains full manifest and split it to [][]byte.
	// Each element of [][]byte contains one yaml manifest
	buffer := make([]byte, 0)
	for i := range y.manifest {
		if y.manifest[i] == "---" {
			if len(buffer) > 0 {
				response = append(response, buffer)
			}
			buffer = make([]byte, 0)
			continue
		}
		buffer = append(buffer, y.manifest[i]+"\n"...)
	}
	response = append(response, buffer) //last manifest should be also append to response

	return
}
