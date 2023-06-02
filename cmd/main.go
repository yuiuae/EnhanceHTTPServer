package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yuiuae/EnhanceHTTPServer/internal/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// "UserCreate" and "CheckUser" are handler that we will implement
	http.HandleFunc("/user", handlers.UserCreate)
	http.HandleFunc("/user/login", handlers.UserLogin)
	http.HandleFunc("/admin", handlers.GetUserAll)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
