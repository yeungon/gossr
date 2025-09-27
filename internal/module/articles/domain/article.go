package domain

import "time"

type Article struct {
	ID        int64
	Title     string
	Content   string
	CreatedAt time.Time
}
