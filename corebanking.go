package main

import (
	delivery "core-banking/app/delivery"
	repository "core-banking/app/repository"
	usecase "core-banking/app/usecase"
	"core-banking/config"
	"flag"
	"net/http"
	"time"
)

func main() {

	// DB configuration
	dbConfig := config.ConfigDB()

	// Router configuration
	router := config.ConfigRouter()

	rCoreBankingRepo := repository.NewRepository(dbConfig)

	uCoreBankingRepo := usecase.NewUsecase(rCoreBankingRepo)

	delivery.NewHTTPHandler(router, uCoreBankingRepo)

	s := &http.Server{
		Addr:         getPortNumber(),
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	router.Logger.Fatal(router.StartServer(s))
}

func getPortNumber() string {
	port := flag.String("port", "8000", "Port number")
	flag.Parse()
	if port == nil || len(*port) == 0 {
		panic("Please enter port")
	}
	return ":" + *port
}
