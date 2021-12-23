package repository

import (
	"tirelease/internal/entity"
)

// Interface
type CIDetailRepo interface {
	Insert(*entity.CIDetail)
	BatchInsert([]*entity.CIDetail)
}

// Implement
