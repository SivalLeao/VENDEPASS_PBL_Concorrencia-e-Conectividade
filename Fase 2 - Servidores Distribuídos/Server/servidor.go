// Para executar com docker, docker-compose up --build

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var rotas_locais = map[string]int{
	"Sao-Paulo":        0,
	"Rio-de-Janeiro":   0,
	"Belo-Horizonte":   0,
	"Porto-Alegre":     0,
	"Cachoeira":        0,
	}

// rotas_locais := map[string]int{
// 	"Serrinha":		 	0,
// 	"Salvador":         0,
// 	"Feira-de-Santana": 0,
// 	"Xique-Xique":      0,
// 	"Aracaju":          0,
//	}

// rotas_locais := map[string]int{
// 	"Maceio":           0,
// 	"Recife":           0,
// 	"Fortaleza":        0,
// 	"Acre":             0,
// 	"Manaus":           0,
//	}

var servidores = []string{
	"http://localhost:8081",
	"http://localhost:8082",
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/rotas", func(c *gin.Context) {
        // Verificar o cabeçalho X-Source
        source := c.GetHeader("X-Source")
        if source == "server" {
            // Responder com as rotas locais
            c.JSON(http.StatusOK, rotas_locais)
            return
        }

        // Create a map to hold the combined rotas
        rotas_combinadas := make(map[string]int)

        // Add local rotas to the combined rotas
        for k, v := range rotas_locais {
            rotas_combinadas[k] = v
        }

        // Fetch rotas from other servidores
        for _, server := range servidores {
            req, err := http.NewRequest("GET", server+"/rotas", nil)
            if err != nil {
                fmt.Printf("Failed to create request to server %s: %v\n", server, err)
                continue
            }
            req.Header.Set("X-Source", "server")

            client := &http.Client{}
            resp, err := client.Do(req)
            if err != nil {
                fmt.Printf("Failed to fetch rotas from server %s: %v\n", server, err)
                continue
            }
            defer resp.Body.Close()

            var rotas_externa map[string]int
            if err := json.NewDecoder(resp.Body).Decode(&rotas_externa); err != nil {
                fmt.Printf("Failed to decode rotas from server %s: %v\n", server, err)
                continue
            }

            // Add server rotas to the combined rotas
            for k, v := range rotas_externa {
                rotas_combinadas[k] = v
            }
        }

        // Return the combined rotas
        c.JSON(http.StatusOK, rotas_combinadas)
    })

	r.POST("/broadcast", func(c *gin.Context) {
		var jsonData map[string]interface{}
		if err := c.ShouldBindJSON(&jsonData); err != nil {
			fmt.Println("Failed to bind JSON (/broadcast):", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		jsonValue, err := json.Marshal(jsonData)
		if err != nil {
			fmt.Println("Failed to serialize JSON (/broadcast):", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for _, server := range servidores {
			go func(server string) {
				resp, err := http.Post(server+"/receive", "application/json", bytes.NewBuffer(jsonValue))
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

	r.Run(":8080")
}