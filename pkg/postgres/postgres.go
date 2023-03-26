package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	_defaultPoolSize    = 1
	_defaultConnAttemps = 10
	_defatulTimeOut     = time.Second
)

type Postgres struct {
	maxPoolSize int
	connAttemps int
	connTimeout time.Duration
	Builder     squirrel.StatementBuilderType
	Pool        *pgxpool.Pool
}

func New(url string, options ...Option) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize: _defaultPoolSize,
		connAttemps: _defaultConnAttemps,
		connTimeout: _defatulTimeOut,
	}
	for _, opt := range options {
		opt(pg)
	}

	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	poolConfig, err := pgxpool.ParseConfig(url)

	if err != nil {
		return nil, fmt.Errorf("pxpool parse config error %w", err)
	}

	poolConfig.MaxConns = int32(pg.maxPoolSize)

	for pg.connAttemps > 0 {
		pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err == nil {
			break
		}

		time.Sleep(pg.connTimeout)

		pg.connAttemps--
	}
	if err != nil {
		return nil, fmt.Errorf("connection failed %w", err)
	}

	return pg, nil

}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
