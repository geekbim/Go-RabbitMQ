package repository

import (
	"context"
	"ubersnap/domain/entity"
)

type UserRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
}
