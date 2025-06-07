package repositories

import entity "github.com/osvaldotcf/pgfcycle/goexpert/apis/internal/entities"

type UsersRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
