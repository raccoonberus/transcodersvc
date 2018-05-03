package entity

import "time"

type BaseEntity struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	RemovedAt time.Time
}
