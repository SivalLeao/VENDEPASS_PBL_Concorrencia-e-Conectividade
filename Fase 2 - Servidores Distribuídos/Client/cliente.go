package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var url string = "http://localhost:8080"

func main() {
	// Requisição GET
	response, err := http.Get(url+"/rotas")
	if err != nil {
		fmt.Println("Erro ao fazer a requisição GET:", err)
		return
	}
	defer response.Body.Close()

	// body, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Println("Erro ao ler o corpo da resposta:", err)
	// 	return
	// }

	// fmt.Println("Resposta GET:", string(body))

	var rotas_recebidas map[string]int
	if err := json.NewDecoder(response.Body).Decode(&rotas_recebidas); err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}

	fmt.Println("Quantidade de rotas recebidas:", len(rotas_recebidas))
	fmt.Println("Rotas recebidas:", rotas_recebidas)

	// Requisição POST
	jsondata := map[string]string{
		"message": "Hello, World!"}
	jsonvalue, err := json.Marshal(jsondata)
	if err != nil {
		fmt.Println("Erro ao serializar o JSON:", err)
		return
	}

	response, err = http.Post(url+"/broadcast", "application/json", bytes.NewBuffer(jsonvalue))
	if err != nil {
		fmt.Println("Erro ao fazer a requisição POST:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err)
		return
	}

	fmt.Println("Resposta POST:", string(body))
}