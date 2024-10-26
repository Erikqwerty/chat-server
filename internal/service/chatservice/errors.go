package chatservice

import "errors"

func ErrCreateChatReq() error {
	return errors.New("запрос на создание чата отсутствует")
}

func ErrJoinChatReq() error {
	return errors.New("запрос на присоединение к чату отсутствует")
}

func ErrSendMessage() error {
	return errors.New("запрос на отправку сообщения отсутсвует")
}
