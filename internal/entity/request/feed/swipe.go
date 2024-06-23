package feed

type SwipeRequest struct {
	ProfileID int    `json:"profile_id" binding:"required"`
	Type      string `json:"type" binding:"required,oneof=pass like"`
}

func (SwipeRequest) GetJsonFieldName(field string) string {
	return map[string]string{
		"ProfileID": "profile_id",
		"Type":      "type",
	}[field]
}

func (SwipeRequest) ErrMessages() map[string]map[string]string {
	return map[string]map[string]string{
		"profile_id": {
			"required": "Profile ID is required",
		},
		"type": {
			"required": "Swipe type is required",
			"oneof":    "Swipe type must be pass or like",
		},
	}
}
