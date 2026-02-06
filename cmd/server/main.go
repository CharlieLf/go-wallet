package main

import (
	"log"
	"net/http"

	"go-wallet/internal/config"
	"go-wallet/internal/db"
	"go-wallet/internal/router"
	"go-wallet/internal/wallet"
)

func main() {

	cfg := config.Load()

	dbConn, err := db.New(cfg.DBDSN)
	if err != nil {
		log.Fatal(err)
	}

	repo := wallet.NewRepository(dbConn)
	service := wallet.NewService(repo)
	handler := wallet.NewHandler(service)

	r := router.Register(handler)

	log.Println("wallet api running on :" + cfg.Port)

	http.ListenAndServe(":"+cfg.Port, r)
}
