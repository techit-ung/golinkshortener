package provider

type NoFilePathError struct{}

func (NoFilePathError) Error() string {
	return "No FilePath set"
}
