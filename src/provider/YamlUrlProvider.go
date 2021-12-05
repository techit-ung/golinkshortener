package provider

type YamlUrlProvider struct {
	FilePath string
}

func (provider YamlUrlProvider) Get(key string) (string, error) {
	if provider.FilePath == "" {
		return "", NoFilePathError{}
	}

	return "http://www.google.com", nil
}
