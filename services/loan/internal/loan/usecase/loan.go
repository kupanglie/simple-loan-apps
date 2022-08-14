package usecase

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kupanglie/simple-loan-apps/services/loan/internal/constant"
	"github.com/kupanglie/simple-loan-apps/services/loan/internal/helper"
	"github.com/kupanglie/simple-loan-apps/services/loan/internal/loan/entity"
	"github.com/kupanglie/simple-loan-apps/services/loan/internal/loan/service"
)

type loan struct {
	loanSvc service.Loan
}

func NewLoanUC(loanSvc service.Loan) *loan {
	return &loan{
		loanSvc: loanSvc,
	}
}

func (l *loan) Add(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	reqBody, _ := ioutil.ReadAll(r.Body)
	request := entity.CreateLoanRequest{}
	json.Unmarshal(reqBody, &request)

	if !validateName(ctx, request.Name) {
		json.NewEncoder(w).Encode(helper.ResponseGenerator(entity.CreateLoanResponse{}, constant.INCORRECT_NAME))
		return
	}
	if !validateIdentityNumber(ctx, request.Sex, request.DateOfBirth, request.IdentityNumber) {
		json.NewEncoder(w).Encode(helper.ResponseGenerator(entity.CreateLoanResponse{}, constant.INCORRECT_IDENTITY_NUMBER))
		return
	}
	if !validateAmount(ctx, request.Amount) {
		json.NewEncoder(w).Encode(helper.ResponseGenerator(entity.CreateLoanResponse{}, constant.INCORRECT_AMOUNT))
		return
	}
	if !validatePurpose(ctx, request.Purpose) {
		json.NewEncoder(w).Encode(helper.ResponseGenerator(entity.CreateLoanResponse{}, constant.INCORRECT_PURPOSE))
		return
	}

	response, err := l.loanSvc.Add(ctx, request)
	if err != nil {
		log.Println("[LOAN][UC][Add][Add] - ", err)
		json.NewEncoder(w).Encode(helper.ResponseGenerator(entity.CreateLoanResponse{}, constant.INTERNAL_SERVER_ERROR))
		return
	}

	json.NewEncoder(w).Encode(helper.ResponseGenerator(response, 0))
	return
}

func (l *loan) FindById(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("[LOAN][UC][FindById][Atoi] - ", err)
		json.NewEncoder(w).Encode(helper.ResponseGenerator(entity.Loan{}, constant.INTERNAL_SERVER_ERROR))
		return
	}

	response, err := l.loanSvc.FindById(ctx, id)
	if err != nil {
		log.Println("[LOAN][UC][FindById][FindById] - ", err)
		json.NewEncoder(w).Encode(helper.ResponseGenerator(entity.Loan{}, constant.INTERNAL_SERVER_ERROR))
		return
	}

	json.NewEncoder(w).Encode(helper.ResponseGenerator(response, 0))
	return
}

func (l *loan) FindByIdentityNumber(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	vars := mux.Vars(r)
	identityNumber := vars["identity-number"]

	response, err := l.loanSvc.FindByIdentityNumber(ctx, identityNumber)
	if err != nil {
		log.Println("[LOAN][UC][FindById][FindById] - ", err)
		json.NewEncoder(w).Encode(helper.ResponseGenerator([]entity.Loan{}, constant.INTERNAL_SERVER_ERROR))
		return
	}

	json.NewEncoder(w).Encode(helper.ResponseGenerator(response, 0))
	return
}
