package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gustavohenrique/e2e-tests-using-docker-cypress-vuejs-golang-postgres/backend/app"
)

func main() {
	port := "11080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatalf("Please set the DATABASE_URL variable. Ex.:\nDATABASE_URL=postgres://admin:password@127.0.0.1:5432/todos?sslmode=disable")
	}

	var application *app.Application
	application = app.New()
	application.Initialize(map[string]string{
		"databaseURL": databaseURL,
	})

	fmt.Println("Server running on", port)
	http.ListenAndServe(":"+port, application.EnableCORS())
}
