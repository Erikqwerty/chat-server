package pg

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/erikqwerty/chat-server/internal/db"
)

// ReadChatMember - достает из базы данных участника (UserEmail) чата (chatID) при наличии
func (pg *PG) ReadChatMember(ctx context.Context, UserEmail string, chatID int) (*db.ChatMember, error) {
	query := pg.sb.
		Select("chat_id", "user_email", "joined_at").
		From("chat_members").
		Where(squirrel.Eq{"chat_id": chatID, "user_email": UserEmail})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("ошибка создания SQL-запроса: %w", err)
	}

	row := pg.pool.QueryRow(ctx, sql, args...)

	member := &db.ChatMember{}
	err = row.Scan(&member.ChatID, &member.UserEmail, &member.JoinedAt)
	if err != nil {
		return nil, fmt.Errorf("ошибка сканирования результата: %w", err)
	}

	return member, nil
}
