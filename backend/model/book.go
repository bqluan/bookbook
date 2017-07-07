package model

import (
	"time"
)

type Book struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Author    string
	Pub       string
	Desc      string
	Qty       int
}
