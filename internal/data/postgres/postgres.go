package postgres

import (
	"database/sql"
	"log"
	"strconv"
	"taskmanager/internal/domain"
	"taskmanager/internal/repo"
)

var _ repo.TaskRepo = &TaskRepository{}

type TaskRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (tr *TaskRepository) GetAll() ([]*domain.Task, error) {
	rows, err := tr.db.Query(`SELECT id, about, author, author_id, assignee, assignee_id FROM tasks`)
	if err != nil {
		return nil, err
	}
	tasks := []*domain.Task{}
	for rows.Next() {
		task := &domain.Task{}
		err = rows.Scan(&task.ID, &task.About, &task.Author, &task.AuthorID, &task.Assignee, &task.AssigneeID)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (tr *TaskRepository) GetByAssignee(assigneeID uint32) ([]*domain.Task, error) {
	rows, err := tr.db.Query(
		`SELECT id, about, author, author_id, assignee, assignee_id FROM tasks WHERE assignee_id=$1`,
		strconv.FormatUint(uint64(assigneeID), 10),
	)
	if err != nil {
		return nil, err
	}

	tasks := []*domain.Task{}
	for rows.Next() {
		task := &domain.Task{}
		err = rows.Scan(&task.ID, &task.About, &task.Author, &task.AuthorID, &task.Assignee, &task.AssigneeID)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (tr *TaskRepository) GetByAuthor(authorID uint32) ([]*domain.Task, error) {
	rows, err := tr.db.Query(
		`SELECT id, about, author, author_id, assignee, assignee_id FROM tasks WHERE author_id=$1`,
		strconv.FormatUint(uint64(authorID), 10),
	)
	if err != nil {
		return nil, err
	}

	tasks := []*domain.Task{}
	for rows.Next() {
		task := &domain.Task{}
		err = rows.Scan(&task.ID, &task.About, &task.Author, &task.AuthorID, &task.Assignee, &task.AssigneeID)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (tr *TaskRepository) Add(task *domain.Task) error {
	res, err := tr.db.Exec(
		`INSERT INTO tasks (about, author, author_id, assignee, assignee_id) VALUES ($1, $2, $3, $4, $5)`,
		task.About,
		task.Author, strconv.FormatUint(uint64(task.AuthorID), 10),
		task.Assignee, strconv.FormatUint(uint64(task.AssigneeID), 10),
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("Rows affected: ", rows)

	return nil
}

func (tr *TaskRepository) Assign(id uint32, assignee string, assigneeID uint32) error {
	res, err := tr.db.Exec(
		`UPDATE tasks SET assignee=$1, assignee_id=$2 WHERE id=$3`,
		assignee, strconv.FormatUint(uint64(assigneeID), 10), strconv.FormatUint(uint64(id), 10),
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("Rows affected: ", rows)

	return nil
}

func (tr *TaskRepository) Unassign(id uint32) error {
	res, err := tr.db.Exec(
		`UPDATE tasks SET assignee='', assignee_id=0 WHERE id=$1`,
		strconv.FormatUint(uint64(id), 10),
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("Rows affected: ", rows)

	return nil
}

func (tr *TaskRepository) Resolve(id uint32) error {
	res, err := tr.db.Exec(
		`DELETE FROM tasks WHERE id=$1`,
		strconv.FormatUint(uint64(id), 10),
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("Rows affected: ", rows)

	return nil
}
