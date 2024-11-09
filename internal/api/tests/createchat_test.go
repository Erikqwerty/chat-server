package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/erikqwerty/chat-server/internal/api"
	"github.com/erikqwerty/chat-server/internal/model"
	"github.com/erikqwerty/chat-server/internal/service"
	serviceMock "github.com/erikqwerty/chat-server/internal/service/mocks"
	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
)

func TestCreateChat(t *testing.T) {
	t.Parallel()

	type chatSericeMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		emails   = []string{gofakeit.Email(), gofakeit.Email()}
		chatName = gofakeit.Name()

		req = &desc.CreateRequest{
			Emails:   emails,
			ChatName: chatName,
		}

		createChat = &model.CreateChat{
			ChatName:     chatName,
			MembersEmail: emails,
		}

		id = gofakeit.Int64()

		res = &desc.CreateResponse{
			Id: id,
		}

		tempErr = errors.New("service create error")
	)

	tests := []struct {
		name               string
		args               args
		want               *desc.CreateResponse
		err                error
		chatSericeMockFunc chatSericeMockFunc
	}{
		{
			name: "api create chat success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatSericeMockFunc: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMock.NewChatServiceMock(mc)

				mock.CreateChatMock.Expect(ctx, createChat).Return(id, nil)

				return mock
			},
		},
		{
			name: "api create chat error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  tempErr,
			chatSericeMockFunc: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMock.NewChatServiceMock(mc)

				mock.CreateChatMock.Expect(ctx, createChat).Return(0, tempErr)

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

			newID, err := api.CreateChat(tt.args.ctx, tt.args.req)

			if tt.err != nil {
				require.Error(t, err)
				require.EqualError(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tt.want, newID)
		})
	}
}
