package chatservice

import "errors"

// ErrCreateChatReq - запрос на создание чата отсутствует
func ErrCreateChatReq() error {
	return errors.New("запрос на создание чата отсутствует")
}

// ErrJoinChatReq - запрос на присоединение к чату отсутствует
func ErrJoinChatReq() error {
	return errors.New("запрос на присоединение к чату отсутствует")
}

// ErrSendMessage - запрос на отправку сообщения отсутсвует
func ErrSendMessage() error {
	return errors.New("запрос на отправку сообщения отсутсвует")
}
