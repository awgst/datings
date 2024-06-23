package premium

type CreatePremiumRequest struct {
	PremiumFeature string `json:"premium_feature" binding:"required,oneof=verified_label no_swipe_quota"`
}

func (CreatePremiumRequest) GetJsonFieldName(field string) string {
	return map[string]string{
		"PremiumFeature": "premium_feature",
	}[field]
}

func (CreatePremiumRequest) ErrMessages() map[string]map[string]string {
	return map[string]map[string]string{
		"premium_feature": {
			"required": "Premium feature is required",
			"oneof":    "Premium feature should be verified_label or no_swipe_quota",
		},
	}
}
