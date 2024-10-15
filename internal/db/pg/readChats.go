package pg

import (
	"context"
	"fmt"

	"github.com/erikqwerty/chat-server/internal/db"
)

// ReadChats - возвращает список чатов хранящихся в БД
func (pg *PG) ReadChats(ctx context.Context) ([]*db.Chat, error) {
	chats := []*db.Chat{}
	query := pg.sb.Select("id", "chat_name", "created_at").From("chats")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, errSQLCreateQwery(err)
	}

	rows, err := pg.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса на получение списка чата: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		chat := &db.Chat{}
		err := rows.Scan(&chat.ID, &chat.ChatName, &chat.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки результата: %w", err)
		}
		chats = append(chats, chat)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("ошибка обработки строк результата: %w", rows.Err())
	}

	return chats, nil
}
