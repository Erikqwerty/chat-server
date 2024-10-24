package pg

import (
	"context"
	"time"
)

// CreateMessage - сохраняет новое сообщение (text) с указанием чата (chatID) и отправителя (userEmail)
// в базе данных и возвращает ID записи.
func (pg *PG) CreateMessage(ctx context.Context, chatID int, userEmail, text string) (int, error) {
	query := pg.sb.
		Insert("messages").
		Columns("chat_id", "user_email", "text", "timestamp").
		Values(chatID, userEmail, text, time.Now()).
		Suffix("RETURNING id")

	sql, args, err := query.ToSql()
	if err != nil {
		return 0, errSQLCreateQwery(err)
	}

	var messageID int
	err = pg.pool.QueryRow(ctx, sql, args...).Scan(&messageID)
	if err != nil {
		return 0, errSQLQwery(err)
	}

	return messageID, nil
}
