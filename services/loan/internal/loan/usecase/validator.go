package usecase

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/kupanglie/simple-loan-apps/services/loan/internal/constant"
)

func validateName(ctx context.Context, name string) bool {
	return len(strings.Split(name, " ")) >= 2
}

func validateIdentityNumber(ctx context.Context, sex string, dateOfBirth string, identityNumber string) bool {
	dateOfBirthParsed, err := time.Parse("02-01-2006", dateOfBirth)
	if err != nil {
		log.Println("[LOAN][UC][validateIdentityNumber][Parse] - ", err)
		return false
	}

	if sex == constant.MALE {
		return strings.Join(strings.Split(identityNumber, "")[6:12], "") == dateOfBirthParsed.Format("020106")
	} else {
		date, _ := strconv.Atoi(dateOfBirthParsed.Format("02"))

		return strings.Join(strings.Split(identityNumber, "")[6:12], "") == strings.Join([]string{strconv.Itoa(date + 40), dateOfBirthParsed.Format("0106")}, "")
	}
}

func validateAmount(ctx context.Context, amount int) bool {
	return amount >= 1000000 && amount <= 10000000
}

func validatePurpose(ctx context.Context, purpose string) bool {
	subPurpose := []string{"vacation", "renovation", "electronics", "wedding", "rent", "car", "investment"}

	for _, val := range subPurpose {
		if strings.Contains(purpose, val) {
			return true
		}
	}

	return false
}
