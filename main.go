package main

import (
	"context"
	"log"
	"os"
	"transactionlabo/repository"
	"transactionlabo/usecase"
)

func main() {
	dsn := os.Getenv("DSN")
	dbRepository, err := repository.NewDBRepository(dsn)
	if err != nil {
		log.Fatalln(err)
	}
	txRepository := repository.NewTxDBRepository(dbRepository.DB())
	u := usecase.NewUsecase(dbRepository, txRepository)
	ctx := context.Background()
	err = u.CreateUser(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
