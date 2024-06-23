package model

import "database/sql"

type User struct {
	ID           int
	Email        string
	PasswordHash string
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
	LastLoginAt  sql.NullTime

	Profile *Profile `gorm:"foreignKey:UserID"`
	Premium *Premium `gorm:"foreignKey:UserID"`
}

func (u User) TableName() string {
	return "users"
}

func (u User) HasUnlimitedSwipe() bool {
	if u.Premium == nil {
		return false
	}

	return u.Premium.Feature == PremiumFeatureNoSwipeQuota
}

type Profile struct {
	ID        int
	UserID    int
	Name      string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime

	// Helper
	IsVerified bool `gorm:"->"`
}

func (p Profile) TableName() string {
	return "profiles"
}
