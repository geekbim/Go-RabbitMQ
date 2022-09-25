package usecase

import (
	"context"
	"ubersnap/domain/entity"
	"ubersnap/pkg/exceptions"
)

type UserUseCase interface {
	Login(ctx context.Context, user *entity.User) (*entity.User, *exceptions.CustomerError)
}
