package provider

import (
	"database/sql"
	"errors"
	"strconv"
)

func (p *Provider) SelectCount() (string, error) {
	var count string

	err := p.conn.QueryRow("SELECT count FROM count").Scan(&count)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}
	return count, nil
}

func (p *Provider) InsertCount(count int) error {
	var oldcount string
	err := p.conn.QueryRow("SELECT count FROM count").Scan(&oldcount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return err
	}
	oldnumber, err := strconv.Atoi(oldcount)
	if err != nil {
		return err
	}

	_, err = p.conn.Exec("UPDATE count SET count = ($1)", oldnumber+count)
	if err != nil {
		return err
	}
	return nil
}
