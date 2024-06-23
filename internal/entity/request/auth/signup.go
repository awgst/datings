package request

type SignupRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (SignupRequest) GetJsonFieldName(field string) string {
	return map[string]string{
		"Email":    "email",
		"Password": "password",
	}[field]
}

func (SignupRequest) ErrMessages() map[string]map[string]string {
	return map[string]map[string]string{
		"email": {
			"required": "Email is required",
		},
		"password": {
			"required": "Password is required",
		},
	}
}
