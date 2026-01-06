package users_models

// David? deberia
// o es recomendable que haya
// un modelo para sign in y otro para sign up?
type SignData struct {
	Email          string
	Password       string
	HashedPassword string
}

func (s *SignData) String() string {
	return s.Email + " " + s.Password
}
