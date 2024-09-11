package main

import (
	"encoding/json"
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
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println("|\033[32m             VENDEPASS: Venda de Passagens         	 \033[0m|")
	fmt.Println("|--------------------------------------------------------|")
	fmt.Println("|\033[34m           Conectado:", endereco + "                \033[0m|")
	fmt.Print("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-\n\n")
}

func enviar(server net.Conn, dado []byte) error{
	_, erro := server.Write(dado)
	if erro != nil {
		return erro
	}
	return nil
}

func receber(server net.Conn) ([]byte, error){
	buffer := make([]byte, 1024)
	tam_bytes, erro := server.Read(buffer)
	if erro != nil {
		return nil, erro
	}
	return buffer[:tam_bytes], nil
}

//Função para enviar mensagens
func enviar_mensagem(server net.Conn, mensagem string) {
	erro := enviar(server, []byte(mensagem))
	if erro != nil {
		fmt.Println("Erro ao enviar mensagem:", erro)
		return
	}
}

//Função para receber mensagens
func receber_mensagem(server net.Conn) string {
	buffer, erro := receber(server)
	if erro != nil {
		fmt.Println("Erro ao receber mensagem:", erro)
		return ""
	}

	return string(buffer)
}

//Função para serializar dados
func serializar_dados[Tipo any](dados Tipo) ([]byte, error){
	jsonData, erro := json.Marshal(dados)
	if erro != nil {
		return nil, erro
	}
	return jsonData, nil
}

//Função para desserializar dados
func desserializar_dados[Tipo any](jsonData []byte) (Tipo, error){
	var dados Tipo
	erro := json.Unmarshal(jsonData, &dados)
	if erro != nil {
		return dados, erro
	}
	return dados, nil
}

//Função para enviar dados de um tipo desconhecido (como um slice ou um map)
func enviar_dados[Tipo any](server net.Conn, dados Tipo) error {
	jsonData, erro := serializar_dados(dados)
	if erro != nil {
		return erro
	}
	erro = enviar(server, jsonData)
	if erro != nil {
		return erro
	}
	return nil
}

//Função para receber dados de um tipo desconhecido (como um slice ou um map)
func receber_dados[Tipo any](server net.Conn) (Tipo, error) {
	buffer, erro := receber(server)
	var dados Tipo
	if erro != nil {
		return dados, erro
	}
	dados, erro = desserializar_dados[Tipo](buffer)
	if erro != nil {
		return dados, erro
	}
	return dados, nil
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
		//fmt.Print("Digite a mensagem a ser enviada: ")
		fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
		fmt.Println("                            MENU")
		fmt.Print("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-\n\n")

		fmt.Println("1 - Comprar passagem")
		fmt.Println("2 - Consultar passagem")
		fmt.Print("3 - Sair\n\n")
		fmt.Print("==========================================================\n\n")


		fmt.Scanln(&scan) // Recebe o comando que se deseja realizar
		
		switch scan {
			case "1":
				cabecalho(endereco)
				fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
				fmt.Println("                    Comprar passagem")
				fmt.Print("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-\n\n")

				fmt.Println("Rotas disponiveis:")
				// Solicitar rotas disponíveis do servidor
				rotas := map[string]int{"Salvador": 1, "Feira de Santana": 1, "Xique-Xique": 0, "Aracaju": 1, "Maceió": 0, "Recife": 0}
				
				var matriz[3][2]string
				for local, vaga := range rotas {
					for i := 0; i < 3; i++ {
						for j := 0; j < 2; j++ {
							if vaga == 1 {
								matriz[i][j] = local
							}
							if vaga == 0 {
								matriz[i][j] = "\033[34m"+local+"\033[0m"
							}
							
						}
					}
				}

				//var lista_rotas = [6]string{"\033[34mRota 1\033[0m", " Rota 2", "Rota 3", "Rota 4", "Rota 5", "Rota 6"} 

				fmt.Println("----------------------------------------------------------")
				fmt.Println(matriz[0][0], matriz[0][1])
				fmt.Println(matriz[1][0], matriz[1][1])
				fmt.Print("----------------------------------------------------------\n\n")

				fmt.Print("Digite a rota desejada: ")
				fmt.Scanln(&scan)
				mens_env = id + ":" + scan // Concatena o id do cliente com o comando que se deseja realizar. A mensagem a ser enviada ao servidor
				enviar_mensagem(server, mens_env) // Envia a mensagem ao servidor
		}
		
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
