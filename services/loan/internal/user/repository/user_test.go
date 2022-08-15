package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"reflect"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/kupanglie/simple-go-app/service/user/helper"
	"github.com/kupanglie/simple-loan-apps/services/loan/internal/user/entity"
	"github.com/kupanglie/simple-loan-apps/services/loan/internal/user/repository"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func Test_user_FindById(t *testing.T) {
	ctx := context.Background()
	dob := helper.CastSQLTimeToDate("1995-01-01")
	timeNow := time.Now().Format("2006-01-02 15:04:05")

	type fields struct {
		db func() *sql.DB
	}
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.User
		wantErr bool
	}{
		{
			name: "Error QueryContext",
			fields: fields{
				db: func() *sql.DB {
					db, dbmock := NewMock()
					dbmock.ExpectQuery("SELECT id, name, identity_number, date_of_birth, sex, created_at, updated_at, deleted_at FROM users").WillReturnError(errors.New("errors"))

					return db
				},
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			want:    entity.User{},
			wantErr: true,
		},
		{
			name: "Error Scan",
			fields: fields{
				db: func() *sql.DB {
					db, dbmock := NewMock()
					rows := sqlmock.NewRows([]string{"id", "name", "identity_number", "date_of_birth", "sex", "created_at", "updated_at", "deleted_at"}).
						AddRow("asdasd", "alfin", "1234560101954567", "1995-01-01", "MALE", time.Now().Format("2006-01-02 15:04:05"), nil, nil)

					dbmock.ExpectQuery("SELECT (.+) FROM users").WillReturnRows(rows)

					return db
				},
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			want:    entity.User{},
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				db: func() *sql.DB {
					db, dbmock := NewMock()
					rows := sqlmock.NewRows([]string{"id", "name", "identity_number", "date_of_birth", "sex", "created_at", "updated_at", "deleted_at"}).
						AddRow(1, "alfin", "1234560101954567", "1995-01-01", "MALE", timeNow, nil, nil)
					dbmock.ExpectQuery("SELECT id, name, identity_number, date_of_birth, sex, created_at, updated_at, deleted_at FROM users").WillReturnRows(rows)

					return db
				},
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			want: entity.User{
				ID:             1,
				Name:           "alfin",
				IdentityNumber: "1234560101954567",
				DateOfBirth:    dob,
				Sex:            "MALE",
				CreatedAt:      helper.CastSQLTimeToTime(timeNow),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := repository.NewUserRp(tt.fields.db())

			got, err := u.FindById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("user.FindById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_user_FindByIdentityNumber(t *testing.T) {
	ctx := context.Background()
	dob := helper.CastSQLTimeToDate("1995-01-01")
	timeNow := time.Now().Format("2006-01-02 15:04:05")

	type fields struct {
		db func() *sql.DB
	}
	type args struct {
		ctx            context.Context
		identityNumber string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.User
		wantErr bool
	}{
		{
			name: "Error QueryContext",
			fields: fields{
				db: func() *sql.DB {
					db, dbmock := NewMock()
					dbmock.ExpectQuery("SELECT id, name, identity_number, date_of_birth, sex, created_at, updated_at, deleted_at FROM users").WillReturnError(errors.New("errors"))

					return db
				},
			},
			args: args{
				ctx:            ctx,
				identityNumber: "1234560101954567",
			},
			want:    entity.User{},
			wantErr: true,
		},
		{
			name: "Error Scan",
			fields: fields{
				db: func() *sql.DB {
					db, dbmock := NewMock()
					rows := sqlmock.NewRows([]string{"id", "name", "identity_number", "date_of_birth", "sex", "created_at", "updated_at", "deleted_at"}).
						AddRow("asdasd", "alfin", "1234560101954567", "1995-01-01", "MALE", time.Now().Format("2006-01-02 15:04:05"), nil, nil)

					dbmock.ExpectQuery("SELECT (.+) FROM users").WillReturnRows(rows)

					return db
				},
			},
			args: args{
				ctx:            ctx,
				identityNumber: "1234560101954567",
			},
			want:    entity.User{},
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				db: func() *sql.DB {
					db, dbmock := NewMock()
					rows := sqlmock.NewRows([]string{"id", "name", "identity_number", "date_of_birth", "sex", "created_at", "updated_at", "deleted_at"}).
						AddRow(1, "alfin", "1234560101954567", "1995-01-01", "MALE", timeNow, nil, nil)

					dbmock.ExpectQuery("SELECT id, name, identity_number, date_of_birth, sex, created_at, updated_at, deleted_at FROM users").WillReturnRows(rows)

					return db
				},
			},
			args: args{
				ctx:            ctx,
				identityNumber: "1234560101954567",
			},
			want: entity.User{
				ID:             1,
				Name:           "alfin",
				IdentityNumber: "1234560101954567",
				DateOfBirth:    dob,
				Sex:            "MALE",
				CreatedAt:      helper.CastSQLTimeToTime(timeNow),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := repository.NewUserRp(tt.fields.db())

			got, err := u.FindByIdentityNumber(tt.args.ctx, tt.args.identityNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.FindByIdentityNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("user.FindByIdentityNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_user_Add(t *testing.T) {
	ctx := context.Background()
	dob := helper.CastSQLTimeToDate("1995-01-01")

	type fields struct {
		db func() *sql.DB
	}
	type args struct {
		ctx     context.Context
		request entity.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "Error ExecContext",
			fields: fields{
				db: func() *sql.DB {
					db, dbmock := NewMock()

					dbmock.ExpectBegin()
					dbmock.ExpectExec("INSERT INTO users").WillReturnError(errors.New("error"))
					dbmock.ExpectRollback()

					return db
				},
			},
			args: args{
				ctx: ctx,
				request: entity.User{
					Name:           "alfin",
					IdentityNumber: "1234560101954567",
					DateOfBirth:    dob,
					Sex:            "MALE",
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "LastInsertId is 0",
			fields: fields{
				db: func() *sql.DB {
					db, dbmock := NewMock()

					dbmock.ExpectBegin()
					dbmock.ExpectExec("INSERT INTO users").WillReturnResult(sqlmock.NewResult(0, 0))
					dbmock.ExpectRollback()

					return db
				},
			},
			args: args{
				ctx: ctx,
				request: entity.User{
					Name:           "alfin",
					IdentityNumber: "1234560101954567",
					DateOfBirth:    dob,
					Sex:            "MALE",
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				db: func() *sql.DB {
					db, dbmock := NewMock()

					dbmock.ExpectBegin()
					dbmock.ExpectExec("INSERT INTO users").WillReturnResult(sqlmock.NewResult(1, 1))

					return db
				},
			},
			args: args{
				ctx: ctx,
				request: entity.User{
					Name:           "alfin",
					IdentityNumber: "1234560101954567",
					DateOfBirth:    dob,
					Sex:            "MALE",
				},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := repository.NewUserRp(tt.fields.db())

			tx, _ := tt.fields.db().Begin()

			_, got1, err := u.Add(tt.args.ctx, tx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("user.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got1 != tt.want {
				t.Errorf("user.Add() got1 = %v, want %v", got1, tt.want)
			}
		})
	}
}
