package provider

import (
	"database/sql"
	"errors"
)

func (p *Provider) SelectRandomHelloUser() (string, error) {
	var msg string

	err := p.conn.QueryRow("SELECT message FROM query ORDER BY RANDOM() LIMIT 1").Scan(&msg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return "hello, " + msg + "!", nil
}

func (p *Provider) InsertUser(msg string) error {
	_, err := p.conn.Exec("INSERT INTO query (message) VALUES ($1)", msg)
	if err != nil {
		return err
	}
	return nil
}
