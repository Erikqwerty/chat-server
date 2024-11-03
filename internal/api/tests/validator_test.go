package tests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/erikqwerty/chat-server/internal/api"
	"github.com/erikqwerty/chat-server/pkg/utils/validator"
)

func TestValidateRequest(t *testing.T) {
	tests := []struct {
		name        string
		req         interface{}
		expectedErr error
	}{
		{
			name: "valid request",
			req: &struct {
				FromUserEmail string
				UserEmail     string
				Emails        []string
				ChatName      string
				ChatID        int64
				Text          string
			}{
				FromUserEmail: "sender@example.com",
				UserEmail:     "user@example.com",
				Emails:        []string{"user1@example.com", "user2@example.com"},
				ChatName:      "ChatRoom",
				ChatID:        1,
				Text:          "Hello!",
			},
			expectedErr: nil,
		},
		{
			name: "missing FromUserEmail",
			req: &struct {
				FromUserEmail string
				UserEmail     string
			}{
				FromUserEmail: "",
				UserEmail:     "user@example.com",
			},
			expectedErr: api.ErrFromUserEmail,
		},
		{
			name: "invalid FromUserEmail format",
			req: &struct {
				FromUserEmail string
			}{
				FromUserEmail: "invalid-email",
			},
			expectedErr: api.ErrInvalidEmail,
		},
		{
			name: "missing UserEmail",
			req: &struct {
				UserEmail string
			}{
				UserEmail: "",
			},
			expectedErr: api.ErrUserEmailJoinChat,
		},
		{
			name: "invalid UserEmail format",
			req: &struct {
				UserEmail string
			}{
				UserEmail: "invalid-email",
			},
			expectedErr: api.ErrInvalidEmail,
		},
		{
			name: "empty Emails slice",
			req: &struct {
				Emails []string
			}{
				Emails: []string{},
			},
			expectedErr: api.ErrChatMembersNotSpecifed,
		},
		{
			name: "invalid Emails element format",
			req: &struct {
				Emails []string
			}{
				Emails: []string{"user1@example.com", "invalid-email"},
			},
			expectedErr: validator.ValidEmails([]string{"user1@example.com", "invalid-email"}),
		},
		{
			name: "missing ChatName",
			req: &struct {
				ChatName string
			}{
				ChatName: "",
			},
			expectedErr: api.ErrChatNameNotSpecified,
		},
		{
			name: "missing ChatID",
			req: &struct {
				ChatID int64
			}{
				ChatID: 0,
			},
			expectedErr: api.ErrChatIDNotSpecifed,
		},
		{
			name: "missing Text",
			req: &struct {
				Text string
			}{
				Text: "",
			},
			expectedErr: api.ErrMessageTextNotSpecifed,
		},
		{
			name:        "invalid request type (non-struct)",
			req:         "invalid request",
			expectedErr: errors.New("ожидалась структура для валидации"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := api.ValidateRequest(tt.req)
			if tt.expectedErr != nil {
				require.Error(t, err)
				require.Equal(t, tt.expectedErr.Error(), err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
