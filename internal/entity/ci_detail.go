package entity

import (
	"time"
)

// Struct of ci_detail
type CIDetail struct {
	ID          int64     `json:"id"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
	SuiteName   string    `json:"suite_name"`
	CaseName    string    `json:"case_name"`
	CaseClass   string    `json:"case_class"`
	Status      *Status   `json:"status"`
	ErrorDetail string    `json:"error_detail"`
	StackTrace  string    `json:"stack_trace"`
}

// Enum type
type Status string

// Enum list...
const (
	StatusPassed = Status("passed")
	StatusFailed = Status("failed")
	StatusSkiped = Status("skiped")
)
