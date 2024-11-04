package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/erikqwerty/erik-platform/clients/db"
	dbMock "github.com/erikqwerty/erik-platform/clients/db/mocks"
	"github.com/gojuno/minimock"
	"github.com/stretchr/testify/require"

	"github.com/erikqwerty/chat-server/internal/model"
	"github.com/erikqwerty/chat-server/internal/repository"
	repoMock "github.com/erikqwerty/chat-server/internal/repository/mocks"
	"github.com/erikqwerty/chat-server/internal/service/chatservice"
)

func TestCreateChat(t *testing.T) {
	t.Parallel()

	type chatServerRepoMockFunc func(mc *minimock.Controller) repository.ChatServerRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req *model.CreateChat
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatID       = gofakeit.Int64()
		chatName     = gofakeit.Name()
		membersEmail = []string{gofakeit.Email()}

		req = &model.CreateChat{
			ChatName:     chatName,
			MembersEmail: membersEmail,
		}

		repoErr = errors.New("репо ошибка")
	)

	tests := []struct {
		name                   string
		args                   args
		want                   int64
		err                    error
		chatServerRepoMockFunc chatServerRepoMockFunc
		txManagerMockFunc      txManagerMockFunc
	}{
		{
			name: "service: успешное создание чата",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: chatID,
			err:  nil,
			chatServerRepoMockFunc: func(_ *minimock.Controller) repository.ChatServerRepository {
				mock := repoMock.NewChatServerRepositoryMock(t)
				mock.CreateChatMock.Expect(ctx, req.ChatName).Return(int(chatID), nil)
				mock.CreateChatMemberMock.Expect(ctx, &model.ChatMember{
					ChatID:    int(chatID),
					UserEmail: membersEmail[0],
				}).Return(nil)
				mock.CreateLogMock.Expect(ctx, &model.Log{
					ActionType:    "CREATE_CHAT",
					ActionDetails: "детальная информация отсутствует",
				}).Return(nil)
				return mock
			},
			txManagerMockFunc: func(_ *minimock.Controller) db.TxManager {
				mock := dbMock.NewTxManagerMock(t)
				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return handler(ctx)
				})
				return mock
			},
		},
		{
			name: "service: ошибка транзакции при создание чата",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  errors.New("transaction failed"),
			chatServerRepoMockFunc: func(_ *minimock.Controller) repository.ChatServerRepository {
				mock := repoMock.NewChatServerRepositoryMock(t)
				return mock
			},
			txManagerMockFunc: func(_ *minimock.Controller) db.TxManager {
				mock := dbMock.NewTxManagerMock(t)
				mock.ReadCommittedMock.Set(func(_ context.Context, _ db.Handler) error {
					return errors.New("transaction failed")
				})
				return mock
			},
		},
		{
			name: "service: ошибка добавления участников чата",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  repoErr,
			chatServerRepoMockFunc: func(_ *minimock.Controller) repository.ChatServerRepository {
				mock := repoMock.NewChatServerRepositoryMock(t)
				mock.CreateChatMock.Expect(ctx, req.ChatName).Return(int(chatID), nil)

				mock.CreateChatMemberMock.Expect(ctx, &model.ChatMember{
					ChatID:    int(chatID),
					UserEmail: membersEmail[0],
				}).Return(repoErr)
				return mock
			},
			txManagerMockFunc: func(_ *minimock.Controller) db.TxManager {
				mock := dbMock.NewTxManagerMock(t)
				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return handler(ctx)
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

			servic := chatservice.NewService(chatRepoMock, txManagerMock)

			ID, err := servic.CreateChat(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, ID)
		})
	}

	t.Cleanup(mc.Finish)
}
