package dog

import (
	"errors"
	"strings"
)

type Dog struct {
	Name string `json:"name"`
}

// Create a constructor for Dog
func New(name string) (*Dog, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if len(name) > 50 {
		return nil, errors.New("name is too long (max 50 characters)")
	}

	return &Dog{
		Name: name,
	}, nil
}

func (d *Dog) GetName() string {
	return d.Name
}

func (d *Dog) Speak() string {
	return "Gau gau!"
}
