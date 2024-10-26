package api

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/erikqwerty/chat-server/pkg/utils/validator"
)

// валидируемые данные поля структур запроса gRPC контракта
const (
	fromUserEmail = "FromUserEmail"
	userEmail     = "UserEmail"
	emails        = "Emails"
	chatName      = "ChatName"
	chatID        = "ChatId"
	messageText   = "Text"
	// timeStamp     = "Timestamp" ??? - нужно ли
)

// Ошибки для валидации
var (
	errUserEmailJoinChat      = errors.New("не передан email пользователя который хочет присоеденится к чату")
	errFromUserEmail          = errors.New("не передан email отправителя сообщения")
	errChatNameNotSpecified   = errors.New("не указанно название чата")
	errChatMembersNotSpecifed = errors.New("не переданы участники чата")
	errChatIDNotSpecifed      = errors.New("не указан id чата для удаления")
	errMessageTextNotSpecifed = errors.New("нельзя отправлять пустое сообщение")

	errInvalidEmail = errors.New("email не валиден")
)

func validateRequest(req interface{}) error {
	v := reflect.ValueOf(req)

	// Проверка на указатель и получение значения
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Проверка, что входные данные являются структурой
	if v.Kind() != reflect.Struct {
		return errors.New("ожидалась структура для валидации")
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := v.Type().Field(i).Name
		switch fieldName {
		case fromUserEmail:
			if field.String() == "" {
				return errFromUserEmail
			}
			if !validator.IsValidEmail(field.String()) {
				return errInvalidEmail
			}
		case userEmail:
			if field.String() == "" {
				return errUserEmailJoinChat
			}
			if !validator.IsValidEmail(field.String()) {
				return errInvalidEmail
			}
		case emails:
			if field.Kind() != reflect.Slice || field.Type().Elem().Kind() != reflect.String {
				return fmt.Errorf("поле Emails должно быть слайсом строк")
			}

			emails := field.Interface().([]string)

			if len(emails) == 0 {
				return errChatMembersNotSpecifed
			}

			if err := validator.ValidEmails(emails); err != nil {
				return err
			}
		case chatName:
			if field.String() == "" {
				return errChatNameNotSpecified
			}
		case chatID:
			if field.Int() == 0 {
				return errChatIDNotSpecifed
			}
		case messageText:
			if field.String() == "" {
				return errMessageTextNotSpecifed
			}
		}
	}
	return nil
}
