package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	loanRp "github.com/kupanglie/simple-loan-apps/services/loan/internal/loan/repository"
	loanSvc "github.com/kupanglie/simple-loan-apps/services/loan/internal/loan/service"
	loanUc "github.com/kupanglie/simple-loan-apps/services/loan/internal/loan/usecase"
	userRp "github.com/kupanglie/simple-loan-apps/services/loan/internal/user/repository"
)

const (
	SERVICE_USER_PORT = ":8900"
)

func main() {
	ctx := context.Background()

	if err := initRouter(ctx, SERVICE_USER_PORT); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func initRouter(ctx context.Context, servicePort string) error {
	db, err := sql.Open("mysql", "root:password@tcp(0.0.0.0:3307)/bank")
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	loanRp := loanRp.NewLoanRp(db)
	userRp := userRp.NewUserRp(db)
	loanSvc := loanSvc.NewLoanSvc(userRp, loanRp, db)
	loanUc := loanUc.NewLoanUC(loanSvc)

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/loan/create", loanUc.Add).Methods("POST")
	myRouter.HandleFunc("/loan/get-by-id/{id}", loanUc.FindById).Methods("GET")
	myRouter.HandleFunc("/loan/get-by-ktp/{identity-number}", loanUc.FindByIdentityNumber).Methods("GET")

	log.Printf("starting server on port %s...", servicePort)
	log.Fatal(http.ListenAndServe(servicePort, myRouter))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			<-ctx.Done()
		}
	}()

	return nil
}
