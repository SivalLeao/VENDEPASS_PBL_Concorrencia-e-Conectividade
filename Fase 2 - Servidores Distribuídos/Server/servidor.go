// Para executar com docker, docker-compose up --build

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Informações de servidor locais
// As rotas pertencentes a esse, os servidores com que vai se comunicar e sua porta de comunicação
type Infos_locais struct {
	rotas map[string]int // Mapa de rotas locais
	servidores []string // Slice de servidores com que vai se comunicar
	porta string // Porta de comunicação
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
					"http://localhost:8081",
					"http://localhost:8082",
				},
				porta: ":8080",
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
					"http://localhost:8080",
					"http://localhost:8082",
				},
				porta: ":8081",
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
					"http://localhost:8080",
					"http://localhost:8081",
				},
				porta: ":8082",
			}
			return serv_local
		default:
			fmt.Println("Entrada inválida")
			fmt.Println("Servidor (A, B ou C):")
			fmt.Scanln(&qual_serv)
			continue
		}
	}
}

// Função para definir o servidor com os métodos POST, GET
func define_servidor(serv_local *Infos_locais) *gin.Engine{
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Método GET para retornar as rotas locais e externas
	r.GET("/rotas", func(c *gin.Context) {
        // Verificar o cabeçalho X-Source para saber se veio de um servidor ou cliente
        source := c.GetHeader("X-Source")
        if source == "server" { // Caso seja de um servidor, responder com as rotas locais
            // Responde com as rotas locais
            c.JSON(http.StatusOK, serv_local.rotas)
            return
        }

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
            req.Header.Set("X-Source", "server") // Adiciona o cabeçalho X-Source para identificar que é uma requisição de servidor

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

		fmt.Println("Quantidade de rotas:", len(rotas_combinadas))

        // Responde com as rotas locais e externas combinadas
        c.JSON(http.StatusOK, rotas_combinadas)
    })

	r.POST("/broadcast", func(c *gin.Context) {
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

	r.POST("/receive", func(c *gin.Context) {
		var json map[string]interface{}
		if err := c.ShouldBindJSON(&json); err != nil {
			fmt.Println("Failed to bind JSON (/receive):", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Received message: %v\n", json)
		c.JSON(http.StatusOK, gin.H{"status": "received"})
	})

	return r
}

func main() {
	serv_local := define_info() // Define as informações locais do servidor
	servidor := define_servidor(&serv_local) // Define os métodos POST, GET e o servidor
	servidor.Run(serv_local.porta) // Executa o servidor
}