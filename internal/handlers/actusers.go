package handlers

import (
	"fmt"
	"net/http"
)

var activeUsers = make(map[string]bool)

func ActiveUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorlog(w, "Only GET method allowed", http.StatusBadRequest)
		return
	}
	// if (len(activeUsers)) == 0 {
	// 	fmt.Fprintln(w, "No active users")
	// } else {
	fmt.Fprintln(w, "Active users:")
	i := 1
	for user, _ := range activeUsers {
		fmt.Fprintf(w, "%d - %s\n", i, user)
		i++
	}
	// }

}
