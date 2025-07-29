package pokeapi

import (
	"errors"
	"fmt"
)

var (
	ErrUndefinedResponse = errors.New("empty or undefined response")
	ErrUnknownCommand    = errors.New("unknown command")
	ErrStatusNotOK       = func(statusCode int) error { return fmt.Errorf("bad status code: %d", statusCode) }
)
