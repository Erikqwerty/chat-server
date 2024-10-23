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

var _ repository.ChatMember = (*repoChatMember)(nil)

const (
	tableChatMember = "chat_members"

	membersChatID    = "chat_id"
	membersUserEmail = "user_email"
	membersJoinedAt  = "joined_at"
)

type repoChatMember struct {
	pool *pgxpool.Pool
}

// CreateChatMember - сохраняет нового пользователя  в таблице chat_members
// с указание чата в котором тот находится
func (pg *repoChatMember) CreateChatMember(ctx context.Context, chatID int, userEmail string) error {

	if err := checkMemberInChat(ctx, pg, chatID, userEmail); err != nil {
		return err
	}

	query := sq.
		Insert(tableChatMember).
		Columns(membersChatID, membersUserEmail, membersJoinedAt).
		Values(chatID, userEmail, time.Now()).PlaceholderFormat(sq.Dollar)

	sql, arg, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = pg.pool.Exec(ctx, sql, arg...)
	if err != nil {
		return err
	}

	return nil
}

// ReadChatMember - достает из базы данных участника (UserEmail) чата (chatID) при наличии
func (pg *repoChatMember) ReadChatMember(ctx context.Context, UserEmail string, chatID int) (*model.ChatMember, error) {
	query := sq.
		Select(membersChatID, membersUserEmail, membersJoinedAt).
		From(tableChatMember).
		Where(sq.Eq{membersChatID: chatID, membersUserEmail: UserEmail}).PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	row := pg.pool.QueryRow(ctx, sql, args...)

	member := &modelrepo.ChatMember{}
	err = row.Scan(&member.ChatID, &member.UserEmail, &member.JoinedAt)
	if err != nil {
		return nil, err
	}

	return convertor.ToChatMemberFromRepo(member), nil
}

// ReadChatMembers - достает из базы данных список участников ([]*db.ChatMember) чата (chatID)
func (pg *repoChatMember) ReadChatMembers(ctx context.Context, chatID int) ([]*model.ChatMember, error) {
	query := sq.
		Select(membersChatID, membersUserEmail, membersJoinedAt).
		From(tableChatMember).
		Where(sq.Eq{membersChatID: chatID}).
		OrderBy("joined_at ASC").PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := pg.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []*modelrepo.ChatMember
	for rows.Next() {
		member := &modelrepo.ChatMember{}
		err := rows.Scan(&member.ChatID, &member.UserEmail, &member.JoinedAt)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return convertor.ToChatMembersFromRepo(members), nil
}

// DeleteChatMember - удаляет участника чата по его (userEmail), в указаном чате (chatID)
func (pg *repoChatMember) DeleteChatMember(ctx context.Context, chatID int, userEmail string) error {
	query := sq.Delete(tableChatMember).Where(sq.Eq{membersChatID: chatID, membersUserEmail: userEmail}).PlaceholderFormat(sq.Dollar)

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

// checkMemberInChat - проверяет, состоит ли пользователь в чате
func checkMemberInChat(ctx context.Context, pg *repoChatMember, chatID int, userEmail string) error {
	checkQuery := sq.
		Select("1").
		From(tableChatMember).
		Where(sq.Eq{membersChatID: chatID, membersUserEmail: userEmail}).PlaceholderFormat(sq.Dollar)

	sql, args, err := checkQuery.ToSql()
	if err != nil {
		return err
	}

	var exists int
	err = pg.pool.QueryRow(ctx, sql, args...).Scan(&exists)
	if err == nil {
		return err
	} else if err.Error() == "no rows in result set" {
		return nil
	}

	return err
}
