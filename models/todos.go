package models

import (
	"time"
)

// structを定義
type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}
