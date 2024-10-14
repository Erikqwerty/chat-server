package pg

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/erikqwerty/chat-server/internal/db"
)

func (pg *PG) ListChatMembers(ctx context.Context, chatID int) ([]*db.ChatMember, error) {
	// Построение SQL-запроса для получения участников чата
	query := pg.sb.
		Select("chat_id", "user_email", "joined_at").
		From("chat_members").
		Where(squirrel.Eq{"chat_id": chatID}).
		OrderBy("joined_at ASC")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, errSQLCreateQwery(err)
	}

	// Выполнение запроса и получение строк
	rows, err := pg.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса на получение участников чата: %w", err)
	}
	defer rows.Close()

	// Инициализация слайса для хранения участников
	var members []*db.ChatMember
	for rows.Next() {
		member := &db.ChatMember{}
		err := rows.Scan(&member.ChatID, &member.UserEmail, &member.JoinedAt)
		if err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки результата: %w", err)
		}
		members = append(members, member)
	}

	// Проверка на ошибки после завершения итерации по строкам
	if rows.Err() != nil {
		return nil, fmt.Errorf("ошибка обработки строк результата: %w", rows.Err())
	}

	return members, nil
}
