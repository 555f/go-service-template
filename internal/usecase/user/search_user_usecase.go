package user

import (
	"context"

	"github.com/555f/go-service-template/internal/usecase/gateway"
	"github.com/555f/go-service-template/internal/usecase/presenter"

	"github.com/555f/go-service-template/pkg/domain/dto"
)

type SearchUserUseCase interface {
	Execute(ctx context.Context) ([]*dto.User, error)
}

var _ SearchUserUseCase = &DefaultSearchUserUseCase{}

type DefaultSearchUserUseCase struct {
	userGateway   gateway.UserGateway
	userPresenter presenter.UserPresenter
}

func (u *DefaultSearchUserUseCase) Execute(ctx context.Context) ([]*dto.User, error) {
	users, err := u.userGateway.Search(ctx)
	if err != nil {
		return nil, err
	}
	return u.userPresenter.Responses(users)
}

func NewDefaultSearchUserUseCase(
	userGateway gateway.UserGateway,
	userPresenter presenter.UserPresenter,
) *DefaultSearchUserUseCase {
	return &DefaultSearchUserUseCase{
		userGateway:   userGateway,
		userPresenter: userPresenter,
	}
}
