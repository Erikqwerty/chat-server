package pg

import (
	"context"

	"github.com/Masterminds/squirrel"
)

// DeleteMessage - удаляет сообщение из базы данных по его id
func (pg *PG) DeleteMessage(ctx context.Context, id int) error {
	query := pg.sb.Delete("messages").Where(squirrel.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return errSQLCreateQwery(err)
	}

	_, err = pg.pool.Exec(ctx, sql, args...)
	if err != nil {
		return errSQLQwery(err)
	}

	return nil
}
