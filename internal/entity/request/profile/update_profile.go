package profile

type UpdateProfileRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
