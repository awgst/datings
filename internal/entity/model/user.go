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
}

func (u User) TableName() string {
	return "users"
}

type Profile struct {
	ID        int
	UserID    int
	Name      string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

func (p Profile) TableName() string {
	return "profiles"
}
