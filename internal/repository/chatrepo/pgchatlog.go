package chatrepo

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/erikqwerty/chat-server/internal/model"
	"github.com/erikqwerty/chat-server/internal/repository"
	"github.com/erikqwerty/chat-server/pkg/db"
)

var _ repository.RepoLoger = (*repoLoger)(nil)

const (
	tableLogs = "user_log"

	actionType      = "action_type"
	actionDetails   = "action_details"
	actionTimestamp = "action_timestamp"
)

type repoLoger struct {
	db db.Client
}

// CreateLog - записываем лог
func (r *repoLoger) CreateLog(ctx context.Context, log *model.Log) error {
	query := sq.
		Insert(tableLogs).
		Columns(actionType, actionDetails, actionTimestamp).
		Values(log.ActionType, log.ActionDetails, sq.Expr("NOW()")).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository_CreateLog",
		QueryRaw: sql,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
