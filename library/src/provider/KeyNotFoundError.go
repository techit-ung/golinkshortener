package provider

import "fmt"

type KeyNotFoundError struct {
	Key string
}

func (e KeyNotFoundError) Error() string {
	return fmt.Sprintf("Key: %s not found", e.Key)
}
