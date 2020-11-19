package decolar

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	decolarGen "github.com/selmison/dojo-1/gen/decolar"
)

// Pais represents a single pais.
type Pais struct {
	gorm.Model
	ID   string
	Nome string
}

// CreatePais implements create_pais.
func (s *service) CreatePais(_ context.Context, dto *decolarGen.CreatePaisDTO) (res *decolarGen.PaisDTO, err error) {
	s.logger.Log("info", fmt.Sprintf("pais.create_pais"))
	pais := Pais{
		ID:   uuid.New().String(),
		Nome: strings.TrimSpace(dto.Nome),
	}
	if result := s.repo.Create(&pais); result.Error != nil {
		return nil, result.Error
	}
	res = &decolarGen.PaisDTO{
		ID:   pais.ID,
		Nome: pais.Nome,
	}
	return res, nil
}
