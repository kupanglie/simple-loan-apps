package entity

import "time"

type Loan struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Amount    int       `db:"amount"`
	Period    int       `db:"period"`
	Purpose   string    `db:"purpose"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}

type CreateLoanRequest struct {
	Name           string `json:"name"`
	IdentityNumber string `json:"identityNumber"`
	DateOfBirth    string `json:"dateOfBirth"`
	Sex            string `json:"sex"`
	Amount         int    `json:"amount"`
	Period         int    `json:"period"`
	Purpose        string `json:"purpose"`
}

type CreateLoanResponse struct {
	IsSuccess bool `json:"isSuccess"`
}
