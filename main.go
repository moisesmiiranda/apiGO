package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() { //endpoint que retornara os usuarios
	http.HandleFunc("/users", getUsers)
	fmt.Println("api is on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type User struct {
	ID   int    `json: "id"`
	Name string `json: "name"`
}

// funcao que chama o endpoint dos usuarios
func getUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content=type", "application/json")
	json.NewEncoder(w).Encode([]User{
		{
			ID:   1,
			Name: "Moisés",
		},
		{
			ID:   2,
			Name: "Mayara",
		},
		{
			ID:   3,
			Name: "Margareth",
		},
	})
}
