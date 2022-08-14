package usecase

import (
	"context"

	"github.com/kupanglie/simple-loan-apps/services/loan/internal/loan/entity"
)

type User interface {
	Add(ctx context.Context, request entity.CreateLoanRequest)
	FindById(ctx context.Context, id int) (entity.Loan, error)
	FindByIdentityNumber(ctx context.Context, identityNumber string) ([]entity.Loan, error)
}
