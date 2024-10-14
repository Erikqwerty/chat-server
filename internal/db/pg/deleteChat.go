package pg

import (
	"context"

	"github.com/Masterminds/squirrel"
)

// DeleteChat удаляет чат по указанному ID
func (pg *PG) DeleteChat(ctx context.Context, id int) error {
	query := pg.sb.
		Delete("chats").
		Where(squirrel.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return errSQLCreateQwery(err)
	}

	// Выполнение запроса на удаление
	_, execErr := pg.pool.Exec(ctx, sql, args...)
	if execErr != nil {
		return errSQLQwery(err)
	}
	return nil
}
