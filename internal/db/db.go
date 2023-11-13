package db

import (
	"database/sql"
	_ "embed"
	"log/slog"
	"time"

	_ "github.com/lib/pq"
)

//go:embed sql/migrate.sql
var migrateSql string

//go:embed sql/update_time.sql
var updateTimeSql string

//go:embed sql/delete_inactive_objects.sql
var deleteObjSql string

type Repository struct {
	db *sql.DB
}

func New() *Repository {
	connStr := "postgres://postgres:1234@localhost/test?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return &Repository{
		db: db,
	}
}

func (repo *Repository) Migrate() {
	_, _ = repo.db.Exec(migrateSql)
}

func (repo *Repository) UpdateUserLastSeen(id int) {
	time := time.Now()
	_, err := repo.db.Exec(updateTimeSql, id, time)
	if err != nil {
		slog.Error("Cant insert", "err", err)
	}
}

func (repo *Repository) DeleteObjectsThatOffline() {
	currentTime := time.Now()
	res, err := repo.db.Exec(deleteObjSql, currentTime.Add(-30*time.Second))
	if err != nil {
		slog.Error("cant delete users where last seen > 30s", "err", err)
	}
	slog.Info("db delete", "res", res)
}
