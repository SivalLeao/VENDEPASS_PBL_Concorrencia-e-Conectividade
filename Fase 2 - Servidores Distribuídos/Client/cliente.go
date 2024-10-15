package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var url string // = "http://localhost:8080"

// Estrutura de requisição de cliente (presente no servidor)
type Client_req struct {
	Id int `json:"id"`
	Rota string `json:"rota"`
}

// Estrutura de requisição de cadastro (presente no servidor)
type Cadastro_req struct{
	Id int `json:"id"`
	Nome string `json:"nome"`
}

// Função para requisição POST de cadastro com retorno de ID
func cadastrar(nome string) int {
	var cadastro = Cadastro_req{Nome: nome} // ID é gerado pelo servidor
	json_valor, err := json.Marshal(cadastro) // Serializa o JSON
	if err != nil {
		fmt.Println("Erro ao serializar o JSON:", err)
		return -1
	}

	resposta, err := http.Post(url+"/cadastro", "application/json", bytes.NewBuffer(json_valor)) // Faz a requisição POST
	if err != nil {
		fmt.Println("Erro ao fazer a requisição POST:", err)
		return -1
	}
	defer resposta.Body.Close()

	var resposta_map map[string]interface{} // Mapa para decodificar o JSON
	if err := json.NewDecoder(resposta.Body).Decode(&resposta_map); err != nil { // Decodifica o JSON
		fmt.Println("Erro ao decodificar o JSON:", err)
		return -1
	}

	id, ok := resposta_map["id"].(float64) // Converte o ID para int
	if !ok {
		fmt.Println("Erro ao converter o ID")
		return -1
	}

	return int(id)
}

// Função para requisição GET de rotas gerais com retorno de map
func rotas_todas() map[string]int {
	resposta, err := http.Get(url+"/rotas") // Faz a requisição GET
	if err != nil {
		fmt.Println("Erro ao fazer a requisição GET:", err)
		return nil
	}
	defer resposta.Body.Close()

	var resposta_map map[string]interface{} // Mapa para decodificar o JSON

	if err := json.NewDecoder(resposta.Body).Decode(&resposta_map); err != nil { // Decodifica o JSON
		fmt.Println("Erro ao decodificar o JSON:", err)
		return nil
	}

    rotas_interface, ok := resposta_map["rotas"].(map[string]interface{}) // Converte as rotas para map interface
    if !ok {
        fmt.Println("Erro ao converter as rotas")
        return nil
    }

	// Montagem do map de rotas para retorno
    rotas_recebidas := make(map[string]int) // Mapa para armazenar as rotas
    for rota, pertence := range rotas_interface { 
        if valor, ok := pertence.(float64); ok { 
            rotas_recebidas[rota] = int(valor) // Converte o valor para int
        } else {
            fmt.Println("Erro ao converter o valor da rota:", rota)
        }
    }

	return rotas_recebidas
}

// Função para requisição GET de rotas do cliente com retorno de slice
func rotas_cliente(id int) []string {
	resposta, err := http.Get(fmt.Sprintf("%s/rotas_cliente?id=%d", url, id)) // Faz a requisição GET com parâmetro de id
	if err != nil {
		fmt.Println("Erro ao fazer a requisição GET:", err)
		return nil
	}
	defer resposta.Body.Close()

	var resposta_map map[string]interface{} // Mapa para decodificar o JSON
	if err := json.NewDecoder(resposta.Body).Decode(&resposta_map); err != nil { // Decodifica o JSON
		fmt.Println("Erro ao decodificar o JSON:", err)
		return nil
	}

	rotas_interface, ok := resposta_map["rotas"].([]interface{}) // Converte as rotas para slice interface
	if !ok {
		resp, ok := resposta_map["error"].(string) // Verifica se houve erro
		if ok {
			fmt.Println(resp) // Imprime o erro
			return nil
		}
		fmt.Println("Erro ao converter as rotas") // Caso não haja erro, imprime mensagem de erro de conversão
		return nil
	}

	rotas_recebidas := make([]string, len(rotas_interface)) // Slice para armazenar as rotas
	for i, rota := range rotas_interface { // Converte as rotas para slice de string
		if valor, ok := rota.(string); ok {
			rotas_recebidas[i] = valor // Adiciona a rota ao slice
		} else {
			fmt.Println("Erro ao converter o valor da rota:", rota)
		}
	}

	return rotas_recebidas
}

// Função para requisição PATCH de compra de rota
func comprar(id int, rota string) {
	var cliente = Client_req{Id: id, Rota: rota} // Cria a estrutura de requisição
	json_valor, err := json.Marshal(cliente) // Serializa o JSON
	if err != nil {
		fmt.Println("Erro ao serializar o JSON:", err)
		return
	}

	req, err := http.NewRequest("PATCH", url+"/comprar_rota", bytes.NewBuffer(json_valor)) // Cria a requisição PATCH
	if err != nil {
		fmt.Println("Erro ao fazer a requisição PATCH:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json") // Adiciona o cabeçalho de Content-Type para JSON

	cliente_http := &http.Client{} // Cria o cliente HTTP
	resposta, err := cliente_http.Do(req) // Faz a requisição PATCH
	if err != nil {
		fmt.Println("Erro ao fazer a requisição PATCH:", err)
		return
	}
	defer resposta.Body.Close()

	var resposta_map map[string]interface{} // Mapa para decodificar o JSON
	if err := json.NewDecoder(resposta.Body).Decode(&resposta_map); err != nil { // Decodifica o JSON
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}
	resp_serv, ok := resposta_map["status"].(string) // Verifica se houve erro
	if !ok {
		fmt.Println("Ocorreu um erro")
	}
	fmt.Println(resp_serv) // Imprime a resposta do servidor
}

// Função para requisição PATCH de cancelamento de rota
func cancelar(id int, rota string) {
	var cliente = Client_req{Id: id, Rota: rota} // Cria a estrutura de requisição
	json_valor, err := json.Marshal(cliente) // Serializa o JSON
	if err != nil {
		fmt.Println("Erro ao serializar o JSON:", err)
		return
	}

	req, err := http.NewRequest("PATCH", url+"/cancelar_rota", bytes.NewBuffer(json_valor)) // Cria a requisição PATCH
	if err != nil {
		fmt.Println("Erro ao fazer a requisição PATCH:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json") // Adiciona o cabeçalho de Content-Type para JSON

	cliente_http := &http.Client{} // Cria o cliente HTTP
	resposta, err := cliente_http.Do(req) // Faz a requisição PATCH
	if err != nil {
		fmt.Println("Erro ao fazer a requisição PATCH:", err)
		return
	}
	defer resposta.Body.Close()

	var resposta_map map[string]interface{} // Mapa para decodificar o JSON
	if err := json.NewDecoder(resposta.Body).Decode(&resposta_map); err != nil { // Decodifica o JSON
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}
	resp_serv, ok := resposta_map["status"].(string) // Verifica se houve erro
	if !ok {
		fmt.Println("Ocorreu um erro")
	}
	fmt.Println(resp_serv) // Imprime a resposta do servidor
}

func main() {
	fmt.Println("Digite o endereço do servidor:")
	fmt.Scanln(&url)
	var entrada string
	var id int = -1
	for id == -1 {
		fmt.Println("Digite um nome de usuário:")
		fmt.Scanln(&entrada)
		id = cadastrar(entrada)
	}

	for { // Loop para interação e seleção de servidor
		fmt.Println("Conectando-se ao servidor", url)
		for entrada != "5"{ // Loop para interação com o servidor selecionado
			fmt.Println("Digite um comando [1 - Rotas, 2 - Minhas Rotas, 3 - Comprar, 4 - Cancelar, 5 - Sair]:")
			fmt.Scanln(&entrada) // Recebe a entrada do usuário
			switch entrada { // Seleção de comando
				case "1": // Caso 1 - Rotas
					rotas_recebidas := rotas_todas() // Recebe as rotas
					fmt.Println("Quantidade de rotas recebidas:", len(rotas_recebidas)) // Imprime a quantidade de rotas
					for rota, disp := range rotas_recebidas { // Imprime as rotas e disponibilidade
						fmt.Print("Rota: ", rota, " => ")
						if disp == 0 {
							fmt.Println("Disponível")
						} else {
							fmt.Println("Indisponível")
						}
					}
				case "2": // Caso 2 - Minhas Rotas
					rotas_recebidas := rotas_cliente(id) // Recebe as rotas do cliente
					fmt.Println("Quantidade de rotas recebidas:", len(rotas_recebidas)) // Imprime a quantidade de rotas
					for _, rota := range rotas_recebidas {
						fmt.Println("Rota:", rota)
					}
				case "3": // Caso 3 - Comprar
					fmt.Println("Digite a rota que deseja comprar:")
					fmt.Scanln(&entrada) // Recebe a rota
					comprar(id, entrada) // Chama a função de compra
				case "4": // Caso 4 - Cancelar
					fmt.Println("Digite a rota que deseja cancelar:")
					fmt.Scanln(&entrada) // Recebe a rota
					cancelar(id, entrada) // Chama a função de cancelamento
				case "5":
					fmt.Println("Saindo...") // Caso 5 - Sair
				default:
					fmt.Println("Comando inválido") // Caso padrão - Comando inválido
			}
		}
		fmt.Println("Digite o endereço do servidor:")
		fmt.Scanln(&url) // Recebe o endereço do servidor
		entrada = "" // Reseta a entrada
	}
}