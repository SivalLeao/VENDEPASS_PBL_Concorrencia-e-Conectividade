// Para executar com docker, docker-compose up --build

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

// Para uso local
var serv_A string = "http://localhost:8080" // URL do servidor A
var serv_B string = "http://localhost:8081" // URL do servidor B
var serv_C string = "http://localhost:8082" // URL do servidor C

// Para uso em laboratório
// var serv_A string = "http://172.16.103.14:8080" // URL do servidor A
// var serv_B string = "http://172.16.103.13:8080" // URL do servidor B
// var serv_C string = "http://172.16.103.2:8080" // URL do servidor C

// Estrutura de dados para o cliente
// O cliente possui um ID, um nome e uma lista de rotas que pertencem a ele
type Cliente struct{
	id int
	nome string
	rotas []string
}

// Estrutura de dados para o cadastro de cliente
// O cadastro de cliente possui um ID e um nome
type Cadastro_req struct{
	Id int `json:"id"`
	Nome string `json:"nome"`
}

// Estrutura de dados para a requisição de cliente
// A requisição de cliente possui um ID e uma rota
type Client_req struct{
	Id int `json:"id"`
	Rota string `json:"rota"`
}

// Informações de servidor locais
// As rotas pertencentes a esse, os servidores com que vai se comunicar e sua porta de comunicação
type Infos_locais struct {
	rotas map[string]int // Mapa de rotas locais
	servidores []string // Slice de servidores com que vai se comunicar
	porta string // Porta de comunicação
	clientes map[int]Cliente
}

// Função para definir as informações locais do servidor a partir de seleção do usuário
func define_info() Infos_locais{
	var qual_serv string
	fmt.Println("Servidor (A, B ou C):")
	fmt.Scanln(&qual_serv)
	for{
		switch qual_serv {
		case "A":
			serv_local := Infos_locais{
				rotas: map[string]int{
						"Sao-Paulo":        0,
						"Rio-de-Janeiro":   0,
						"Belo-Horizonte":   0,
						"Porto-Alegre":     0,
						"Cachoeira":        0,
				},
				servidores: []string{
					serv_B,
					serv_C,
				},
				porta: ":8080",
				clientes: make(map[int]Cliente),
			}
			return serv_local
		case "B":
			serv_local := Infos_locais{
				rotas: map[string]int{
						"Serrinha":		 	0,
						"Salvador":         0,
						"Feira-de-Santana": 0,
						"Xique-Xique":      0,
						"Aracaju":          0,
				},
				servidores: []string{
					serv_A,
					serv_C,
				},
				porta: ":8081",
				// porta: "8080"
				clientes: make(map[int]Cliente),
			}
			return serv_local
		case "C":
			serv_local := Infos_locais{
				rotas: map[string]int{
						"Maceio":           0,
						"Recife":           0,
						"Fortaleza":        0,
						"Acre":             0,
						"Manaus":           0,
				},
				servidores: []string{
					serv_A,
					serv_B,
				},
				porta: ":8082",
				// porta: "8080"
				clientes: make(map[int]Cliente),
			}
			return serv_local
		default: // Caso a entrada seja inválida, solicitar novamente
			fmt.Println("Entrada inválida")
			fmt.Println("Servidor (A, B ou C):")
			fmt.Scanln(&qual_serv)
			continue
		}
	}
}

// Função para reunir todas as rotas locais e externas
func reunir_rotas(serv_local *Infos_locais) map[string]int{
	// Cria um mapa para concatenar as rotas locais e externas (que serão recebidas)
	rotas_combinadas := make(map[string]int)

	// Adiciona as rotas locais ao mapa combinado
	for k, v := range serv_local.rotas {
		rotas_combinadas[k] = v
	}

	// Itera sobre os servidores externos para buscar suas rotas
	for _, server := range serv_local.servidores {
		req, err := http.NewRequest("GET", server+"/rotas", nil) // Cria uma requisição GET para o servidor
		if err != nil {
			fmt.Printf("Failed to create request to server %s: %v\n", server, err)
			continue
		}
		req.Header.Set("X-Source", "servidor") // Adiciona o cabeçalho X-Source para identificar que é uma requisição de servidor

		client := &http.Client{} // Cria um cliente HTTP
		resp, err := client.Do(req) // Envia a requisição
		if err != nil {
			fmt.Printf("Failed to fetch rotas from server %s: %v\n", server, err)
			continue
		}
		defer resp.Body.Close()

		var rotas_externas map[string]int // Cria um mapa para armazenar as rotas recebidas
		if err := json.NewDecoder(resp.Body).Decode(&rotas_externas); err != nil { // Decodifica o JSON recebido para o mapa de rotas externas
			fmt.Printf("Failed to decode rotas from server %s: %v\n", server, err)
			continue
		}

		// Adiciona as rotas externas ao mapa combinado
		for k, v := range rotas_externas {
			rotas_combinadas[k] = v
		}
	}

	return rotas_combinadas
}

// Função para atribuir rota a um cliente
func atribuir_rota(serv_local *Infos_locais, cliente_req Client_req) {
	cliente := serv_local.clientes[cliente_req.Id] // Busca o cliente pelo ID
	cliente.rotas = append(cliente.rotas, cliente_req.Rota) // Adiciona a rota ao cliente
	serv_local.clientes[cliente_req.Id] = cliente // Atualiza o cliente no mapa de clientes

	_, existe := serv_local.rotas[cliente_req.Rota] // Verifica se a rota pertence a este servidor
	if existe{ // Caso a rota pertença a este servidor, atualiza o mapa de rotas
		serv_local.rotas[cliente_req.Rota] = cliente_req.Id // Atualiza o mapa de rotas
	}
}

// Função para remover uma rota de um cliente
func cancelar_rota(serv_local *Infos_locais, cliente_req Client_req) {
	var rotas_atualizada []string // Cria um slice para armazenar as rotas atualizadas
	cliente := serv_local.clientes[cliente_req.Id] // Busca o cliente pelo ID
	for _, rota := range cliente.rotas { // Itera sobre as rotas do cliente
		if rota != cliente_req.Rota { // Caso a rota seja diferente da rota a ser removida
			rotas_atualizada = append(rotas_atualizada, rota) // Adiciona a rota ao slice de rotas atualizadas
		}
	}
	cliente.rotas = rotas_atualizada // Atualiza as rotas do cliente
	serv_local.clientes[cliente_req.Id] = cliente // Atualiza o cliente no mapa de clientes

	_, existe := serv_local.rotas[cliente_req.Rota] // Verifica se a rota pertence a este servidor
	if existe{ // Caso a rota pertença a este servidor, atualiza o mapa de rotas
		serv_local.rotas[cliente_req.Rota] = 0 // Atualiza o mapa de rotas
	}
}

// Função para definir os métodos GET do servidor
func define_metodo_get(serv_local *Infos_locais, serv *gin.Engine){
	serv.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Método GET para retornar as rotas locais e externas
	serv.GET("/rotas", func(c *gin.Context) {
        // Verificar o cabeçalho X-Source para saber se veio de um servidor ou cliente
        fonte := c.GetHeader("X-Source")
        if fonte == "servidor" { // Caso seja de um servidor, responder com as rotas locais
            // Responde com as rotas locais
            c.JSON(http.StatusOK, serv_local.rotas)
            return
        }

        // Cria um mapa para concatenar as rotas locais e externas (que serão recebidas)
        rotas_combinadas := reunir_rotas(serv_local) // Reúne as rotas locais e externas

		fmt.Println("Quantidade de rotas:", len(rotas_combinadas))

        // Responde com as rotas locais e externas combinadas
        c.JSON(http.StatusOK, gin.H{"rotas": rotas_combinadas})
    })

	// Método GET para retornar as rotas pertencentes ao cliente
	serv.GET("/rotas_cliente", func(c *gin.Context) {
		cliente_id := c.Query("id") // Obtém o ID do cliente a partir dos parâmetros de consulta

		if cliente_id == "" { // Verifica se o ID do cliente foi fornecido
			c.JSON(http.StatusBadRequest, gin.H{"error": "id_do_cliente_nao_fornecido"})
			return
		}
	
		cliente_id_int, err := strconv.Atoi(cliente_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id_do_cliente_invalido"})
			return
		}
		cliente, existe := serv_local.clientes[cliente_id_int] // Busca o cliente pelo ID
		if !existe { // Caso o cliente não exista, responde com um erro
			c.JSON(http.StatusBadRequest, gin.H{"error": "cliente_nao_encontrado"})
			return
		}

		if (len(cliente.rotas) == 0) { // Caso o cliente não tenha rotas, responde com um erro
			c.JSON(http.StatusBadRequest, gin.H{"error": "cliente_sem_rotas"})
			return
		}


		// Responde com as rotas do cliente
		c.JSON(http.StatusOK, gin.H{"rotas": cliente.rotas})
	})
}

// Função para definir os métodos POST do servidor
func define_metodo_post(serv_local *Infos_locais, serv *gin.Engine, id_cont *int){
	// MÉTODO PARA TESTES
	serv.POST("/broadcast", func(c *gin.Context) {
		var json_dado map[string]interface{}
		if err := c.ShouldBindJSON(&json_dado); err != nil {
			fmt.Println("Failed to bind JSON (/broadcast):", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		json_valor, err := json.Marshal(json_dado)
		if err != nil {
			fmt.Println("Failed to serialize JSON (/broadcast):", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for _, server := range serv_local.servidores {
			go func(server string) {
				resp, err := http.Post(server+"/receive", "application/json", bytes.NewBuffer(json_valor))
				if err != nil {
					fmt.Printf("Failed to send to server %s: %v\n", server, err)
					return
				}
				defer resp.Body.Close()
			}(server)
		}

		c.JSON(http.StatusOK, gin.H{"status": "broadcasted"})
	})

	// MÉTODO PARA TESTES
	serv.POST("/receive", func(c *gin.Context) {
		var json map[string]interface{}
		if err := c.ShouldBindJSON(&json); err != nil {
			fmt.Println("Failed to bind JSON (/receive):", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Received message: %v\n", json)
		c.JSON(http.StatusOK, gin.H{"status": "received"})
	})

	// Método POST para cadastro de um cliente nos servidores
	serv.POST("/cadastro", func(c *gin.Context) {
		var cadastro Cadastro_req // Cria uma variável para armazenar o cadastro do cliente
		if err := c.ShouldBindJSON(&cadastro); err != nil { // Faz o bind do JSON recebido para a variável de cadastro
			fmt.Println("Erro ao fazer o bind JSON (/cadastro):", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fonte := c.GetHeader("X-Source") // Verifica o cabeçalho X-Source para saber se veio de um servidor ou cliente
		if fonte == "servidor"{ // Caso seja de um servidor, realizar apenas cadastro próprio do cliente
			serv_local.clientes[cadastro.Id] = Cliente{
				id: cadastro.Id,
				nome: cadastro.Nome,
				rotas: []string{},
			}
			*id_cont = cadastro.Id + 1
			c.JSON(http.StatusOK, gin.H{"status": "cadastrado"})
			return
		}

		//Verificando se o cliente já está cadastrado
		for _, cliente := range serv_local.clientes{
			if cliente.nome == cadastro.Nome{ // Caso o cliente já esteja cadastrado, responde com o ID do cliente
				c.JSON(http.StatusOK, gin.H{"status": "logado", "id": cliente.id}) // Responde com o ID do cliente já cadastrado
				return
			}
		}

		//Cadastrando o cliente localmente
		serv_local.clientes[*id_cont] = Cliente{
			id: *id_cont,
			nome: cadastro.Nome,
			rotas: []string{},
		}

		// Enviando cadastro de cliente aos outros servidores
		cadastro = Cadastro_req{ // Monta a estrutuda de dados para enviar os servidores
			Id: *id_cont,
			Nome: cadastro.Nome,
		}
		json_valor, err := json.Marshal(cadastro) // Serializa o JSON para enviar aos servidores
		if err != nil {
			fmt.Println("Erro ao serializar o JSON (/cadastro):", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		for _, server := range serv_local.servidores { // Loop para acessar outros servidores
			go func(server string) {
				req, err := http.NewRequest("POST", server+"/cadastro", bytes.NewBuffer(json_valor)) // Cria uma requisição POST para o servidor
				if err != nil {
					fmt.Printf("Failed to create request to server %s: %v\n", server, err)
					return
				}
				req.Header.Set("Content-Type", "application/json") // Adiciona o cabeçalho Content-Type para identificar que é um JSON
				req.Header.Set("X-Source", "servidor") // Adiciona o cabeçalho X-Source para identificar que é uma requisição de servidor

				client := &http.Client{} // Cria um cliente HTTP
				resp, err := client.Do(req) // Envia a requisição
				if err != nil {
					fmt.Printf("Failed to send to server %s: %v\n", server, err)
					return
				}
				defer resp.Body.Close()
			}(server)
		}
		c.JSON(http.StatusOK, gin.H{"status": "cadastrado", "id": *id_cont}) // Responde com o status de cadastrado e o ID do cliente
		*id_cont++ // Incrementa o contador de ID
	})
}

// Função para definir os métodos PATCH do servidor
func define_metodo_patch(serv_local *Infos_locais, serv *gin.Engine){
	// Método PATCH para adicionar uma rota ao cliente
	serv.PATCH("/comprar_rota", func(c *gin.Context){
		var cliente_req Client_req // Cria uma variável para armazenar a requisição do cliente
		if err := c.ShouldBindJSON(&cliente_req); err != nil { // Faz o bind do JSON recebido para a variável de requisição do cliente
			fmt.Println("Erro ao fazer o bind JSON (/comprar_rota):", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fonte := c.GetHeader("X-Source") // Verifica o cabeçalho X-Source para saber se veio de um servidor ou cliente
		if fonte == "servidor"{ // Caso seja de um servidor, realizar apenas a adição da rota ao cliente
			atribuir_rota(serv_local, cliente_req) // Atribui a rota ao cliente
			c.JSON(http.StatusOK, gin.H{"status": "rota_adicionada"}) // Responde com o status de rota adicionada
			return
		}

		rotas_combinadas := reunir_rotas(serv_local) // Reúne as rotas locais e externas
		pertence, existe := rotas_combinadas[cliente_req.Rota] // Verifica se a rota pertence a algum cliente
		if (!existe || pertence != 0) { // Caso a rota não exista ou já pertença ao próprio cliente ou a algum outro
			c.JSON(http.StatusBadRequest, gin.H{"status": "rota_indiponivel"}) // Responde com o status de rota indiponível
			return
		}
		// Caso a rota esteja disponível...
		atribuir_rota(serv_local, cliente_req) // Atribui a rota ao cliente

		// Enviando a requisição de compra de rota aos outros servidores
		json_valor, err := json.Marshal(cliente_req) // Serializa o JSON para enviar aos servidores
		if err != nil {
			fmt.Println("Erro ao serializar o JSON (/comprar_rota):", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		for _, server := range serv_local.servidores { // Loop para acessar outros servidores
			go func(server string) {
				req, err := http.NewRequest("PATCH", server+"/comprar_rota", bytes.NewBuffer(json_valor)) // Cria uma requisição PATCH para o servidor
				if err != nil {
					fmt.Printf("Failed to create request to server %s: %v\n", server, err)
					return
				}
				req.Header.Set("Content-Type", "application/json") // Adiciona o cabeçalho Content-Type para identificar que é um JSON
				req.Header.Set("X-Source", "servidor") // Adiciona o cabeçalho X-Source para identificar que é uma requisição de servidor

				client := &http.Client{} // Cria um cliente HTTP
				resp, err := client.Do(req) // Envia a requisição
				if err != nil {
					fmt.Printf("Failed to send to server %s: %v\n", server, err)
					return
				}
				defer resp.Body.Close()
			}(server)
		}
		c.JSON(http.StatusOK, gin.H{"status": "rota_adicionada"}) // Responde com o status de rota adicionada
	})

	// Método PATCH para cancelar uma rota de um cliente
	serv.PATCH("/cancelar_rota", func(c *gin.Context){
		var cliente_req Client_req // Cria uma variável para armazenar a requisição do cliente
		if err := c.ShouldBindJSON(&cliente_req); err != nil { // Faz o bind do JSON recebido para a variável de requisição do cliente
			fmt.Println("Erro ao fazer o bind JSON (/cancelar_rota):", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fonte := c.GetHeader("X-Source") // Verifica o cabeçalho X-Source para saber se veio de um servidor ou cliente
		if fonte == "servidor"{ // Caso seja de um servidor, realizar apenas o cancelamento da rota do cliente
			cancelar_rota(serv_local, cliente_req) // Cancela a rota do cliente
			c.JSON(http.StatusOK, gin.H{"status": "rota_cancelada"}) // Responde com o status de rota cancelada
			return
		}

		cliente, existe := serv_local.clientes[cliente_req.Id] // Busca o cliente pelo ID
		if !existe { // Caso o cliente não exista, responde com um erro
			c.JSON(http.StatusBadRequest, gin.H{"error": "cliente_nao_encontrado"})
			return
		}

		// Verifica se a rota pertence ao cliente
		indice := -1
		for i, rota := range cliente.rotas {
			if rota == cliente_req.Rota {
				indice = i
				break
			}
		}
		if indice == -1 { // Caso a rota não pertença ao cliente, responde com um erro
			c.JSON(http.StatusBadRequest, gin.H{"status": "rota_nao_encontrada"})
			return
		}

		// Caso a rota pertença ao cliente...
		cancelar_rota(serv_local, cliente_req) // Cancela a rota do cliente

		// Enviando a requisição de cancelamento de rota aos outros servidores
		json_valor, err := json.Marshal(cliente_req) // Serializa o JSON para enviar aos servidores
		if err != nil {
			fmt.Println("Erro ao serializar o JSON (/cancelar_rota):", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		for _, server := range serv_local.servidores { // Loop para acessar outros servidores
			go func(server string) {
				req, err := http.NewRequest("PATCH", server+"/cancelar_rota", bytes.NewBuffer(json_valor)) // Cria uma requisição PATCH para o servidor
				if err != nil {
					fmt.Printf("Failed to create request to server %s: %v\n", server, err)
					return
				}
				req.Header.Set("Content-Type", "application/json") // Adiciona o cabeçalho Content-Type para identificar que é um JSON
				req.Header.Set("X-Source", "servidor") // Adiciona o cabeçalho X-Source para identificar que é uma requisição de servidor

				client := &http.Client{} // Cria um cliente HTTP
				resp, err := client.Do(req) // Envia a requisição
				if err != nil {
					fmt.Printf("Failed to send to server %s: %v\n", server, err)
					return
				}
				defer resp.Body.Close()
			}(server)
		}
		c.JSON(http.StatusOK, gin.H{"status": "rota_cancelada"}) // Responde com o status de rota cancelada
	})
}

// Função para definir o servidor com os métodos POST, GET
func define_servidor(serv_local *Infos_locais, id_cont *int) *gin.Engine{
	r := gin.Default()

	// Configuração do middleware CORS
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:8088"}, // Permitir a origem do frontend
        AllowMethods:     []string{"GET", "POST", "PATCH", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "X-Source"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

	// Define os métodos GET
	define_metodo_get(serv_local, r)

	// Define os métodos POST
	define_metodo_post(serv_local, r, id_cont)

	// Define os métodos PATCH
	define_metodo_patch(serv_local, r)

	return r
}

func main() {
	var int_cont int = 1
	serv_local := define_info() // Define as informações locais do servidor
	servidor := define_servidor(&serv_local, &int_cont) // Define os métodos POST, GET e o servidor
	servidor.Run(serv_local.porta) // Executa o servidor
}