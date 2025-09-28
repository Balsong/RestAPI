package postgres

import (
	"API/internal/config"
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	goose "github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
)

const driverName = "pgx"

func NewConnection(
	ctx context.Context,
	cfg config.Postgres,
	log *zerolog.Logger,
	dsn ...string,
) (*sqlx.DB, error) {
	var confDsn pgx.ConnConfig
	var err error

	if len(dsn) == 1 {
		confDsn, err = pgx.ParseDSN(dsn[0])
		if err != nil {
			return nil, fmt.Errorf("parse dsn: %w", err)
		}
	}

	addr := strings.Split(cfg.Host, ":")
	port, err := strconv.ParseUint(addr[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parse port: %w", err)
	}

	confPgx := pgx.ConnConfig{
		Host:                 addr[0],
		Port:                 uint16(port),
		User:                 cfg.User,
		Password:             cfg.Password,
		Database:             cfg.DB,
		PreferSimpleProtocol: cfg.SimpleProtocol,
	}

	conf := confDsn.Merge(confPgx)

	db := sqlx.NewDb(stdlib.OpenDB(conf), driverName)

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	if err = applyMigrations(db, log); err != nil {
		return nil, err
	}

	return db, nil
}

const lockKey = 111

func applyMigrations(db *sqlx.DB, log *zerolog.Logger) (err error) {
	err = goose.SetDialect("postgres")
	if err != nil {
		return err
	}

	if _, err := db.Exec("SELECT pg_advisory_lock($1)", lockKey); err != nil {
		return err
	}

	defer func() {
		if _, err := db.Exec("SELECT pg_advisory_unlock($1)", lockKey); err != nil {
			log.Error().Err(err).Msg("got err unlocking DB")
		}
	}()

	err = goose.RunWithOptionsContext(
		context.Background(),
		"up",
		db.DB,
		"./migrations/postgres",
		nil,
		goose.WithAllowMissing(),
	)
	if err != nil {
		return err
	}

	return nil
}
