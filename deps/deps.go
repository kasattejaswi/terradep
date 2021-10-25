package deps

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type dcf struct {
	Metadata struct {
		Name    string `yaml: "name"`
		Version string `yaml: "version"`
	} `yaml: "metadata"`
	Dependencies []struct {
		Name    string `yaml: "name"`
		Version string `yaml: "version"`
	} `yaml: "dependencies"`
}

func ReadDcf(filePath string) (*dcf, error) {
	buff, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	d := &dcf{}
	err = yaml.Unmarshal(buff, d)

	if err != nil {
		return nil, err
	}
	return d, nil
}
