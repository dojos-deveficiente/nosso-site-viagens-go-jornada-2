package pais

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/google/uuid"
	"gorm.io/gorm"

	paisGen "github.com/selmison/dojo-1/gen/pais"
)

type paissrvc struct {
	repo   *gorm.DB
	logger log.Logger
}

// NewPais returns the pais service implementation.
func NewPais(repo *gorm.DB, logger log.Logger) paisGen.Service {
	return &paissrvc{repo, logger}
}

// CreatePais implements create_pais.
func (s *paissrvc) CreatePais(_ context.Context, dto *paisGen.CreatePaisDTO) (res *paisGen.PaisDTO, err error) {
	s.logger.Log("info", fmt.Sprintf("paisGen.create_pais"))
	pais := Pais{
		ID:   uuid.New().String(),
		Nome: strings.TrimSpace(dto.Nome),
	}
	if err := pais.Validate(); err != nil {
		return nil, err
	}
	res = &paisGen.PaisDTO{
		ID:   pais.ID,
		Nome: pais.Nome,
	}
	return res, nil
}
