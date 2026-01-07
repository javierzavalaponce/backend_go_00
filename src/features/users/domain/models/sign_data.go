package users_models

type SignData struct {
	Email          string
	Password       string
	HashedPassword string
}

func (s *SignData) String() string {
	return s.Email + " " + s.Password
}
