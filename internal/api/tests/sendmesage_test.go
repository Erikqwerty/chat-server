package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/erikqwerty/chat-server/internal/api"
	"github.com/erikqwerty/chat-server/internal/model"
	"github.com/erikqwerty/chat-server/internal/service"
	serviceMock "github.com/erikqwerty/chat-server/internal/service/mocks"
	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()

	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.SendMessageRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		userEmail = gofakeit.Email()
		textMess  = gofakeit.Language()
		timet     = timestamppb.Now()
		chatID    = gofakeit.Int64()

		req = &desc.SendMessageRequest{
			FromUserEmail: userEmail,
			Text:          textMess,
			Timestamp:     timet,
			ChatId:        chatID,
		}

		message = &model.Message{
			ChatID:    int(chatID),
			UserEmail: userEmail,
			Text:      textMess,
			Timestamp: timet.AsTime(),
		}

		tempErr = errors.New("service send message error")
	)

	tests := []struct {
		name                string
		args                args
		want                *desc.SendMessageRequest
		err                 error
		chatServiceMockFunc chatServiceMockFunc
	}{
		{
			name: "api send message success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: nil,
			chatServiceMockFunc: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMock.NewChatServiceMock(mc)

				mock.SendMessageMock.Expect(ctx, message).Return(nil)

				return mock
			},
		},
		{
			name: "api send message error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: tempErr,
			chatServiceMockFunc: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMock.NewChatServiceMock(mc)

				mock.SendMessageMock.Expect(ctx, message).Return(tempErr)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatServiceMock := tt.chatServiceMockFunc(mc)
			api := api.NewChatServerGRPCImplementation(chatServiceMock)

			_, err := api.SendMessage(tt.args.ctx, tt.args.req)

			if tt.err != nil {
				require.Error(t, err)
				require.EqualError(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
