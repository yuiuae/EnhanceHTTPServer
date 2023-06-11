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
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	token := r.URL.Query().Get("token")
	username, err := parseToken(token, []byte(tokenSecretKey))
	if err != nil {
		log.Print("parse:", err)
		return
	}

	usersTable[username].Token = ""
	usersTable[username].ExpireAt = 0

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

// parse token from a websocket user connection request
func parseToken(accessToken string, signingKey []byte) (string, error) {
	// claims := jwt.MapClaims{}
	// token, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) { return []byte(signingKey), nil })
	// fmt.Println(token.Claims)
	// if err != nil {
	// 	return "", err
	// }
	// for key, val := range claims {
	// 	fmt.Printf("Key: %v, value: %v\n", key, val)
	// }

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
