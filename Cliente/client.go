package main

import (
	"fmt"
	"net"
)

func manipularConexao(server net.Conn) {
	//fechar a conexao no fim da operacao
	defer server.Close()

	//Manipulando dados [Ler/Escrever]

	// Enviar mensagem
	mens_envio := "hello world!!!"
	_, erro := server.Write([]byte(mens_envio))
	if erro != nil {
		fmt.Println("Erro ao enviar mensagem:", erro)
		return
	}
	
	fmt.Println("Mensagem enviada")

	// Recebendo mensagem do server
	// Buffer de 1K
	buffer := make([]byte, 1024)
	//Tamanho da mensagem recebida
	tam_bytes, erro := server.Read(buffer)
	if erro != nil {
		fmt.Println("Erro ao le os dados:", erro)
		return
	}

	// Guardando a mensagem
	mensagem := string(buffer[:tam_bytes])

	// Exibindo a mensagem
	fmt.Println("Mensagem recebida:", mensagem)
}

func main() {
	/*
		    * Acessando o servidor
			* A função Dial conecta-se a um servidor
	*/
	conectando, erro := net.Dial("tcp", "127.0.0.1:8088")

	if erro != nil {
		fmt.Println("Erro ao se conectar ao servidor:", erro)
		return
	}
	//fechando a conexao
	defer conectando.Close()

	fmt.Println("Conectado ao servidor na porta 8088")

	manipularConexao(conectando)

}
