package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() { //endpoint que retornarar os usuarios
	http.HandleFunc("/users", getUsers)
	fmt.Println("api is on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type User struct {
	ID int 
	Name string
}
//funcao que chama o endpoint dos usuarios
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content=type", "application/json")
	json.NewEncoder(w).Encode([]User{{
		ID: 1,
        Name: "Moisés",
	}})
}
