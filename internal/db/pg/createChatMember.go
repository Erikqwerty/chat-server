package pg

import (
	"context"
	"time"
)

// CreateChatMember - Добавляет нового пользователя в таблицу chat_members в бд
func (pg *PG) CreateChatMember(ctx context.Context, chatID int, userEmail string) error {

	if err := checkMemberInChat(ctx, pg, chatID, userEmail); err != nil {
		return err
	}

	query := pg.sb.
		Insert("chat_members").
		Columns("chat_id", "user_email", "joined_at").
		Values(chatID, userEmail, time.Now())

	sql, arg, err := query.ToSql()
	if err != nil {
		return errSQLCreateQwery(err)
	}

	_, err = pg.pool.Exec(ctx, sql, arg...)
	if err != nil {
		return errSQLQwery(err)
	}

	return nil
}
