package decolar

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	decolarGen "github.com/selmison/dojo-1/gen/decolar"
)

// Aeroporto represents a single aeroporto.
type Aeroporto struct {
	gorm.Model
	ID     string `gorm:"primarykey"`
	Nome   string
	PaisID string
}

// CreateAeroporto implements create_aeroporto.
func (s *service) CreateAeroporto(ctx context.Context, dto *decolarGen.CreateAeroportoDTO) (res *decolarGen.AeroportoDTO, err error) {
	s.logger.Log("info", fmt.Sprintf("companhia.create_companhia"))
	aeroporto := Aeroporto{
		ID:   uuid.New().String(),
		Nome: strings.TrimSpace(dto.Nome),
	}
	if result := s.repo.Create(&aeroporto); result.Error != nil {
		return nil, result.Error
	}
	res = &decolarGen.AeroportoDTO{
		ID:   aeroporto.ID,
		Nome: aeroporto.Nome,
	}
	return res, nil
}
