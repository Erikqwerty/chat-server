package pg

import (
	"context"

	"github.com/Masterminds/squirrel"
)

// DeleteChatMember - удаляет участника чата по его (userEmail), в указаном чате (chatID)
func (pg *PG) DeleteChatMember(ctx context.Context, chatID int, userEmail string) error {
	query := pg.sb.Delete("chat_members").Where(squirrel.Eq{"chat_id": chatID, "user_email": userEmail})

	sql, args, err := query.ToSql()
	if err != nil {
		return errSQLCreateQwery(err)
	}

	_, execErr := pg.pool.Exec(ctx, sql, args...)
	if execErr != nil {
		return errSQLQwery(err)
	}
	return nil
}
