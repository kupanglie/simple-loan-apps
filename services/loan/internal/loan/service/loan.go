package service

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	loanEntity "github.com/kupanglie/simple-loan-apps/services/loan/internal/loan/entity"
	loanRp "github.com/kupanglie/simple-loan-apps/services/loan/internal/loan/repository"
	userEntity "github.com/kupanglie/simple-loan-apps/services/loan/internal/user/entity"
	userRp "github.com/kupanglie/simple-loan-apps/services/loan/internal/user/repository"
)

type loan struct {
	db     *sql.DB
	loanRp loanRp.Loan
	userRp userRp.User
}

func NewLoanSvc(userRp userRp.User, loanRp loanRp.Loan, db *sql.DB) *loan {
	return &loan{
		db:     db,
		userRp: userRp,
		loanRp: loanRp,
	}
}

func (l *loan) Add(ctx context.Context, request loanEntity.CreateLoanRequest) (loanEntity.CreateLoanResponse, error) {
	userId := 0
	loanId := 0
	isSuccess := false

	tx, err := l.db.Begin()
	if err != nil {
		log.Println("[LOAN][SVC][Add][Begin] - ", err)
		return loanEntity.CreateLoanResponse{IsSuccess: isSuccess}, err
	}

	dateOfBirthParsed, err := time.Parse("02-01-2006", request.DateOfBirth)
	if err != nil {
		log.Println("[LOAN][SVC][Add][Parse] - ", err)
		return loanEntity.CreateLoanResponse{IsSuccess: isSuccess}, err
	}

	userReq := userEntity.User{
		Name:           request.Name,
		IdentityNumber: request.IdentityNumber,
		DateOfBirth:    dateOfBirthParsed,
		Sex:            request.Sex,
	}
	user, err := l.userRp.FindByIdentityNumber(ctx, request.IdentityNumber)
	if err != nil {
		tx.Rollback()

		log.Println("[LOAN][SVC][Add][FindByIdentityNumber] - ", err)
		return loanEntity.CreateLoanResponse{IsSuccess: isSuccess}, err
	}

	if user.ID > 0 {
		userId = user.ID
	} else {
		tx, userLastInsertedId, err := l.userRp.Add(ctx, tx, userReq)
		if err != nil {
			tx.Rollback()

			log.Println("[LOAN][SVC][Add][Add] - ", err)
			return loanEntity.CreateLoanResponse{IsSuccess: isSuccess}, err
		}

		userId = int(userLastInsertedId)
	}

	loanReq := loanEntity.Loan{
		UserID:  userId,
		Amount:  request.Amount,
		Period:  request.Period,
		Purpose: request.Purpose,
	}
	tx, loanId, err = l.loanRp.Add(ctx, tx, loanReq)
	if err != nil {
		tx.Rollback()

		log.Println("[LOAN][SVC][Add][Add] - ", err)
		return loanEntity.CreateLoanResponse{IsSuccess: isSuccess}, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()

		log.Println("[LOAN][SVC][Add][Commit] - ", err)
		return loanEntity.CreateLoanResponse{IsSuccess: isSuccess}, err
	}

	err = l.GenerateLoanDocument(ctx, loanId)
	if err != nil {
		tx.Rollback()

		log.Println("[LOAN][SVC][Add][GenerateLoanDocument] - FAIL")
		return loanEntity.CreateLoanResponse{IsSuccess: isSuccess}, err
	}

	return loanEntity.CreateLoanResponse{IsSuccess: true}, nil
}

func (l *loan) GenerateLoanDocument(ctx context.Context, loanId int) error {
	loan, err := l.FindById(ctx, loanId)
	if err != nil {
		log.Println("[LOAN][SVC][GenerateLoanDocument][FindById] - ", err)
		return err
	}

	user, err := l.userRp.FindById(ctx, loan.UserID)
	if err != nil {
		log.Println("[LOAN][SVC][GenerateLoanDocument][FindById] - ", err)
		return err
	}

	docData := [][]string{
		{"Loan ID", "Amount", "Period", "Purpose", "Name", "Identity Number", "Date of Birth", "Sex", "Loan Created At"},
		{strconv.Itoa(loan.ID), strconv.Itoa(loan.Amount), strconv.Itoa(loan.Period), loan.Purpose, user.Name, user.IdentityNumber, user.DateOfBirth.Format("02-01-2006"), user.Sex, loan.CreatedAt.Format("02-01-2006")},
	}

	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Println("[LOAN][SVC][GenerateLoanDocument][UserHomeDir] - ", err)
		return err
	}

	csvFile, err := os.Create(fmt.Sprintf("%s/data-%d.csv", os.DirFS(fmt.Sprint(homePath)), loanId))
	if err != nil {
		log.Println("[LOAN][SVC][GenerateLoanDocument][Create] - ", err)
		return err
	}

	csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range docData {
		err = csvwriter.Write(empRow)
		if err != nil {
			log.Println("[LOAN][SVC][GenerateLoanDocument][Write] - ", err)
			return err
		}
	}

	csvwriter.Flush()

	return nil
}

func (l *loan) FindById(ctx context.Context, id int) (loanEntity.Loan, error) {
	return l.loanRp.FindById(ctx, id)
}

func (l *loan) FindByIdentityNumber(ctx context.Context, identityNumber string) ([]loanEntity.Loan, error) {
	return l.loanRp.FindByIdentityNumber(ctx, identityNumber)
}
