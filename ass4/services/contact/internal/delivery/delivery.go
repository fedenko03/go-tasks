package delivery

import (
	"ass3/services/contact/internal/repository"
	"ass3/services/contact/internal/useCase"
)

type Delivery struct {
	useCase    useCase.UseCase
	repository repository.Repository
}

func NewDelivery(useCase useCase.UseCase, repository repository.Repository) *Delivery {
	return &Delivery{
		useCase:    useCase,
		repository: repository,
	}
}
