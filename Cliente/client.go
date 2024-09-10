package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func lipar_terminal() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
		case "windows":
			cmd = exec.Command("cmd", "/c", "cls")
		default: //linux e mac
			cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	erro := cmd.Run()
	if erro != nil {
		fmt.Println("Erro ao limpar o terminal:", erro)
		return
	}
}

func cabecalho(endereco string) {
	lipar_terminal()
	fmt.Println("=-=-=-=-=-=-==-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println("|\033[32m             VENDEPASS: Venda de Passagens         	 \033[0m|")
	fmt.Println("|--------------------------------------------------------|")
	fmt.Println("|\033[34m           Conectado:", endereco + "                \033[0m|")
	fmt.Println("=-=-=-=-=-=-==-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-\n")
}

func manipularConexao(server net.Conn, endereco string) {
	//fechar a conexao no fim da operacao
	defer server.Close()

	buffer := make([]byte, 1024)
	tam_bytes, erro := server.Read(buffer)
	if erro != nil {
		fmt.Println("Erro ao receber mensagem:", erro)
		return
	}

	mens_receb := string(buffer[:tam_bytes])
	_, erro = strconv.Atoi(mens_receb) // Testa se a mensagem recebida pode ser convertida em um inteiro

	if (erro != nil){ // Se a mensagem recebida for um inteiro, o cliente foi identificado
		fmt.Println("Erro ao identificar o cliente")
		return
	}

	id := mens_receb

	mens_env := "ID_ok"
	_, erro = server.Write([]byte(mens_env))
	if erro != nil {
		fmt.Println("Erro ao enviar mensagem:", erro)
		return
	}

	fmt.Println("Identificação concluída")

	var scan string

	//Manipulando dados [Ler/Escrever]
	for {

		// Enviar mensagem
		fmt.Print("Digite a mensagem a ser enviada: ")
		fmt.Scanln(&scan)
		mens_env = id + ":" + scan
		_, erro = server.Write([]byte(mens_env))
		if erro != nil {
			fmt.Println("Erro ao enviar mensagem:", erro)
			return
		}
		
		fmt.Println("\nMensagem enviada\n")

		// Recebendo mensagem do server
		// Buffer de 1K
		buffer = make([]byte, 1024)
		//Tamanho da mensagem recebida
		tam_bytes, erro = server.Read(buffer)
		if erro != nil {
			fmt.Println("Erro ao receber mensagem:", erro)
			return
		}

		// Guardando a mensagem
		mens_receb = string(buffer[:tam_bytes])

		// Exibindo a mensagem
		fmt.Println("Mensagem recebida:", mens_receb +"\n")
		fmt.Println("=============================================\n")

		if (scan == "exit" && mens_receb == "exit_ok") {
			return
		}
	}
}

func main() {
	lipar_terminal()
	/*
		    * Acessando o servidor
			* A função Dial conecta-se a um servidor
	*/

	var endereco_alvo string
	fmt.Print("Digite o endereço alvo: ")
	fmt.Scanln(&endereco_alvo)

	conexao, erro := net.Dial("tcp", endereco_alvo)

	if erro != nil {
		fmt.Println("Erro ao se conectar ao servidor:", erro)
		return
	}
	//fechando a conexao
	defer conexao.Close()

	cabecalho(endereco_alvo)
	//fmt.Println("Conectado ao servidor no endereço", endereco_alvo)

	manipularConexao(conexao, endereco_alvo)

}
