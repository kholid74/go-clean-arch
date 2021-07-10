package mysql

import (
	"context"
	"golang-testing/domain"

	"gorm.io/gorm"
)

type mysqlProjectRepository struct {
	db *gorm.DB
}

func NewMysqlProjectRepository(db *gorm.DB) *mysqlProjectRepository {
	return &mysqlProjectRepository{
		db: db,
	}
}

func (m *mysqlProjectRepository) Create(ctx context.Context, a *domain.Project) error {

	data := domain.Project{
		Title:          a.Title,
		Description:    a.Description,
		Image:          a.Image,
		NumberOfTester: a.NumberOfTester,
	}

	tx := m.db.WithContext(ctx).
		Create(&data)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (m *mysqlProjectRepository) Update(ctx context.Context, a *domain.Project) error {
	var project domain.Project

	checkProjectID := m.db.WithContext(ctx).
		Select("*").
		Where("id = ?", a.ID).
		Find(&project)

	if checkProjectID.RowsAffected < 1 {
		return checkProjectID.Error
	}

	project.Title = a.Title
	project.Description = a.Description
	project.Image = a.Image
	project.NumberOfTester = a.NumberOfTester

	data := m.db.WithContext(ctx).
		Select("*").
		Where("id = ?", a.ID).
		Find(&project).
		Updates(&project)

	if data.Error != nil {
		return data.Error
	}

	return nil

}

func (m *mysqlProjectRepository) Delete(ctx context.Context, id int64) error {
	var project domain.Project

	checkProjectID := m.db.WithContext(ctx).
		Select("*").
		Where("id = ?", id).
		Find(&project)

	if checkProjectID.RowsAffected < 1 {
		return checkProjectID.Error
	}

	data := m.db.WithContext(ctx).
		Select("*").
		Where("id = ?", id).
		Find(&project).
		Delete(&project)

	if data.Error != nil {
		return data.Error
	}

	return nil
}

func (m *mysqlProjectRepository) GetByID(ctx context.Context, id int64) (*domain.Project, error) {
	var project domain.Project

	tx := m.db.WithContext(ctx).
		Where("projects.id = ?", id).
		First(&project)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &project, nil
}

func (m *mysqlProjectRepository) GetAll(ctx context.Context) ([]domain.Project, error) {
	var project []domain.Project

	tx := m.db.WithContext(ctx).
		Find(&project)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return project, nil
}
