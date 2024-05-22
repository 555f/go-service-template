package presenter

import (
	"github.com/555f/go-service-template/internal/domain/entity"

	"github.com/555f/go-service-template/internal/usecase/presenter"

	"github.com/555f/go-service-template/pkg/domain/dto"
)

var _ presenter.UserPresenter = &UserPresenter{}

type UserPresenter struct{}

// Responses implements presenter.UserPresenter.
func (p *UserPresenter) Responses(ms []*entity.User) (ds []*dto.User, err error) {
	ds = make([]*dto.User, 0, len(ms))
	for _, m := range ms {
		d, err := p.Response(m)
		if err != nil {
			return nil, err
		}
		ds = append(ds, d)
	}
	return
}

// Response implements presenter.FooPresenter.
func (p *UserPresenter) Response(m *entity.User) (d *dto.User, err error) {
	return &dto.User{
		FirstName: m.FirstName,
		LastName:  m.LastName,
	}, nil
}

func NewUserPresenter() *UserPresenter {
	return &UserPresenter{}
}
