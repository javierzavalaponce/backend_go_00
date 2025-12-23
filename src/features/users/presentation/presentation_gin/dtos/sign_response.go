package users_presentation_dtos

type SignResponse struct {
	Token string `json:"token_jwt" binding:"required"`
}

func NewSignResponse(token string) *SignResponse {
	return &SignResponse{Token: token}
}
