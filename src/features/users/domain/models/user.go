package users_models

type User struct {
	ID             string
	Email          string
	HashedPassword string
}

func (s *User) String() string {
	return s.Email + " " + s.HashedPassword
}
