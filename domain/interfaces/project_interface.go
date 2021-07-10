package interfaces

import (
	"context"
	"golang-testing/domain"
)

type ProjectUsecase interface {
	GetAll(ctx context.Context) ([]domain.Project, error)
	GetByID(ctx context.Context, id int64) (*domain.Project, error)
	Create(context.Context, *domain.Project) error
	Update(ctx context.Context, ar *domain.Project) error
	Delete(ctx context.Context, id int64) error
}

type ProjectRepository interface {
	GetAll(ctx context.Context) (res []domain.Project, err error)
	GetByID(ctx context.Context, id int64) (*domain.Project, error)
	Update(ctx context.Context, ar *domain.Project) error
	Create(ctx context.Context, a *domain.Project) error
	Delete(ctx context.Context, id int64) error
}
