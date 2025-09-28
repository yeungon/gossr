package domain

import (
	"errors"
	"time"
)

type Article struct {
	ID        int64
	Title     string
	Content   string
	CreatedAt time.Time
}

// Validate checks if the Article fields meet the required constraints.
func (a *Article) Validate() error {
	if a.Title == "" {
		return errors.New("title is required")
	}
	if len(a.Title) > 255 {
		return errors.New("title must not exceed 255 characters")
	}
	if a.Content == "" {
		return errors.New("content is required")
	}
	return nil
}

func NewArticle(title, content string) (*Article, error) {
	a := &Article{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}

	if err := a.Validate(); err != nil {
		return nil, err
	}

	return a, nil
}
