package models

import (
	"time"
)

// Userのstructを定義
type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
	Todos     []Todo
}

// Sessionのstructを定義
type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
}
