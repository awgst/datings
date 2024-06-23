package feed

import "github.com/awgst/datings/internal/entity/model"

type RecommendationResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (r RecommendationResponse) Makes(profile []model.Profile) []RecommendationResponse {
	res := make([]RecommendationResponse, 0, len(profile))
	for _, p := range profile {
		res = append(res, RecommendationResponse{
			ID:   p.ID,
			Name: p.Name,
		})
	}

	return res
}
