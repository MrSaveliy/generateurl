package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"log"
	"os"
	"sync"
)

var (
	store *Store
	once  = &sync.Once{}
)

type Store struct {
	*sql.DB
}

func GetStore() *Store {
	once.Do(func() {
		connectionString := fmt.Sprintf(
			"%s://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("DB"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)

		s, err := sql.Open("postgres", connectionString)
		if err != nil {
			log.Fatal(err)
		}

		store = &Store{s}
	})
	return store
}

func (s *Store) Get(short string) (string, error) {
	return s.GetUrl(short)
}

func (s *Store) Set(url, short string) error {
	return s.SetShort(short, url)
}

func (s *Store) GetLastID() (int64, error) {
	sq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	var id int64
	if err := sq.Select("id").From("urls").OrderBy("id desc").Limit(1).RunWith(s.DB).QueryRow().Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 1, nil
		}
		return 0, err
	}
	return id + 1, nil
}
