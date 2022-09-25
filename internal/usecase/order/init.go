package order

import (
	"ubersnap/domain/repository"
	"ubersnap/domain/usecase"
)

type orderInteractor struct {
	orderRepository   repository.OrderRepository
	productRepository repository.ProductRepository
}

func NewOrderInteractor(
	orderRepository repository.OrderRepository,
	productRepository repository.ProductRepository,
) usecase.OrderUseCase {
	return &orderInteractor{
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
}
