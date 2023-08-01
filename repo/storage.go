package repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type Storage struct {
	DBPool *pgxpool.Pool
}

func New(pgConn string, logger *log.Logger) Storage {
	PGPool, err := pgxpool.New(context.Background(), pgConn)
	if err != nil {
		logger.Println(err)
		log.Fatal(err)
	}

	return Storage{DBPool: PGPool}
}

func (s Storage) NewUser(id int64) error {
	_, err := s.DBPool.Exec(context.Background(), "INSERT INTO users (id, state) VALUES ($1, $2);", id, "Main")
	if err != nil && err.Error() == "ERROR: duplicate key value violates unique constraint \"users_pkey\" (SQLSTATE 23505)" {
		return nil
	}

	return err
}

func (s Storage) GetUserState(id int64) (string, error) {
	var userState string
	row := s.DBPool.QueryRow(context.Background(), "SELECT state FROM users WHERE id = $1;", id)
	err := row.Scan(&userState)
	if err != nil {
		return "", err
	}

	return userState, nil
}

func (s Storage) SetUserState(id int64, state string) error {
	_, err := s.DBPool.Exec(context.Background(), "UPDATE users SET state = $1 WHERE id = $2;", state, id)
	return err
}
