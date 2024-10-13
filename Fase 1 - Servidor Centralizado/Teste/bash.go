package main

import (
	"fmt"
	"os/exec"
	"sync"
)

func main() {
	const numClients = 5 // Número de clientes que você deseja executar
	var wg sync.WaitGroup

	cmd := exec.Command("cd Cliente") // Executa o client.go
	for i := 1; i <= numClients; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cmd = exec.Command("go", "run", "Cliente/clientTest.go", fmt.Sprint(id)) // Executa o client.go
			output, err := cmd.CombinedOutput()                              // Captura a saída e erro
			if err != nil {
				fmt.Printf("Erro ao executar cliente %d: %v\n", id, err)
			}
			fmt.Printf("Saída do cliente %d: %s\n", id, output)
		}(i)
	}

	wg.Wait() // Espera todos os goroutines terminarem
}
