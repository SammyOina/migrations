package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mainflux/mainflux/pkg/errors"
	"github.com/mainflux/migrations/tools/configs"
)

var (
	errConnect = errors.New("failed to connect to postgresql server")
)

// Connect creates a connection to the PostgreSQL instance
func Connect(cfg configs.PostgresConfig) (*sqlx.DB, error) {
	url := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s sslcert=%s sslkey=%s sslrootcert=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DB, cfg.Password, cfg.SSLMode, cfg.SSLCert, cfg.SSLKey, cfg.SSLRootCert)

	db, err := sqlx.Open("pgx", url)
	if err != nil {
		return nil, errors.Wrap(errConnect, err)
	}

	return db, nil
}
