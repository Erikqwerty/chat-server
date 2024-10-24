package chatrepo

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/erikqwerty/chat-server/internal/client/db"
	"github.com/erikqwerty/chat-server/internal/model"
	"github.com/erikqwerty/chat-server/internal/repository"

	"github.com/erikqwerty/chat-server/internal/repository/chatrepo/convertor"
	"github.com/erikqwerty/chat-server/internal/repository/chatrepo/modelrepo"
)

var _ repository.Chat = (*repoChat)(nil)

const (
	tableChat = "chats"

	chatID        = "id"
	chatName      = "chat_name"
	chatCreatedAt = "created_at"
)

type repoChat struct {
	db db.Client
}

// CreateChat - сохраняет запись о новом чате  в базе данных
func (repo *repoChat) CreateChat(ctx context.Context, chat string) (int, error) {
	query := sq.
		Insert(tableChat).
		Columns(chatName, chatCreatedAt).
		Values(chat, time.Now()).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "chat_repository_CreateChat",
		QueryRaw: sql,
	}

	var chatID int

	err = repo.db.DB().ScanOneContext(ctx, &chatID, q, args...)
	if err != nil {
		return 0, err
	}

	return chatID, nil
}

// ReadChat - возвращает информацию о чате из базы данных
func (repo *repoChat) ReadChat(ctx context.Context, id int) (*model.Chat, error) {
	query := sq.
		Select(chatID, chatName, chatCreatedAt).
		From(tableChat).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "chat_repository_ReadChat",
		QueryRaw: sql,
	}

	chat := &modelrepo.Chat{}

	err = repo.db.DB().ScanOneContext(ctx, chat, q, args...)
	if err != nil {
		return nil, err
	}

	return convertor.ToChatFromRepo(chat), nil
}

// ReadChats - достает список чатов хранящихся в БД
func (repo *repoChat) ReadChats(ctx context.Context) ([]*model.Chat, error) {
	chats := []*modelrepo.Chat{}

	query := sq.
		Select(chatID, chatName, chatCreatedAt).
		From(tableChat).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "chat_repository_ReadChats",
		QueryRaw: sql,
	}

	err = repo.db.DB().ScanAllContext(ctx, &chats, q, args...)
	if err != nil {
		return nil, err
	}

	return convertor.ToChatsFromRepo(chats), nil
}

// DeleteChat удаляет чат по указанному (Id) из базы данных
func (repo *repoChat) DeleteChat(ctx context.Context, id int) error {
	query := sq.
		Delete(tableChat).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository_DeleteChat",
		QueryRaw: sql,
	}

	_, execErr := repo.db.DB().ExecContext(ctx, q, args...)
	if execErr != nil {
		return err
	}

	return nil
}
