package interfaces

import "github.com/sminoeee/sample-app/go/domain/model"

type UserRepository interface {
	FindByID(id uint64) (*model.User, error)
	Store(user model.User) (*uint64, error)
}