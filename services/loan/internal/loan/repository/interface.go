package repository

import (
	"context"
	"database/sql"

	"github.com/kupanglie/simple-loan-apps/services/loan/internal/loan/entity"
)

type Loan interface {
	Add(ctx context.Context, tx *sql.Tx, request entity.Loan) (*sql.Tx, int, error)
	FindById(ctx context.Context, id int) (entity.Loan, error)
	FindByIdentityNumber(ctx context.Context, identityNumber string) ([]entity.Loan, error)
}
