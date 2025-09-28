package domain

import (
	"errors"
	"time"
)

type Category struct {
	ID        int64
	Name      string
	CreatedAt time.Time
}

// Validate checks the Category fields.
func (c *Category) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if len(c.Name) > 255 {
		return errors.New("name must not exceed 255 characters")
	}
	return nil
}

func NewCategory(name string) (*Category, error) {
	c := &Category{
		Name:      name,
		CreatedAt: time.Now(),
	}
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return c, nil
}
