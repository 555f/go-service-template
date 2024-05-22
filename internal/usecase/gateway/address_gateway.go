package gateway

import (
	"context"

	"github.com/555f/go-service-template/internal/domain/entity"
)

type AddressGateway interface {
	Save(ctx context.Context, address entity.Address)
}
