package model

import (
	"time"
)

type Borrow struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	BookID    uint
	WechatID  string
}
