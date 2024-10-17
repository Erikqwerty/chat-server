package pg

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

// checkMemberInChat - проверяет, состоит ли пользователь (userEmail) в чате (chatID) в базе данных
func checkMemberInChat(ctx context.Context, pg *PG, chatID int, userEmail string) error {
	checkQuery := pg.sb.
		Select("1").
		From("chat_members").
		Where(squirrel.Eq{"chat_id": chatID, "user_email": userEmail})

	sql, args, err := checkQuery.ToSql()
	if err != nil {
		return fmt.Errorf("ошибка построения SQL-запроса: %w", err)
	}

	var exists int
	err = pg.pool.QueryRow(ctx, sql, args...).Scan(&exists)
	if err == nil {
		return errUserAlreadyExists(err)
	} else if err.Error() == "no rows in result set" {
		return nil
	}

	return errSQLQwery(err)
}
