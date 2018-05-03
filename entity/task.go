package entity

import "time"

type Task struct {
	BaseEntity

	ResourceID       int64
	Resource         Resource
	CallbackUrl      string
	CallbackResponse string
	StartedAt        time.Time
	FinishedAt       time.Time
	Output           string
}
