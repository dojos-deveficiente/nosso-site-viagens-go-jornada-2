package decolar

import (
	"github.com/go-kit/kit/log"
	"gorm.io/gorm"

	decolarGen "github.com/selmison/dojo-1/gen/decolar"
)

type service struct {
	repo   *gorm.DB
	logger log.Logger
}

// NewService returns the decolar service implementation.
func NewService(repo *gorm.DB, logger log.Logger) decolarGen.Service {
	return &service{repo, logger}
}
