package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	Db *sql.DB
}
