package out

import "go-clean-arch-ddd/account/entity"

type UserRepository interface {
	NextID() string
	Save(user entity.User) error
}
