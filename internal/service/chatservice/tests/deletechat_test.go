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

func TestDeleteChat(t *testing.T) {
	t.Parallel()

	type chatServerRepoMockFunc func(mc *minimock.Controller) repository.ChatServerRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		id  int64
	}

	var (
		ctx        = context.Background()
		mc         = minimock.NewController(t)
		chatID     = gofakeit.Int64()
		actionType = "DELETE_CHAT"
		logDetails = "детальная информация отсутствует"

		repoErr = errors.New("repository error")
		txErr   = errors.New("transaction error")
	)

	tests := []struct {
		name                   string
		args                   args
		err                    error
		chatServerRepoMockFunc chatServerRepoMockFunc
		txManagerMockFunc      txManagerMockFunc
	}{
		{
			name: "service: успешное удаление чата",
			args: args{
				ctx: ctx,
				id:  chatID,
			},
			err: nil,
			chatServerRepoMockFunc: func(_ *minimock.Controller) repository.ChatServerRepository {
				mock := repoMock.NewChatServerRepositoryMock(t)
				mock.DeleteChatMock.Expect(ctx, int(chatID)).Return(nil)
				mock.CreateLogMock.Expect(ctx, &model.Log{
					ActionType:    actionType,
					ActionDetails: logDetails,
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
			name: "service: ошибка при удалении чата",
			args: args{
				ctx: ctx,
				id:  chatID,
			},
			err: repoErr,
			chatServerRepoMockFunc: func(_ *minimock.Controller) repository.ChatServerRepository {
				mock := repoMock.NewChatServerRepositoryMock(t)
				mock.DeleteChatMock.Expect(ctx, int(chatID)).Return(repoErr)
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
			name: "service: ошибка при создании лога",
			args: args{
				ctx: ctx,
				id:  chatID,
			},
			err: repoErr,
			chatServerRepoMockFunc: func(_ *minimock.Controller) repository.ChatServerRepository {
				mock := repoMock.NewChatServerRepositoryMock(t)
				mock.DeleteChatMock.Expect(ctx, int(chatID)).Return(nil)
				mock.CreateLogMock.Expect(ctx, &model.Log{
					ActionType:    actionType,
					ActionDetails: logDetails,
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
		{
			name: "service: ошибка транзакции",
			args: args{
				ctx: ctx,
				id:  chatID,
			},
			err: txErr,
			chatServerRepoMockFunc: func(_ *minimock.Controller) repository.ChatServerRepository {
				mock := repoMock.NewChatServerRepositoryMock(t)
				return mock
			},
			txManagerMockFunc: func(_ *minimock.Controller) db.TxManager {
				mock := dbMock.NewTxManagerMock(t)
				mock.ReadCommittedMock.Set(func(_ context.Context, _ db.Handler) error {
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

			err := service.DeleteChat(tt.args.ctx, tt.args.id)

			require.Equal(t, tt.err, err)
		})
	}

	t.Cleanup(mc.Finish)
}
