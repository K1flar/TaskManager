package postgres

import (
	"strconv"
	"taskmanager/internal/domain"
)

func (pr *PostgresRepo) GetUserByID(id uint32) (*domain.User, error) {
	row := pr.db.QueryRow(
		`SELECT id, name, tg_id, chat_id FROM users WHERE id=$1`,
		strconv.FormatUint(uint64(id), 10),
	)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.TgID, &user.ChatID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (pr *PostgresRepo) GetUserByTgID(tgID uint32) (*domain.User, error) {
	row := pr.db.QueryRow(
		`SELECT id, name, tg_id, chat_id FROM users WHERE tg_id=$1`,
		strconv.FormatUint(uint64(tgID), 10),
	)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.TgID, &user.ChatID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
