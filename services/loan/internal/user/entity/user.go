package entity

import "time"

type User struct {
	ID             int       `db:"id"`
	Name           string    `db:"name"`
	IdentityNumber string    `db:"identity_number"`
	DateOfBirth    time.Time `db:"date_of_birth"`
	Sex            string    `db:"sex"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	DeletedAt      time.Time `db:"deleted_at"`
}
