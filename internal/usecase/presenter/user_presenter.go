package presenter

import (
	"github.com/555f/go-service-template/internal/domain/entity"

	"github.com/555f/go-service-template/pkg/domain/dto"
)

type UserPresenter interface {
	Responses(ms []*entity.User) (ds []*dto.User, err error)
	Response(m *entity.User) (d *dto.User, err error)
}
