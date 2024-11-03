package tests

import (
	"context"
	"errors"
	"testing"
	"time"

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

func TestJoinChat(t *testing.T) {
	t.Parallel()

	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.JoinChatRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatID    = gofakeit.Int64()
		userEmail = gofakeit.Email()

		ChatName = gofakeit.BeerName()

		member1 = gofakeit.Email()
		member2 = gofakeit.Email()
		Members = []string{member1, member2, userEmail}

		text  = gofakeit.Name()
		timen = time.Now()

		messageAPI = &desc.Message{
			FromUserEmail: member1,
			Text:          text,
			Timestamp:     timestamppb.New(timen),
		}

		messageService = &model.Message{
			ID:        gofakeit.Int(),
			ChatID:    int(chatID),
			UserEmail: member1,
			Text:      text,
			Timestamp: timen,
		}

		req = &desc.JoinChatRequest{
			ChatId:    chatID,
			UserEmail: userEmail,
		}

		res = &desc.JoinChatResponse{
			ChatId:       chatID,
			ChatName:     ChatName,
			Participants: Members,
			Messages:     []*desc.Message{messageAPI},
		}

		joinMemberServiceReq = &model.ChatMember{
			ChatID:    int(chatID),
			UserEmail: userEmail,
		}

		joinMemberServiceRes = &model.JoinChat{
			Chat: &model.Chat{
				ID:        int(chatID),
				ChatName:  ChatName,
				CreatedAt: timen,
			},
			Members:  []string{member1, member2, userEmail},
			Messages: []*model.Message{messageService},
		}

		tempErr = errors.New("service joinChat error")
	)

	tests := []struct {
		name                string
		args                args
		want                *desc.JoinChatResponse
		err                 error
		chatServiceMockFunc chatServiceMockFunc
	}{
		{
			name: "api join chat success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatServiceMockFunc: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMock.NewChatServiceMock(mc)
				mock.JoinChatMock.Expect(ctx, joinMemberServiceReq).Return(joinMemberServiceRes, nil)
				return mock
			},
		},
		{
			name: "api join chat error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  tempErr,
			chatServiceMockFunc: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMock.NewChatServiceMock(mc)
				mock.JoinChatMock.Expect(ctx, joinMemberServiceReq).Return(nil, tempErr)
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

			joinChat, err := api.JoinChat(tt.args.ctx, tt.args.req)

			if tt.err != nil {
				require.Error(t, err)
				require.EqualError(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tt.want, joinChat)
		})
	}
}
