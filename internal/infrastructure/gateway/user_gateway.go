package gateway

import (
	"context"

	"github.com/555f/go-service-template/internal/domain/entity"

	"github.com/555f/go-service-template/internal/usecase/gateway"
)

var _ gateway.UserGateway = &UserGateway{}

type UserGateway struct {
}

// Search implements gateway.UserGateway.
func (u *UserGateway) Search(ctx context.Context) ([]*entity.User, error) {
	return []*entity.User{{
		FirstName: "Uncle",
		LastName:  "Bob",
	}}, nil
}

func NewUserGateway() *UserGateway {
	return &UserGateway{}
}
