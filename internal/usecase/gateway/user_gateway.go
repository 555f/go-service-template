package gateway

import (
	"context"

	"github.com/555f/go-service-template/internal/domain/entity"
)

type UserGateway interface {
	Search(ctx context.Context) ([]*entity.User, error)
}
