package pg

import "fmt"

// errSQLCreateQwery - ошибка преобразования запроса
func errSQLCreateQwery(err error) error {
	return fmt.Errorf("ошибка построения SQL-запроса: %w", err)
}

// errSQLQwery - ошибка выполнения sql запроса
func errSQLQwery(err error) error {
	return fmt.Errorf("ошибка выполнения sql запроса: %w", err)
}

// errDSN - Ошибка парсинга dsn для подключения к БД
func errDSN(err error) error {
	return fmt.Errorf("ошибка парсинга dsn: %w", err)
}

// errDBConect - Ошибка подключения к базе данных
func errDBConect(err error) error {
	return fmt.Errorf("ошибка подключения к database: %w", err)
}

// errUserAlreadyExists - Ошибка  пользователь уже существует
func errUserAlreadyExists(err error) error {
	return fmt.Errorf("пользователь уже существует %w", err)
}
