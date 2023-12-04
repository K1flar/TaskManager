package repo

import "taskmanager/internal/domain"

type TaskRepo interface {
	GetAll() ([]*domain.Task, error)
	GetByAssignee(assigneeID uint32) ([]*domain.Task, error)
	GetByAuthor(authorID uint32) ([]*domain.Task, error)
	Add(task *domain.Task) error
	Assign(id uint32, assignee string, assigneeID uint32) error
	Unassign(id uint32) error
	Resolve(id uint32) error
}
