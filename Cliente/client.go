package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

//Função para limpar o terminal
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

//Função para exibir o cabeçalho com o endereço do servidor para conexão
func cabecalho(endereco string) {
	lipar_terminal()
	fmt.Println("=-=-=-=-=-=-==-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println("|\033[32m             VENDEPASS: Venda de Passagens         	 \033[0m|")
	fmt.Println("|--------------------------------------------------------|")
	fmt.Println("|\033[34m           Conectado:", endereco + "                \033[0m|")
	fmt.Println("=-=-=-=-=-=-==-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-\n")
}

//Função para enviar mensagens
func enviar_mensagem(server net.Conn, mensagem string) {
	_, erro := server.Write([]byte(mensagem))
	if erro != nil {
		fmt.Println("Erro ao enviar mensagem:", erro)
		return
	}
}

//Função para receber mensagens
func receber_mensagem(server net.Conn) string {
	buffer := make([]byte, 1024)
	tam_bytes, erro := server.Read(buffer)
	if erro != nil {
		fmt.Println("Erro ao receber mensagem:", erro)
		return ""
	}

	return string(buffer[:tam_bytes])
}

//Função para manipular a conexão com o servidor
func manipularConexao(server net.Conn, endereco string) {
	//fechar a conexao no fim da operacao
	defer server.Close()

	mens_receb := receber_mensagem(server) // Recebe a mensagem de identificação do cliente
	_, erro := strconv.Atoi(mens_receb) // Testa se a mensagem recebida pode ser convertida em um inteiro

	if (erro != nil){ // Se a mensagem recebida for um inteiro, o cliente foi identificado
		fmt.Println("Erro ao identificar o cliente")
		return
	}

	id := mens_receb // Guarda o id que foi recebido e atribuído pelo servidor

	mens_env := "ID_ok" // Envia a mensagem de confirmação de identificação
	enviar_mensagem(server, mens_env) // Envia a mensagem de confirmação de identificação

	fmt.Println("Identificação concluída")

	var scan string // Variável para armazenar os comandos digitados pelo usuário

	// Loop para enviar e receber mensagens
	for {

		// Enviar mensagem
		fmt.Print("Digite a mensagem a ser enviada: ")
		fmt.Scanln(&scan) // Recebe o comando que se deseja realizar
		mens_env = id + ":" + scan // Concatena o id do cliente com o comando que se deseja realizar. A mensagem a ser enviada ao servidor
		enviar_mensagem(server, mens_env) // Envia a mensagem ao servidor
		
		fmt.Println("\nMensagem enviada\n")

		// Recebendo mensagem do server
		// Guardando a mensagem
		mens_receb = receber_mensagem(server)

		// Exibindo a mensagem
		fmt.Println("Mensagem recebida:", mens_receb +"\n")
		fmt.Println("=============================================\n")

		if (scan == "exit" && mens_receb == "exit_ok") { // Se o comando digitado for "exit" e a mensagem recebida for "exit_ok", encerra a conexão
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
	fmt.Scanln(&endereco_alvo) // Recebe o endereço do servidor a que se deseja conectar

	conexao, erro := net.Dial("tcp", endereco_alvo) // Conecta-se ao servidor

	if erro != nil {
		fmt.Println("Erro ao se conectar ao servidor:", erro)
		return
	}
	//fechando a conexao
	defer conexao.Close()

	cabecalho(endereco_alvo)
	//fmt.Println("Conectado ao servidor no endereço", endereco_alvo)

	manipularConexao(conexao, endereco_alvo) // Manipula a conexão com o servidor

}
