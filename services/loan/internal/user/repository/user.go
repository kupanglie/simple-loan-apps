package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/kupanglie/simple-go-app/service/user/helper"
	"github.com/kupanglie/simple-loan-apps/services/loan/internal/user/entity"
)

type user struct {
	db *sql.DB
}

func NewUserRp(db *sql.DB) *user {
	return &user{
		db: db,
	}
}

func (u *user) Add(ctx context.Context, tx *sql.Tx, request entity.User) (*sql.Tx, int64, error) {

	query := `INSERT INTO users (name, identity_number, date_of_birth, sex, created_at) VALUES (?, ?, ?, ?, NOW())`
	result, err := tx.ExecContext(ctx, query, request.Name, request.IdentityNumber, request.DateOfBirth, request.Sex)
	if err != nil {
		tx.Rollback()

		log.Println("[USER][RP][Add][ExecContext] - ", err)
		return tx, 0, err
	}

	lastInsertedId, _ := result.LastInsertId()
	if lastInsertedId == 0 {
		tx.Rollback()

		log.Println("[USER][RP][Add][LastInsertId] - Last inserted id is 0")
		return tx, 0, errors.New("last inserted id is 0")
	}

	return tx, lastInsertedId, nil
}

func (u *user) FindById(ctx context.Context, id int) (entity.User, error) {
	user := entity.User{}

	query := `SELECT id, name, identity_number, date_of_birth, sex, created_at, updated_at, deleted_at FROM users WHERE id = ? AND deleted_at IS NULL`
	rows, err := u.db.QueryContext(ctx, query, id)
	if err != nil {
		log.Println("[USER][RP][FindById][QueryContext] - ", err)
		return user, err
	}

	defer rows.Close()

	for rows.Next() {
		var dateOfBirth, createdAt, updatedAt, deletedAt sql.NullString
		err = rows.Scan(&user.ID, &user.Name, &user.IdentityNumber, &dateOfBirth, &user.Sex, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			log.Println("[USER][RP][FindById][Scan] - ", err)
			return user, err
		}

		user.DateOfBirth = helper.CastSQLTimeToDate(dateOfBirth.String)
		user.CreatedAt = helper.CastSQLTimeToTime(createdAt.String)
		user.UpdatedAt = helper.CastSQLTimeToTime(updatedAt.String)
		user.DeletedAt = helper.CastSQLTimeToTime(deletedAt.String)
	}

	return user, nil
}

func (u *user) FindByIdentityNumber(ctx context.Context, identityNumber string) (entity.User, error) {
	user := entity.User{}

	query := `SELECT id, name, identity_number, date_of_birth, sex, created_at, updated_at, deleted_at FROM users WHERE identity_number = ? AND deleted_at IS NULL`
	rows, err := u.db.QueryContext(ctx, query, identityNumber)
	if err != nil {
		log.Println("[USER][RP][FindByIdentityNumber][QueryContext] - ", err)
		return user, err
	}

	defer rows.Close()

	for rows.Next() {
		var dateOfBirth, createdAt, updatedAt, deletedAt sql.NullString
		err = rows.Scan(&user.ID, &user.Name, &user.IdentityNumber, &dateOfBirth, &user.Sex, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			log.Println("[USER][RP][FindByIdentityNumber][Scan] - ", err)
			return user, err
		}

		user.DateOfBirth = helper.CastSQLTimeToDate(dateOfBirth.String)
		user.CreatedAt = helper.CastSQLTimeToTime(createdAt.String)
		user.UpdatedAt = helper.CastSQLTimeToTime(updatedAt.String)
		user.DeletedAt = helper.CastSQLTimeToTime(deletedAt.String)
	}

	return user, nil
}
