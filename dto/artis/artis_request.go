package artisdto

type ArtisRequest struct {
	Name        string `json:"name"`
	Old         int    `json:"old"`
	Type        string `json:"type"`
	StartCareer string `json:"startCareer"`
}
