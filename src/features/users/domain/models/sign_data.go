package users_models

// David?
// o es recomendable que haya
// un modelo para sign in y otro para sign up?
// en otras palabras, cada caso de uso deberia
// tener su propio modelo?

type SignData struct {
	Email          string
	Password       string
	HashedPassword string
}

func (s *SignData) String() string {
	return s.Email + " " + s.Password
}
