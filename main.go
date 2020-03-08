package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Tarea struct {
	UserId    int    `json:userId`
	Id        int    `ìd`
	Title     string `title`
	Completed bool   `completed`
}

func main() {

	var urlApi = "https://jsonplaceholder.typicode.com/todos"

	var cliente = &http.Client{Timeout: 10 * time.Second}

	response, err := cliente.Get(urlApi)

	if err != nil {
		panic(err.Error())
	}

	var tareas []Tarea

	err = json.NewDecoder(response.Body).Decode(&tareas)

	fmt.Println(tareas)
}
