package domain

type User struct {
	ID     uint32 `sql:"id"`
	Name   string `sql:"name"`
	TgID   uint32 `sql:"tg_id"`
	ChatID uint32 `sql:"chat_id"`
}
