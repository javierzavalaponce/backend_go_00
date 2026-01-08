package users_presentation_dtos

type SignResponse struct {
	Token string `json:"token_jwt" binding:"required"`
}

type SignInResponse struct {
	UserId string `json:"user_id" binding:"required"`
	Token  string `json:"token_jwt" binding:"required"`
}

// javier  aqui constructor de los responses
// next: refactor sign in response
func NewSignResponse(token string) *SignResponse {
	return &SignResponse{Token: token}
}

func NewSignInResponse(userId string, token string) *SignInResponse {
	return &SignInResponse{
		UserId: userId,
		Token:  token,
	}
}
