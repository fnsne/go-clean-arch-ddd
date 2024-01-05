package user

type UserFactory func(name string, email string, password string) (*User, error)
