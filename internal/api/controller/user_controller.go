package controller

import (
	"context"

	"github.com/555f/go-service-template/internal/usecase/controller"
	"github.com/555f/go-service-template/internal/usecase/user"

	"github.com/555f/go-service-template/pkg/domain/dto"
)

var _ controller.UserController = &UserController{}

type UserController struct {
	searchUserUseCase user.SearchUserUseCase
}

// Hello implements controller.UserController.
func (c *UserController) Search(ctx context.Context, name string) (users []*dto.User, err error) {
	return c.searchUserUseCase.Execute(ctx)
}

func NewUserController(
	searchUserUseCase user.SearchUserUseCase,
) *UserController {
	return &UserController{
		searchUserUseCase: searchUserUseCase,
	}
}
