package repo

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/amlcx/tablero/backend/sentinel"
	"github.com/charmbracelet/log"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/uptrace/bun"
)

type BaseRepo interface {
	getDB() *bun.DB
	log() *log.Logger
	parse(error) error
}

type baseRepo struct {
	db     *bun.DB
	logger *log.Logger
}

var _ BaseRepo = (*baseRepo)(nil)

func NewBaseRepo(
	db *bun.DB,
	logger *log.Logger,
) BaseRepo {
	sentinel.Assert(db != nil, "failed to initialize base repo: nil db")
	sentinel.Assert(logger != nil, "failed to initialize base repo: nil logger")

	return &baseRepo{
		db:     db,
		logger: logger,
	}
}

func (r *baseRepo) log() *log.Logger {
	return r.logger
}

func (r *baseRepo) getDB() *bun.DB {
	return r.db
}

func (r *baseRepo) parse(err error) error {
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError

	if !errors.As(err, &pgErr) {
		if errors.Is(err, sql.ErrNoRows) {
			return &DBErr{
				Code:    NoResults,
				Message: err.Error(),
				Wrapped: err,
			}
		}

		return &DBErr{
			Code:    Unknown,
			Message: err.Error(),
			Wrapped: err,
		}
	}

	dbErr := &DBErr{
		Code:    pgCodeToErrorCode(pgErr.Code),
		Message: err.Error(),
		Wrapped: err,
	}

	return dbErr
}

const (
	pgTooLong = "22001"
	pgNotNull = "23502"
	pgFK      = "23503"
	pgUnique  = "23505"
)

type ErrorCode string

const (
	TooLong        = "TOO_LONG_CONFLICT"
	NotNull        = "NOT_NULL_CONFLICT"
	ForeignKey     = "FOREIGN_KEY_CONFLICT"
	UniqueConflict = "UNIQUE_CONFLICT"
	NoResults      = "NO_RESULTS"
	Unknown        = "UNKNOWN"
)

type DBErr struct {
	Code    ErrorCode
	Message string
	Wrapped error
}

func (e *DBErr) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *DBErr) Unwrap() error {
	return e.Wrapped
}

func pgCodeToErrorCode(pgCode string) ErrorCode {
	switch pgCode {
	case pgTooLong:
		return TooLong
	case pgNotNull:
		return NotNull
	case pgFK:
		return ForeignKey
	case pgUnique:
		return UniqueConflict
	default:
		return Unknown
	}
}
