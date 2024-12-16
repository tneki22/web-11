package provider

import (
	"database/sql"
	"errors"
)

func (p *Provider) SelectRandomHello() (string, error) {
	var msg string

	// Получаем одно сообщение из таблицы hello, отсортированной в случайном порядке
	err := p.conn.QueryRow("SELECT message FROM hello_11 ORDER BY RANDOM()").Scan(&msg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return msg, nil
}

func (p *Provider) CheckHelloExitByMsg(msg string) (bool, error) {
	// Получаем одно сообщение из таблицы hello
	err := p.conn.QueryRow("SELECT message FROM hello_11 WHERE message = $1 LIMIT 1", msg).Scan(&msg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (p *Provider) InsertHello(msg string) error {
	_, err := p.conn.Exec("INSERT INTO hello_11 (message) VALUES ($1)", msg)
	if err != nil {
		return err
	}

	return nil
}
