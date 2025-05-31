package repositories

import entity "github.com/OsvaldoTCF/pgFCycle/goexpert/apis/internal/entities"

type UsersRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
