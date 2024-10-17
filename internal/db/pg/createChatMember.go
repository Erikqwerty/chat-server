package pg

import (
	"context"
	"time"
)

// CreateChatMember - сохраняет нового пользователя (userEmail) в таблице chat_members
// с указание чата в котором тот находится (chatID)
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
