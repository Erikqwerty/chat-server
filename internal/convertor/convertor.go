package convertor

import (
	"github.com/erikqwerty/chat-server/internal/model"
	desc "github.com/erikqwerty/chat-server/pkg/chatapi_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToModelCreateChatFromCreateReq - конвертор из api слоя в сервисыный создание чата
func ToModelCreateChatFromCreateReq(req *desc.CreateRequest) *model.CreateChat {
	return &model.CreateChat{
		ChatName:     req.ChatName,
		MembersEmail: req.Emails,
	}
}

// ToModelChatMemberFromJoinChatRequest desc.JoinChatRequest --> model.ChatMember
func ToModelChatMemberFromJoinChatRequest(req *desc.JoinChatRequest) *model.ChatMember {
	return &model.ChatMember{
		ChatID:    int(req.ChatId),
		UserEmail: req.UserEmail,
	}
}

// ToChatAPIJoinRespFromModelJoinChat - преобразует model.JoinChat --> desc.JoinChatResponse
func ToChatAPIJoinRespFromModelJoinChat(joinChat *model.JoinChat) *desc.JoinChatResponse {
	messages := make([]*desc.Message, len(joinChat.Messages))
	for i, mess := range joinChat.Messages {
		messages[i] = &desc.Message{
			From:      mess.UserEmail,
			Text:      mess.Text,
			Timestamp: timestamppb.New(mess.Timestamp),
		}
	}
	return &desc.JoinChatResponse{
		ChatId:       int64(joinChat.ID),
		ChatName:     joinChat.ChatName,
		Participants: joinChat.Members,
		Messages:     messages,
	}
}

// ToModelMessageFromReqSendMessage desc.SendMessageRequest --> model.Message
func ToModelMessageFromReqSendMessage(req *desc.SendMessageRequest) *model.Message {
	return &model.Message{
		ChatID:    int(req.ChatId),
		UserEmail: req.From,
		Text:      req.Text,
	}
}