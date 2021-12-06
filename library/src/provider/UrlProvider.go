package provider

type UrlProvider interface {
	Get(key string) (string, error)
}
