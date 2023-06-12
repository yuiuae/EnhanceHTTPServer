package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// websocket connection request
func RequestWithToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorlog(w, "Only GET method allowed", http.StatusBadRequest)
		return
	}

	reqToken := r.URL.Query().Get("token")
	username, err := parseToken(reqToken, []byte(tokenSecretKey))
	if err != nil {
		errorlog(w, "Internal Server Error (parse)", http.StatusInternalServerError)
		return
	}

	_, exists := usersTable[username]
	if !exists {
		errorlog(w, "User error", http.StatusBadRequest)
		return
	}

	if reqToken != usersTable[username].Token {
		errorlog(w, "Invalid token", http.StatusBadRequest)
		return
	}

	usersTable[username].Token = ""
	usersTable[username].ExpireAt = 0
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		errorlog(w, "Internal Server Error (upgrade)", http.StatusInternalServerError)
		return
	}
	activeUsers[username] = true
	defer delete(activeUsers, username)
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			errorlog(w, "Internal Server Error (ws read)", http.StatusInternalServerError)
			break
		}
		log.Printf("recv: %s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			errorlog(w, "Internal Server Error (ws write)", http.StatusInternalServerError)
			break
		}
	}
}

// parse token from a websocket user connection request
func parseToken(accessToken string, signingKey []byte) (string, error) {
	var username string
	token, _, err := new(jwt.Parser).ParseUnverified(accessToken, jwt.MapClaims{})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		username = fmt.Sprint(claims["username"])
	}

	if username == "" {
		return "", fmt.Errorf("invalid token payload")
	}
	return username, nil

}
