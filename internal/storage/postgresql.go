package storage

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Storage struct {
	config *Config
	db     *sql.DB
	mrepo  *Repo
}

func New(cfg *Config) *Storage {
	return &Storage{
		config: cfg,
	}
}
func (s *Storage) Open() error {
	var err error
	s.db, err = sql.Open("postgres", s.database_url())
	if err != nil {
		return fmt.Errorf("error to opening database: %w", err)
	}
	if err = s.ping(); err != nil {
		return fmt.Errorf("error connecting to database : %w", err)
	}
	return nil
}
func (s *Storage) MigUp() error {
	m, err := migrate.New("file://"+s.config.MigPath, s.database_url())
	if err != nil {
		return fmt.Errorf("error with file migration: %w", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error with 'Up' migration: %w", err)
	}
	return nil
}
func (s *Storage) ping() error {
	var err error
	for i := 0; i < 3; i++ {
		err = s.db.Ping()
		if err != nil {
			time.Sleep(time.Second * 2)
		} else {
			return nil
		}
	}
	return err
}
func (s *Storage) database_url() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		s.config.User, s.config.Password, s.config.Host, s.config.Port, s.config.DbName)
}

func (s *Storage) Interact() *Repo {
	if s.mrepo != nil {
		return s.mrepo
	}
	s.mrepo = &Repo{
		storage: s.db,
	}
	return s.mrepo
}
