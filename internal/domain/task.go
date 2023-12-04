package domain

type Task struct {
	ID         uint32 `sql:"id"`
	About      string `sql:"about"`
	Author     string `sql:"author"`
	AuthorID   uint32 `sql:"author_id"`
	Assignee   string `sql:"assignee"`
	AssigneeID uint32 `sql:"assignee_id"`
}
