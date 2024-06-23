package model

import "database/sql"

type PremiumFeature string

const (
	PremiumFeatureVerified     PremiumFeature = "verified_label"
	PremiumFeatureNoSwipeQuota PremiumFeature = "no_swipe_quota"
)

type Premium struct {
	ID        int
	UserID    int
	Feature   PremiumFeature
	CreatedAt sql.NullTime
}

func (Premium) TableName() string {
	return "premiums"
}
