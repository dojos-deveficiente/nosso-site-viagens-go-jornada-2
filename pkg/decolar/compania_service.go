package decolar

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	decolarGen "github.com/selmison/dojo-1/gen/decolar"
)

// Compania represents a single companhia.
type Compania struct {
	gorm.Model
	ID     string `gorm:"primarykey"`
	Nome   string
	PaisID string
}

// CreateCompania implements create_companhia.
func (s *service) CreateCompania(ctx context.Context, dto *decolarGen.CreateCompaniaDTO) (res *decolarGen.CompaniaDTO, err error) {
	s.logger.Log("info", fmt.Sprintf("companhia.create_companhia"))
	compania := Compania{
		ID:     uuid.New().String(),
		Nome:   dto.Nome,
		PaisID: dto.PaisID,
	}
	if result := s.repo.Create(&compania); result.Error != nil {
		return nil, result.Error
	}
	return &decolarGen.CompaniaDTO{
		ID:        compania.ID,
		Nome:      compania.Nome,
		PaisID:    compania.PaisID,
		CreatedAt: time.Now().String(),
	}, nil
}
