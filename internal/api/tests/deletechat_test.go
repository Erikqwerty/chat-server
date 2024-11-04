package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/erikqwerty/chat-server/internal/api"
	"github.com/erikqwerty/chat-server/internal/service"
	serviceMock "github.com/erikqwerty/chat-server/internal/service/mocks"
	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

func TestDeleteChat(t *testing.T) {
	t.Parallel()

	type chatSericeMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.DeleteRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		req = &desc.DeleteRequest{
			ChatId: id,
		}

		tempErr = errors.New("service delete error")
	)

	tests := []struct {
		name               string
		args               args
		err                error
		chatSericeMockFunc chatSericeMockFunc
	}{
		{
			name: "api delete chat success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: nil,
			chatSericeMockFunc: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMock.NewChatServiceMock(mc)

				mock.DeleteChatMock.Expect(ctx, req.ChatId).Return(nil)

				return mock
			},
		},
		{
			name: "api delete chat error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: tempErr,
			chatSericeMockFunc: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMock.NewChatServiceMock(mc)

				mock.DeleteChatMock.Expect(ctx, req.ChatId).Return(tempErr)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatServiceMock := tt.chatSericeMockFunc(mc)
			api := api.NewChatServerGRPCImplementation(chatServiceMock)

			_, err := api.DeleteChat(tt.args.ctx, tt.args.req)

			if tt.err != nil {
				require.Error(t, err)
				require.EqualError(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
