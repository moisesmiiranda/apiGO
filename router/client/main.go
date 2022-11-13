package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	ID   int    `json: "id"`
	Name string `json: "name"`
}

func main() {

	resp, err := http.Get("http://localhost:8080/users")
	// verifica se houve problema no endpoint (erro de conexão ou etc...)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if resp.StatusCode != 200 { //se o status code não for o ideal mostrar erro
		fmt.Println("Not sucess", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body) // joga no body da request os dados de obteve como retorno

	var response []User
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Erro ao recuperar dados", err.Error())
		return
	}

	fmt.Println(response)
}
