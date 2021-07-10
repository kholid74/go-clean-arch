package domain

import (
	"errors"
	"time"
)

type Project struct {
	ID             int64
	Title          string
	Description    string
	Image          string
	NumberOfTester int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (p *Project) Validate() error {

	if p.Title == "" {
		return errors.New("Required Title")
	}
	if p.Description == "" {
		return errors.New("Required Description")
	}

	return nil
}
