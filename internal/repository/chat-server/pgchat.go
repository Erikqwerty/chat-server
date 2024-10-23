package chatserverrepository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/erikqwerty/chat-server/internal/model"
	"github.com/erikqwerty/chat-server/internal/repository"
	"github.com/erikqwerty/chat-server/internal/repository/chat-server/convertor"
	modelrepo "github.com/erikqwerty/chat-server/internal/repository/chat-server/model"
)

var _ repository.Chat = (*repoChat)(nil)

const (
	tableChat = "chats"

	chatID        = "id"
	chatName      = "chat_name"
	chatCreatedAt = "created_at"
)

type repoChat struct {
	pool *pgxpool.Pool
}

// CreateChat - сохраняет запись о новом чате  в базе данных
func (pg *repoChat) CreateChat(ctx context.Context, chatName string) (int, error) {

	query := sq.
		Insert(tableChat).
		Columns(chatName, chatCreatedAt).
		Values(chatName, time.Now()).
		Suffix("RETURNING id").PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	var chatID int
	err = pg.pool.QueryRow(ctx, sql, args...).Scan(&chatID)
	if err != nil {
		return 0, err
	}

	return chatID, nil
}

// ReadChat - возвращает информацию о чате из базы данных
func (pg *repoChat) ReadChat(ctx context.Context, id int) (*model.Chat, error) {
	query := sq.Select(chatID, chatName, chatCreatedAt).
		From(tableChat).Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	chat := &modelrepo.Chat{}
	row := pg.pool.QueryRow(ctx, sql, args...)
	err = row.Scan(&chat.ID, &chat.ChatName, &chat.CreatedAt)
	if err != nil {
		return nil, err
	}

	return convertor.ToChatFromRepo(chat), nil
}

// ReadChats - достает список чатов хранящихся в БД
func (pg *repoChat) ReadChats(ctx context.Context) ([]*model.Chat, error) {
	chats := []*modelrepo.Chat{}
	query := sq.Select(chatID, chatName, chatCreatedAt).From(tableChat).PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := pg.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		chat := &modelrepo.Chat{}
		err := rows.Scan(&chat.ID, &chat.ChatName, &chat.CreatedAt)
		if err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return convertor.ToChatsFromRepo(chats), nil
}

// DeleteChat удаляет чат по указанному (Id) из базы данных
func (pg *repoChat) DeleteChat(ctx context.Context, id int) error {
	query := sq.
		Delete(tableChat).
		Where(sq.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, execErr := pg.pool.Exec(ctx, sql, args...)
	if execErr != nil {
		return err
	}
	return nil
}
