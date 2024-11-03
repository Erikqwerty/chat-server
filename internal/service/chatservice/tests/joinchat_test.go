package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock"
	"github.com/stretchr/testify/require"

	"github.com/erikqwerty/chat-server/internal/model"
	"github.com/erikqwerty/chat-server/internal/repository"
	repoMock "github.com/erikqwerty/chat-server/internal/repository/mocks"
	"github.com/erikqwerty/chat-server/internal/service/chatservice"
	"github.com/erikqwerty/chat-server/pkg/db"
	dbMock "github.com/erikqwerty/chat-server/pkg/db/mocks"
)

func TestJoinChat(t *testing.T) {
	t.Parallel()

	type chatServerRepoMockFunc func(mc *minimock.Controller) repository.ChatServerRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx        context.Context
		chatMember *model.ChatMember
	}

	var (
		ctx        = context.Background()
		mc         = minimock.NewController(t)
		chatID     = gofakeit.Int()
		userEmail  = gofakeit.Email()
		actionType = "JOIN_CHAT"
		logDetails = "детальная информация отсутствует"

		chat = &model.Chat{
			ID:       chatID,
			ChatName: "Test Chat",
		}

		members = []*model.ChatMember{
			{
				ChatID:    chatID,
				UserEmail: gofakeit.Email(),
			},
			{
				ChatID:    chatID,
				UserEmail: gofakeit.Email(),
			},
		}

		messages = []*model.Message{
			{
				ID:        gofakeit.Int(),
				ChatID:    chatID,
				Text:      "Hello!",
				UserEmail: members[0].UserEmail,
			},
			{
				ID:        gofakeit.Int(),
				ChatID:    chatID,
				Text:      "Welcome!",
				UserEmail: members[1].UserEmail,
			},
		}

		repoErr = errors.New("ошибка репозитория")
		txErr   = errors.New("ошибка транзакции")
	)

	tests := []struct {
		name                   string
		args                   args
		want                   *model.JoinChat
		err                    error
		chatServerRepoMockFunc chatServerRepoMockFunc
		txManagerMockFunc      txManagerMockFunc
	}{
		{
			name: "успешное добавление участника в чат",
			args: args{
				ctx: ctx,
				chatMember: &model.ChatMember{
					ChatID:    chatID,
					UserEmail: userEmail,
				},
			},
			want: &model.JoinChat{
				Chat:     chat,
				Members:  membersEmails(members),
				Messages: messages,
			},
			err: nil,
			chatServerRepoMockFunc: func(mc *minimock.Controller) repository.ChatServerRepository {
				mock := repoMock.NewChatServerRepositoryMock(t)
				mock.CreateChatMemberMock.Expect(ctx, &model.ChatMember{ChatID: chatID, UserEmail: userEmail}).Return(nil)
				mock.ReadChatMock.Expect(ctx, chatID).Return(chat, nil)
				mock.ReadChatMembersMock.Expect(ctx, chatID).Return(members, nil)
				mock.ReadMessagesMock.Expect(ctx, chatID).Return(messages, nil)
				mock.CreateLogMock.Expect(ctx, &model.Log{
					ActionType:    actionType,
					ActionDetails: logDetails,
				}).Return(nil)
				return mock
			},
			txManagerMockFunc: func(mc *minimock.Controller) db.TxManager {
				mock := dbMock.NewTxManagerMock(t)
				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return handler(ctx)
				})
				return mock
			},
		},
		{
			name: "ошибка при передаче nil для chatMember",
			args: args{
				ctx:        ctx,
				chatMember: nil,
			},
			want: nil,
			err:  chatservice.ErrJoinChatReq(),
			chatServerRepoMockFunc: func(mc *minimock.Controller) repository.ChatServerRepository {
				return repoMock.NewChatServerRepositoryMock(t)
			},
			txManagerMockFunc: func(mc *minimock.Controller) db.TxManager {
				return dbMock.NewTxManagerMock(t)
			},
		},
		{
			name: "ошибка при создании участника чата",
			args: args{
				ctx: ctx,
				chatMember: &model.ChatMember{
					ChatID:    chatID,
					UserEmail: userEmail,
				},
			},
			want: nil,
			err:  repoErr,
			chatServerRepoMockFunc: func(mc *minimock.Controller) repository.ChatServerRepository {
				mock := repoMock.NewChatServerRepositoryMock(t)
				mock.CreateChatMemberMock.Expect(ctx, &model.ChatMember{ChatID: chatID, UserEmail: userEmail}).Return(repoErr)
				return mock
			},
			txManagerMockFunc: func(mc *minimock.Controller) db.TxManager {
				mock := dbMock.NewTxManagerMock(t)
				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return handler(ctx)
				})
				return mock
			},
		},
		{
			name: "service: ошибка транзакции",
			args: args{
				ctx: ctx,
				chatMember: &model.ChatMember{
					ChatID:    chatID,
					UserEmail: userEmail,
				},
			},
			want: nil,
			err:  txErr,
			chatServerRepoMockFunc: func(mc *minimock.Controller) repository.ChatServerRepository {
				mock := repoMock.NewChatServerRepositoryMock(t)
				return mock
			},
			txManagerMockFunc: func(mc *minimock.Controller) db.TxManager {
				mock := dbMock.NewTxManagerMock(t)
				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return txErr
				})
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatRepoMock := tt.chatServerRepoMockFunc(mc)
			txManagerMock := tt.txManagerMockFunc(mc)

			service := chatservice.NewService(chatRepoMock, txManagerMock)

			got, err := service.JoinChat(tt.args.ctx, tt.args.chatMember)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, got)
		})
	}

	t.Cleanup(mc.Finish)
}

func membersEmails(members []*model.ChatMember) []string {
	emails := make([]string, len(members))
	for i, member := range members {
		emails[i] = member.UserEmail
	}
	return emails
}
