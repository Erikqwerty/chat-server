package chatserverrepository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/erikqwerty/chat-server/internal/model"
	"github.com/erikqwerty/chat-server/internal/repository"
	"github.com/erikqwerty/chat-server/internal/repository/chat-server/convertor"
	modelrepo "github.com/erikqwerty/chat-server/internal/repository/chat-server/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	tableMessages = "messages"

	messagesID        = "id"
	messagesChatID    = "chat_id"
	messagesUserEmail = "user_email"
	messagesText      = "text"
	messagesTimestamp = "timestamp"
)

var _ repository.Message = (*repoMessage)(nil)

type repoMessage struct {
	pool *pgxpool.Pool
}

// CreateMessage - сохраняет новое сообщение в базе данных и возвращает ID записи.
func (pg *repoMessage) CreateMessage(ctx context.Context, chatID int, userEmail, text string) (int, error) {
	query := sq.
		Insert(tableMessages).
		Columns(messagesChatID, messagesUserEmail, messagesText, messagesTimestamp).
		Values(chatID, userEmail, text, time.Now()).
		Suffix("RETURNING id").PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	var messageID int
	err = pg.pool.QueryRow(ctx, sql, args...).Scan(&messageID)
	if err != nil {
		return 0, err
	}

	return messageID, nil
}

// ReadMessages - достает из указанного чата список сообщений
func (pg *repoMessage) ReadMessages(ctx context.Context, chatID int) ([]*model.Message, error) {
	query := sq.
		Select(messagesID, messagesChatID, messagesUserEmail, messagesText, messagesTimestamp).
		From(tableMessages).
		Where(sq.Eq{membersChatID: chatID}).
		OrderBy("timestamp ASC").PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := pg.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*modelrepo.Message
	for rows.Next() {
		msg := &modelrepo.Message{}
		err := rows.Scan(&msg.ID, &msg.ChatID, &msg.UserEmail, &msg.Text, &msg.Timestamp)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return convertor.ToMessagesFromRepo(messages), nil
}

// DeleteMessage - удаляет сообщение из базы данных по его (id)
func (pg *repoMessage) DeleteMessage(ctx context.Context, id int) error {
	query := sq.Delete(tableMessages).Where(sq.Eq{messagesID: id}).PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = pg.pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
