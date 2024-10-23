package chatserverrepository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/erikqwerty/chat-server/internal/client/db"
	"github.com/erikqwerty/chat-server/internal/model"
	"github.com/erikqwerty/chat-server/internal/repository"
	"github.com/erikqwerty/chat-server/internal/repository/chat/convertor"
	modelrepo "github.com/erikqwerty/chat-server/internal/repository/chat/model"
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
	db db.Client
}

// CreateMessage - сохраняет новое сообщение в базе данных и возвращает ID записи.
func (repo *repoMessage) CreateMessage(ctx context.Context, message *model.Message) (int, error) {
	query := sq.
		Insert(tableMessages).
		Columns(messagesChatID, messagesUserEmail, messagesText, messagesTimestamp).
		Values(message.ChatID, message.UserEmail, message.Text, time.Now()).
		Suffix("RETURNING id").PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "chat_repository_CreateMessage",
		QueryRaw: sql,
	}

	var messageID int
	err = repo.db.DB().ScanOneContext(ctx, &messageID, q, args...)
	if err != nil {
		return 0, err
	}

	return messageID, nil
}

// ReadMessages - достает из указанного чата список сообщений
func (repo *repoMessage) ReadMessages(ctx context.Context, chatID int) ([]*model.Message, error) {
	query := sq.
		Select(messagesID, messagesChatID, messagesUserEmail, messagesText, messagesTimestamp).
		From(tableMessages).
		Where(sq.Eq{membersChatID: chatID}).
		OrderBy("timestamp ASC").PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "chat_repository_ReadMessages",
		QueryRaw: sql,
	}

	var messages []*modelrepo.Message
	err = repo.db.DB().ScanAllContext(ctx, &messages, q, args...)
	if err != nil {
		return nil, err
	}

	return convertor.ToMessagesFromRepo(messages), nil
}

// DeleteMessage - удаляет сообщение из базы данных по его (id)
func (repo *repoMessage) DeleteMessage(ctx context.Context, id int) error {
	query := sq.Delete(tableMessages).Where(sq.Eq{messagesID: id}).PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository_DeleteMessage",
		QueryRaw: sql,
	}

	_, err = repo.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
