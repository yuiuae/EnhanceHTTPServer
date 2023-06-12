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
	http.HandleFunc("/", middleware.middleLog(handlers.Index))
	// http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/user", middleware.middleLog(handlers.UserCreate))
	http.HandleFunc("/user/login", middleware.middleLog(handlers.UserLogin))
	http.HandleFunc("/admin", middleware.middleLog(handlers.GetUserAll))

	http.HandleFunc("/chat", middleware.middleLog(handlers.RequestWithToken))

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
