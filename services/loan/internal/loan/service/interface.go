package service

import (
	"context"

	"github.com/kupanglie/simple-loan-apps/services/loan/internal/loan/entity"
)

type Loan interface {
	Add(ctx context.Context, request entity.CreateLoanRequest) (entity.CreateLoanResponse, error)
	FindById(ctx context.Context, id int) (entity.Loan, error)
	FindByIdentityNumber(ctx context.Context, identityNumber string) ([]entity.Loan, error)
}
