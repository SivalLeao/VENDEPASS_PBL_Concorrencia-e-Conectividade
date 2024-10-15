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
	var cadastro = Cadastro_req{Nome: nome}
	json_valor, err := json.Marshal(cadastro)
	if err != nil {
		fmt.Println("Erro ao serializar o JSON:", err)
		return -1
	}

	resposta, err := http.Post(url+"/cadastro", "application/json", bytes.NewBuffer(json_valor))
	if err != nil {
		fmt.Println("Erro ao fazer a requisição POST:", err)
		return -1
	}
	defer resposta.Body.Close()

	var resposta_map map[string]interface{}
	if err := json.NewDecoder(resposta.Body).Decode(&resposta_map); err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return -1
	}

	id, ok := resposta_map["id"].(float64)
	if !ok {
		fmt.Println("Erro ao converter o ID")
		return -1
	}

	return int(id)
}

// Função para requisição GET de rotas gerais com retorno de map
func rotas_todas() map[string]int {
	resposta, err := http.Get(url+"/rotas")
	if err != nil {
		fmt.Println("Erro ao fazer a requisição GET:", err)
		return nil
	}
	defer resposta.Body.Close()

	var resposta_map map[string]interface{}

	if err := json.NewDecoder(resposta.Body).Decode(&resposta_map); err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return nil
	}

    rotas_interface, ok := resposta_map["rotas"].(map[string]interface{})
    if !ok {
        fmt.Println("Erro ao converter as rotas")
        return nil
    }

    rotas_recebidas := make(map[string]int)
    for rota, pertence := range rotas_interface {
        if valor, ok := pertence.(float64); ok {
            rotas_recebidas[rota] = int(valor)
        } else {
            fmt.Println("Erro ao converter o valor da rota:", rota)
        }
    }

	return rotas_recebidas
}

// Função para requisição GET de rotas do cliente com retorno de slice
func rotas_cliente(id int) []string {
	resposta, err := http.Get(fmt.Sprintf("%s/rotas_cliente?id=%d", url, id))
	if err != nil {
		fmt.Println("Erro ao fazer a requisição GET:", err)
		return nil
	}
	defer resposta.Body.Close()

	var resposta_map map[string]interface{}
	if err := json.NewDecoder(resposta.Body).Decode(&resposta_map); err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return nil
	}

	rotas_interface, ok := resposta_map["rotas"].([]interface{})
	if !ok {
		resp, ok := resposta_map["error"].(string)
		if ok {
			fmt.Println(resp)
			return nil
		}
		fmt.Println("Erro ao converter as rotas")
		return nil
	}

	rotas_recebidas := make([]string, len(rotas_interface))
	for i, rota := range rotas_interface {
		if valor, ok := rota.(string); ok {
			rotas_recebidas[i] = valor
		} else {
			fmt.Println("Erro ao converter o valor da rota:", rota)
		}
	}

	return rotas_recebidas
}

// Função para requisição PATCH de compra de rota
func comprar(id int, rota string) {
	var cliente = Client_req{Id: id, Rota: rota}
	json_valor, err := json.Marshal(cliente)
	if err != nil {
		fmt.Println("Erro ao serializar o JSON:", err)
		return
	}

	req, err := http.NewRequest("PATCH", url+"/comprar_rota", bytes.NewBuffer(json_valor))
	if err != nil {
		fmt.Println("Erro ao fazer a requisição PATCH:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	cliente_http := &http.Client{}
	resposta, err := cliente_http.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição PATCH:", err)
		return
	}
	defer resposta.Body.Close()

	var resposta_map map[string]interface{}
	if err := json.NewDecoder(resposta.Body).Decode(&resposta_map); err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}
	resp_serv, ok := resposta_map["status"].(string)
	if !ok {
		fmt.Println("Ocorreu um erro")
		return
	}
	fmt.Println(resp_serv)
}

// Função para requisição PATCH de cancelamento de rota
func cancelar(id int, rota string) {
	var cliente = Client_req{Id: id, Rota: rota}
	json_valor, err := json.Marshal(cliente)
	if err != nil {
		fmt.Println("Erro ao serializar o JSON:", err)
		return
	}

	req, err := http.NewRequest("PATCH", url+"/cancelar_rota", bytes.NewBuffer(json_valor))
	if err != nil {
		fmt.Println("Erro ao fazer a requisição PATCH:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	cliente_http := &http.Client{}
	resposta, err := cliente_http.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição PATCH:", err)
		return
	}
	defer resposta.Body.Close()

	var resposta_map map[string]interface{}
	if err := json.NewDecoder(resposta.Body).Decode(&resposta_map); err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}
	resp_serv, ok := resposta_map["status"].(string)
	if !ok {
		fmt.Println("Ocorreu um erro")
		return
	}
	fmt.Println(resp_serv)
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

	for {
		fmt.Println("Conectando-se ao servidor", url)
		for entrada != "5"{
			fmt.Println("Digite um comando [1 - Rotas, 2 - Minhas Rotas, 3 - Comprar, 4 - Cancelar, 5 - Sair]:")
			fmt.Scanln(&entrada)
			switch entrada {
				case "1":
					rotas_recebidas := rotas_todas()
					fmt.Println("Quantidade de rotas recebidas:", len(rotas_recebidas))
					for rota, disp := range rotas_recebidas {
						fmt.Print("Rota: ", rota, " => ")
						if disp == 0 {
							fmt.Println("Disponível")
						} else {
							fmt.Println("Indisponível")
						}
					}
				case "2":
					rotas_recebidas := rotas_cliente(id)
					fmt.Println("Quantidade de rotas recebidas:", len(rotas_recebidas))
					for _, rota := range rotas_recebidas {
						fmt.Println("Rota:", rota)
					}
				case "3":
					fmt.Println("Digite a rota que deseja comprar:")
					fmt.Scanln(&entrada)
					comprar(id, entrada)
				case "4":
					fmt.Println("Digite a rota que deseja cancelar:")
					fmt.Scanln(&entrada)
					cancelar(id, entrada)
				case "5":
					fmt.Println("Saindo...")
				default:
					fmt.Println("Comando inválido")
			}
		}
		fmt.Println("Digite o endereço do servidor:")
		fmt.Scanln(&url)
		entrada = ""
	}
}