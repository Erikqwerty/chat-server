package chatserverrepository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/erikqwerty/chat-server/internal/client/db"
	"github.com/erikqwerty/chat-server/internal/model"
	"github.com/erikqwerty/chat-server/internal/repository"
	"github.com/erikqwerty/chat-server/internal/repository/chat/convertor"

	modelrepo "github.com/erikqwerty/chat-server/internal/repository/chat/model"
)

var _ repository.ChatMember = (*repoChatMember)(nil)

const (
	tableChatMember = "chat_members"

	membersChatID    = "chat_id"
	membersUserEmail = "user_email"
	membersJoinedAt  = "joined_at"
)

type repoChatMember struct {
	db db.Client
}

// CreateChatMember - сохраняет нового пользователя  в таблице chat_members
// с указание чата в котором тот находится
func (repo *repoChatMember) CreateChatMember(ctx context.Context, member *model.ChatMember) error {

	if err := checkMemberInChat(ctx, repo, member); err != nil {
		return err
	}

	query := sq.
		Insert(tableChatMember).
		Columns(membersChatID, membersUserEmail, membersJoinedAt).
		Values(member.ChatID, member.UserEmail, member.JoinedAt).PlaceholderFormat(sq.Dollar)

	sql, arg, err := query.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository_CCreateChatMember",
		QueryRaw: sql,
	}

	_, err = repo.db.DB().ExecContext(ctx, q, arg...)
	if err != nil {
		return err
	}

	return nil
}

// ReadChatMember - достает из базы данных участника (UserEmail) чата (chatID) при наличии
func (repo *repoChatMember) ReadChatMember(ctx context.Context, member *model.ChatMember) (*model.ChatMember, error) {
	query := sq.
		Select(membersChatID, membersUserEmail, membersJoinedAt).
		From(tableChatMember).
		Where(sq.Eq{membersChatID: member.ChatID, membersUserEmail: member.UserEmail}).PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "chat_repository_ReadChatMember",
		QueryRaw: sql,
	}

	chatmember := &modelrepo.ChatMember{}
	err = repo.db.DB().ScanOneContext(ctx, chatmember, q, args...)
	if err != nil {
		return nil, err
	}

	return convertor.ToChatMemberFromRepo(chatmember), nil
}

// ReadChatMembers - достает из базы данных список участников ([]*db.ChatMember) чата (chatID)
func (repo *repoChatMember) ReadChatMembers(ctx context.Context, chatID int) ([]*model.ChatMember, error) {
	query := sq.
		Select(membersChatID, membersUserEmail, membersJoinedAt).
		From(tableChatMember).
		Where(sq.Eq{membersChatID: chatID}).
		OrderBy("joined_at ASC").PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "chat_repository_ReadChatMembers",
		QueryRaw: sql,
	}

	var members []*modelrepo.ChatMember
	err = repo.db.DB().ScanAllContext(ctx, &members, q, args...)
	if err != nil {
		return nil, err
	}

	return convertor.ToChatMembersFromRepo(members), nil
}

// DeleteChatMember - удаляет участника чата по его (userEmail), в указаном чате (chatID)
func (repo *repoChatMember) DeleteChatMember(ctx context.Context, member *model.ChatMember) error {
	query := sq.Delete(tableChatMember).Where(sq.Eq{membersChatID: member.ChatID, membersUserEmail: member.UserEmail}).PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository_DeleteChatMember",
		QueryRaw: sql,
	}
	_, execErr := repo.db.DB().ExecContext(ctx, q, args...)
	if execErr != nil {
		return err
	}
	return nil
}

// checkMemberInChat - проверяет, состоит ли пользователь в чате
func checkMemberInChat(ctx context.Context, repo *repoChatMember, member *model.ChatMember) error {
	checkQuery := sq.
		Select("1").
		From(tableChatMember).
		Where(sq.Eq{membersChatID: member.ChatID, membersUserEmail: member.UserEmail}).PlaceholderFormat(sq.Dollar)

	sql, args, err := checkQuery.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository_checkMemberInChat",
		QueryRaw: sql,
	}

	var exists int
	err = repo.db.DB().ScanOneContext(ctx, &exists, q, args...)
	if err == nil {
		return err
	} else if err.Error() == "no rows in result set" {
		return nil
	}

	return err
}
