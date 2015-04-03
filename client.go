package palmetto

import "errors"

var ErrNotFound = errors.New("not found")

type Client interface {
	// Get gets the value of the given key
	Get(string) (*PlainResult, error)

	// Put sets the value of the given key
	Put(string, []byte) (*PlainResult, error)
}
