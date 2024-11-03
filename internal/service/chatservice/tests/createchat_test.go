package tests

import (
	"testing"

	"github.com/gojuno/minimock"

	"github.com/erikqwerty/chat-server/internal/repository"
	"github.com/erikqwerty/chat-server/pkg/db"
)

func TestCreateChat(t *testing.T) {
	t.Parallel()

	type authRepoMockFunc func(mc *minimock.Controller) repository.Chat
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

}
