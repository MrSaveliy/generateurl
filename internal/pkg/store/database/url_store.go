package database

import (
	"github.com/Masterminds/squirrel"
)

func (s *Store) SetShort(short, url string) error {
	sq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	if _, err := sq.Insert("urls").SetMap(map[string]any{
		"short": short,
		"url":   url,
	}).RunWith(s.DB).Exec(); err != nil {
		return err
	}
	return nil
}

func (s *Store) GetUrl(short string) (string, error) {
	sq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	var url string
	if err := sq.Select("url").From("urls").Where(squirrel.Eq{"short": short}).RunWith(s.DB).QueryRow().Scan(&url); err != nil {
		return "", err
	}
	return url, nil
}
