package pg

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/erikqwerty/chat-server/internal/db"
)

// ReadMessages - достает из указанного чата (chatID) список сообщений в формате ([]*db.Message)
func (pg *PG) ReadMessages(ctx context.Context, chatID int) ([]*db.Message, error) {
	query := pg.sb.
		Select("id", "chat_id", "user_email", "text", "timestamp").
		From("messages").
		Where(squirrel.Eq{"chat_id": chatID}).
		OrderBy("timestamp ASC")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, errSQLCreateQwery(err)
	}

	rows, err := pg.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, errSQLQwery(err)
	}
	defer rows.Close()

	var messages []*db.Message
	for rows.Next() {
		msg := &db.Message{}
		err := rows.Scan(&msg.ID, &msg.ChatID, &msg.UserEmail, &msg.Text, &msg.Timestamp)
		if err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки результата: %w", err)
		}
		messages = append(messages, msg)
	}

	// Проверка на ошибки после завершения итерации по строкам
	if rows.Err() != nil {
		return nil, fmt.Errorf("ошибка обработки строк результата: %w", rows.Err())
	}

	return messages, nil
}
