package pg

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/erikqwerty/chat-server/internal/db"
)

// ReadChat - достает чат db.Chat по его (id) из базы данных
func (pg *PG) ReadChat(ctx context.Context, id int) (*db.Chat, error) {
	query := pg.sb.Select("id", "chat_name", "created_at").From("chats").Where(squirrel.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, errSQLCreateQwery(err)
	}

	var chat db.Chat
	row := pg.pool.QueryRow(ctx, sql, args...)
	err = row.Scan(&chat.ID, &chat.ChatName, &chat.CreatedAt)
	if err != nil {
		return nil, errSQLQwery(err)
	}

	return &chat, nil
}
