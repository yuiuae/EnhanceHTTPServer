// Copyright 2023 Serhii Khrystenko. All rights reserved.

/*
Package handler implements user password verification.

This package is designed as an example of the Godoc
documentation and does not have any functionality:)
*/

package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/yuiuae/EnhanceHTTPServer/pkg/hasher"
)

// calls per hour allowed by the user
var callperhour int = 100

// token validity time (in hours)
var tokentime = 1

// Table with users
var usersTable = map[string]UserInfo{}

type UserInfo struct {
	Passhash string
	Id       string
}

// Create a struct that models the structure for a user creating
// Request
type CrRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// Response
type CrResponse struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
}

// Create new use and add to user table
func UserCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method allowed", http.StatusBadRequest)
		return
	}
	req := &CrRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, "Bad request, empty username or password", http.StatusBadRequest)
		return
	}
	if len(req.UserName) <= 4 {
		http.Error(w, "Username should be at least 4 characters", http.StatusBadRequest)
		return
	}

	if len(req.Password) <= 8 {
		http.Error(w, "Password should be at least 8 characters", http.StatusBadRequest)
		return
	}
	_, b := usersTable[req.UserName]
	if b {
		http.Error(w, "A user with this name already exists", http.StatusConflict)
		return
	}

	// Hash the password using the bcrypt algorithm
	hashedPassword, err := hasher.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Internal Server Error (hash error)", http.StatusInternalServerError)
		return
	}

	// Generate UUID
	uid, err := uuid.NewUUID()
	if err != nil {
		http.Error(w, "Internal Server Error (UUID error)", http.StatusInternalServerError)
		return
	}

	// Create response
	resp := &CrResponse{uid.String(), req.UserName}
	err = json.NewEncoder(w).Encode(&resp) //&resp
	if err != nil {
		http.Error(w, "Internal Server Error (json Encoder error)", http.StatusInternalServerError)
		return
	}
	// add new user to to user table
	usersTable[req.UserName] = UserInfo{hashedPassword, uid.String()}
}

// Create a struct that models the structure for a user login
// Request
type LogRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// Response
type LogResponse struct {
	URL string `json:"username"`
}

// Check username and passoword in user table
func UserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method allowed", http.StatusBadRequest)
		return
	}
	req := &LogRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, "Bad request, empty username or password", http.StatusBadRequest)
		return
	}

	ok := hasher.CheckPasswordHash(usersTable[req.UserName].Passhash, req.Password)
	if !ok {
		http.Error(w, "Invalid username/password", http.StatusBadRequest)
		return
	}

	url := "ws://fancy-chat.io/ws&token=one-time-token"
	resp := &LogResponse{url}
	w.Header().Add("X-Rate-Limit", strconv.Itoa(callperhour))
	w.Header().Add("X-Expires-After", time.Now().UTC().Add(time.Hour*time.Duration(tokentime)).String())
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "Internal Server Error (json Encoder)", http.StatusInternalServerError)
		return
	}
}
