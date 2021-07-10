package project

import (
	"context"
	"golang-testing/domain"
	"golang-testing/domain/interfaces"
	"time"
)

type Service struct {
	r              interfaces.ProjectRepository
	contextTimeout time.Duration
}

func NewProjectService(r interfaces.ProjectRepository, timeout time.Duration) interfaces.ProjectUsecase {
	return &Service{
		r:              r,
		contextTimeout: timeout,
	}
}

func (s *Service) GetAll(c context.Context) ([]domain.Project, error) {

	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	listProject, err := s.r.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return listProject, nil
}

func (s *Service) GetByID(c context.Context, id int64) (*domain.Project, error) {

	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	res, err := s.r.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) Create(c context.Context, m *domain.Project) error {

	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	m.CreatedAt = time.Now()
	err := s.r.Create(ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(c context.Context, ar *domain.Project) error {

	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	ar.UpdatedAt = time.Now()
	return s.r.Update(ctx, ar)
}

func (s *Service) Delete(c context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	existedArticle, err := s.r.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existedArticle == nil {
		return domain.ErrNotFound
	}
	return s.r.Delete(ctx, id)
}
