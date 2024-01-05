package user

type Repository interface {
	Create(user *User) (id string, err error)
	Get(id string) (*User, error)
	Edit(id string, user *User) error
	Delete(id string) error
	GetByEmail(email string) (*User, error)
}
