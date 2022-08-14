package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/kupanglie/simple-loan-apps/services/loan/internal/helper"
	"github.com/kupanglie/simple-loan-apps/services/loan/internal/loan/entity"
)

type loan struct {
	db *sql.DB
}

func NewLoanRp(db *sql.DB) *loan {
	return &loan{
		db: db,
	}
}

func (l *loan) Add(ctx context.Context, tx *sql.Tx, request entity.Loan) (*sql.Tx, int, error) {
	query := `INSERT INTO loans (user_id, amount, period, purpose, created_at) VALUES (?, ?, ?, ?, NOW())`
	response, err := tx.ExecContext(ctx, query, request.UserID, request.Amount, request.Period, request.Purpose)
	if err != nil {
		tx.Rollback()

		log.Println("[LOAN][RP][Add][ExecContext] - ", err)
		return tx, 0, err
	}

	lastInsertedId, err := response.LastInsertId()
	if err != nil {
		tx.Rollback()

		log.Println("[LOAN][RP][Add][LastInsertId] - ", err)
		return tx, 0, err
	}
	return tx, int(lastInsertedId), nil
}

func (l *loan) FindById(ctx context.Context, id int) (entity.Loan, error) {
	loan := entity.Loan{}

	query := `SELECT id, user_id, amount, period, purpose, created_at, updated_at, deleted_at FROM loans WHERE id = ? AND deleted_at IS NULL`
	rows, err := l.db.QueryContext(ctx, query, id)
	if err != nil {
		log.Println("[LOAN][RP][FindById][QueryContext] - ", err)
		return loan, err
	}

	defer rows.Close()

	for rows.Next() {
		var createdAt, updatedAt, deletedAt sql.NullString
		err = rows.Scan(&loan.ID, &loan.UserID, &loan.Amount, &loan.Period, &loan.Purpose, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			log.Println("[LOAN][RP][FindById][Scan] - ", err)
			return loan, err
		}

		loan.CreatedAt = helper.CastSQLTimeToTime(createdAt.String)
		loan.UpdatedAt = helper.CastSQLTimeToTime(updatedAt.String)
		loan.DeletedAt = helper.CastSQLTimeToTime(deletedAt.String)
	}

	return loan, nil
}

func (l *loan) FindByIdentityNumber(ctx context.Context, identityNumber string) ([]entity.Loan, error) {
	var loans []entity.Loan

	query := `SELECT l.id, l.user_id, l.amount, l.period, l.purpose, l.created_at, l.updated_at, l.deleted_at FROM loans l JOIN users u ON l.user_id = u.id  WHERE u.identity_number = ? AND l.deleted_at IS NULL`
	rows, err := l.db.QueryContext(ctx, query, identityNumber)
	if err != nil {
		log.Println("[LOAN][RP][FindByIdentityNumber][QueryContext] - ", err)
		return loans, err
	}

	defer rows.Close()

	for rows.Next() {
		var loan entity.Loan
		var createdAt, updatedAt, deletedAt sql.NullString

		err = rows.Scan(&loan.ID, &loan.UserID, &loan.Amount, &loan.Period, &loan.Purpose, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			log.Println("[LOAN][RP][FindByIdentityNumber][Scan] - ", err)
			return loans, err
		}

		loan.CreatedAt = helper.CastSQLTimeToTime(createdAt.String)
		loan.UpdatedAt = helper.CastSQLTimeToTime(updatedAt.String)
		loan.DeletedAt = helper.CastSQLTimeToTime(deletedAt.String)

		loans = append(loans, loan)
	}

	return loans, nil
}
