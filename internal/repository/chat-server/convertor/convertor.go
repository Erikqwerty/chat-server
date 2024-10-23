package convertor

import (
	"github.com/erikqwerty/chat-server/internal/model"
	modelrepo "github.com/erikqwerty/chat-server/internal/repository/chat-server/model"
)

// ToChatFromRepo - конвертор преобразующий структуру chat repo слоя в структуру бизнес-логики
func ToChatFromRepo(modeldb *modelrepo.Chat) *model.Chat {
	return &model.Chat{
		ID:        modeldb.ID,
		ChatName:  modeldb.ChatName,
		CreatedAt: modeldb.CreatedAt,
	}
}

// ToChatFromRepo - конвертор преобразующий набор структур chat repo слоя в структуру бизнес-логики
func ToChatsFromRepo(modeldb []*modelrepo.Chat) []*model.Chat {

	chats := make([]*model.Chat, len(modeldb))

	for i, chat := range modeldb {
		chats[i] = &model.Chat{
			ID:        chat.ID,
			ChatName:  chat.ChatName,
			CreatedAt: chat.CreatedAt,
		}
	}

	return chats
}

// ToChatMemberFromRepo - конвертор преобразующий структуру ChatMember repo слоя в структуру бизнес-логики
func ToChatMemberFromRepo(modeldb *modelrepo.ChatMember) *model.ChatMember {
	return &model.ChatMember{
		ChatID:    modeldb.ChatID,
		UserEmail: modeldb.UserEmail,
		JoinedAt:  modeldb.JoinedAt,
	}
}

// ToChatMemberFromRepo - конвертор преобразующий набор структур ChatMember  repo слоя в структуру бизнес-логики
func ToChatMembersFromRepo(modeldb []*modelrepo.ChatMember) []*model.ChatMember {
	chatMember := make([]*model.ChatMember, len(modeldb))

	for i, chat := range modeldb {
		chatMember[i] = &model.ChatMember{
			ChatID:    chat.ChatID,
			UserEmail: chat.UserEmail,
			JoinedAt:  chat.JoinedAt,
		}
	}

	return chatMember
}

// ToMessagesFromRepo - конвертор преобразующий набор структур Messages  repo слоя в структуру бизнес-логики
func ToMessagesFromRepo(modeldb []*modelrepo.Message) []*model.Message {
	mesages := make([]*model.Message, len(modeldb))

	for i, chat := range modeldb {
		mesages[i] = &model.Message{
			ID:        chat.ID,
			ChatID:    chat.ChatID,
			UserEmail: chat.UserEmail,
			Text:      chat.Text,
			Timestamp: chat.Timestamp,
		}
	}

	return mesages
}
