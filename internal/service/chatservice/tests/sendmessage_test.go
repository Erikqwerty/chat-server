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

func TestSendMessage(t *testing.T) {
	t.Parallel()

	type chatRepoMockFunc func(mc *minimock.Controller) repository.ChatServerRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	var (
		ctx        = context.Background()
		mc         = minimock.NewController(t)
		actionType = "SEND_MESSAGE"
		logDetails = "детальная информация отсутствует"

		message = &model.Message{
			ID:        gofakeit.Int(),
			ChatID:    gofakeit.Int(),
			Text:      gofakeit.Sentence(5),
			UserEmail: gofakeit.Email(),
		}

		repoErr = errors.New("ошибка репозитория")
		txErr   = errors.New("ошибка транзакции")
	)

	tests := []struct {
		name              string
		msg               *model.Message
		err               error
		chatRepoMockFunc  chatRepoMockFunc
		txManagerMockFunc txManagerMockFunc
	}{
		{
			name: "service: ошибка: сообщение nil",
			msg:  nil,
			err:  chatservice.ErrSendMessage(),
			chatRepoMockFunc: func(_ *minimock.Controller) repository.ChatServerRepository {
				return repoMock.NewChatServerRepositoryMock(t)
			},
			txManagerMockFunc: func(_ *minimock.Controller) db.TxManager {
				return dbMock.NewTxManagerMock(t)
			},
		},
		{
			name: "service: успешная отправка сообщения",
			msg:  message,
			err:  nil,
			chatRepoMockFunc: func(_ *minimock.Controller) repository.ChatServerRepository {
				mock := repoMock.NewChatServerRepositoryMock(t)

				mock.CreateMessageMock.Expect(ctx, message).Return(message.ID, nil)
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
			name: "service: ошибка при создании сообщения",
			msg:  message,
			err:  repoErr,
			chatRepoMockFunc: func(_ *minimock.Controller) repository.ChatServerRepository {
				mock := repoMock.NewChatServerRepositoryMock(t)

				mock.CreateMessageMock.Expect(ctx, message).Return(0, repoErr)

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
			msg:  message,
			err:  repoErr,
			chatRepoMockFunc: func(_ *minimock.Controller) repository.ChatServerRepository {
				mock := repoMock.NewChatServerRepositoryMock(t)

				mock.CreateMessageMock.Expect(ctx, message).Return(message.ID, nil)
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
			msg:  message,
			err:  txErr,
			chatRepoMockFunc: func(_ *minimock.Controller) repository.ChatServerRepository {
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

			chatRepoMock := tt.chatRepoMockFunc(mc)
			txManagerMock := tt.txManagerMockFunc(mc)

			service := chatservice.NewService(chatRepoMock, txManagerMock)

			err := service.SendMessage(ctx, tt.msg)

			require.Equal(t, tt.err, err)
		})
	}

	t.Cleanup(mc.Finish)
}
