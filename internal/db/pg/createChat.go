package pg

import (
	"context"
	"time"
)

// CreateChat - сохраняет запись о новом чате (chatName) в базе данных и возвращает его id
func (pg *PG) CreateChat(ctx context.Context, chatName string) (int, error) {

	query := pg.sb.
		Insert("chats").
		Columns("chat_name", "created_at").
		Values(chatName, time.Now()).
		Suffix("RETURNING id")

	sql, args, err := query.ToSql()
	if err != nil {
		return 0, errSQLCreateQwery(err)
	}

	var chatID int
	err = pg.pool.QueryRow(ctx, sql, args...).Scan(&chatID)
	if err != nil {
		return 0, errSQLQwery(err)
	}

	return chatID, nil
}
