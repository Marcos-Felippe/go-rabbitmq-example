package usecase

import (
	"github.com/projetosgo/exemplocompleto/internal/order/entity"
)

type GetTotalOutputDTO struct {
	Total int
}

type GetTotalUseCase struct {
	OrderReposotory entity.OrderRepositoryInterface
}

func NewGetTotalUseCase(orderRepository entity.OrderRepositoryInterface) *GetTotalUseCase {
	return &GetTotalUseCase{
		OrderReposotory: orderRepository,
	}
}

func (c *GetTotalUseCase) Execute() (*GetTotalOutputDTO, error) {
	total, err := c.OrderReposotory.GetTotal()
	if err != nil {
		return nil, err
	}
	return &GetTotalOutputDTO{
		Total: total,
	}, nil
}
