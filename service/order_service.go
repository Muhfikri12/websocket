package service

import (
	"project/domain"
	"project/repository"
)

type OrderService interface {
	All(page, limit uint) (int, int, []domain.OrderTotal, error)
	Update(order *domain.Order) error
	Get(orderId uint) (domain.OrderTotal, error)
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo: repo}
}

func (s *orderService) All(page, limit uint) (int, int, []domain.OrderTotal, error) {
	return s.repo.All(page, limit)
}

func (s *orderService) Update(order *domain.Order) error {
	return s.repo.Update(order)
}

func (s *orderService) Get(orderId uint) (domain.OrderTotal, error) {
	return s.repo.Get(orderId)
}
