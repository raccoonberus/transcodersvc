package repository

import "github.com/racoonberus/transcodersvc/entity"

type TaskRepository struct{}

func (r TaskRepository) Select(limit int) []entity.Task {
	// TODO: implement me
	return []entity.Task{}
}

func (r TaskRepository) Update(task entity.Task) {
	// TODO: implement me
}
