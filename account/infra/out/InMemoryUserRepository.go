package out

import (
	"github.com/lithammer/shortuuid/v4"
	"go-clean-arch-ddd/account/entity"
)

type InMemoryUserRepository struct {
	userMap map[string]entity.User
}

func (i *InMemoryUserRepository) FindByID(id string) (entity.User, error) {
	user, ok := i.userMap[id]
	if !ok {
		return entity.User{}, entity.ErrUserNotFound
	}
	return user, nil
}

func (i *InMemoryUserRepository) NextID() string {
	return shortuuid.New()
}

func (i *InMemoryUserRepository) Save(user entity.User) error {
	i.userMap[user.ID] = user
	return nil
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{userMap: make(map[string]entity.User)}
}
