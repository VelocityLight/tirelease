package repository

import (
	"tirelease/internal/entity"
)

// Interface
type CIDetailRepo interface {
	Insert()
	BatchInsert([]*entity.CIDetail)
}

// Implement
