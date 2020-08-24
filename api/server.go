package api

import (
	"./common"
	"./connect"
	"log"
	"net/http"
	"os"
	"time"
)

func Init() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	//version := os.Getenv("VERSION")

	connect.Init()
	defer connect.CloseConn()
	connect.Migrate()

	server := &http.Server{
		Addr:           host + ":" + port,
		Handler:        NewRoutes(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	common.WaitForShutdown(server)
}
