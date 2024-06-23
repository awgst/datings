package model

import (
	"database/sql"
)

type SwipeType string

const (
	SwipeTypePass SwipeType = "pass"
	SwipeTypeLike SwipeType = "like"
)

type Swipe struct {
	ID        int
	UserID    int
	ProfileID int
	Type      SwipeType
	CreatedAt sql.NullTime
}

func (Swipe) TableName() string {
	return "swipes"
}
