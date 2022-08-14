package repository

import (
	"context"
	"database/sql"

	"github.com/kupanglie/simple-loan-apps/services/loan/internal/user/entity"
)

type User interface {
	Add(ctx context.Context, tx *sql.Tx, request entity.User) (*sql.Tx, int64, error)
	FindById(ctx context.Context, id int) (entity.User, error)
	FindByIdentityNumber(ctx context.Context, identityNumber string) (entity.User, error)
}
