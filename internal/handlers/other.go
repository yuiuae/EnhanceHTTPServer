// Copyright 2023 Serhii Khrystenko. All rights reserved.

/*
Package hasher implements user password verification.

This package is designed as an example of the Godoc
documentation and does not have any functionality:)
*/

package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET method allowed", http.StatusBadRequest)
		return
	}

	var msg string = `
	<html>
	<h1>Welcome on main page!</h1>
	</html>	
	`
	w.Write([]byte(msg))
}

func GetUserAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET method allowed", http.StatusBadRequest)
		return
	}

	for key, val := range usersTable {
		fmt.Fprintf(w, "\n%s: ID = %v, PassHash = %v", key, val.Id, val.Passhash)
	}
}

var upgrader = websocket.Upgrader{} // use default options

func RequestWithToken(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
