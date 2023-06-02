// Copyright 2023 Serhii Khrystenko. All rights reserved.

/*
Package hasher implements user password verification.

This package is designed as an example of the Godoc
documentation and does not have any functionality:)
*/

package handlers

import (
	"fmt"
	"net/http"
)

func GetUserAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET method allowed", http.StatusBadRequest)
		return
	}

	for key, val := range usersTable {
		fmt.Fprintf(w, "\n%s = %v, %v", key, val.Id, val.Passhash)
	}
}
