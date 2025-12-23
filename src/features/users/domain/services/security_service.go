package users_services

type SecurityService interface {
	GenerateToken(userId string) (string, error)
	GeneratePasswordHash(password string) (string, error)
}
