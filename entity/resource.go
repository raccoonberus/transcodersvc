package entity

type Resource struct {
	BaseEntity

	Size             int64
	OriginalFilename string
	InternalFilename string
}
