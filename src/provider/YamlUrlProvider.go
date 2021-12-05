package provider

import (
	"os"

	"gopkg.in/yaml.v2"
)

type YamlUrlProvider struct {
	FilePath string
}

func (provider YamlUrlProvider) Get(key string) (string, error) {
	if provider.FilePath == "" {
		return "", NoFilePathError{}
	}

	content, err := os.ReadFile(provider.FilePath)
	if err != nil {
		return "", err
	}

	m := linkYaml{}

	yaml.Unmarshal(content, &m)

	link, ok := m.Links[key]

	if ok {
		return link, nil
	} else {
		return "", KeyNotFoundError{}
	}
}

type linkYaml struct {
	Links map[string]string
}
