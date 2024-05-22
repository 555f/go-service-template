package controller

import (
	"context"

	"github.com/555f/go-service-template/pkg/domain/dto"
)

// UserController
// @gg:"http"
// @gg:"middleware"
// @gg:"slog"
// @http-type:"chi"
// @http-openapi
// @http-api-doc
// @http-openapi-tags:"user"
// @http-client
// @http-server
type UserController interface {
	// Search Поиск и список пользователей
	// @http-path:"/users"
	// @http-method:"GET"
	// @http-content-type:"json"
	Search(
		ctx context.Context,
		// @http-name:""name
		// @http-type:"query"
		name string,
	) (users []*dto.User, err error)
}
