// Code generated by GG version . DO NOT EDIT.

package middleware

import (
	"context"
	controller "github.com/555f/go-service-template/internal/usecase/controller"
	dto "github.com/555f/go-service-template/pkg/domain/dto"
)

type UserControllerMiddleware func(controller.UserController) controller.UserController

func UserControllerMiddlewareChain(outer UserControllerMiddleware, others ...UserControllerMiddleware) UserControllerMiddleware {
	return func(next controller.UserController) controller.UserController {
		for i := len(others) - 1; i >= 0; i-- {
			next = others[i](next)
		}
		return outer(next)
	}
}

type userControllerBaseMiddleware struct {
	next     controller.UserController
	mediator any
}

func (m *userControllerBaseMiddleware) Search(ctx context.Context, name string) (users []*dto.User, err error) {
	defer func() {
		if s, ok := m.mediator.(userControllerSearchBaseMiddleware); ok {
			s.Search(ctx, name)
		}
	}()
	return m.next.Search(ctx, name)
}

type userControllerSearchBaseMiddleware interface {
	Search(ctx context.Context, name string)
}

func UserControllerBaseMiddleware(mediator any) UserControllerMiddleware {
	return func(next controller.UserController) controller.UserController {
		return &userControllerBaseMiddleware{next: next, mediator: mediator}
	}
}
