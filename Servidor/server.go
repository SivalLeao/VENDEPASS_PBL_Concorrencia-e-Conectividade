package main

import (
	"fmt"
	"net"
	"strings"
)

func endereco_local() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println("Erro ao obter endereco local:", err)
		return ""
	}
	defer conn.Close()

	endr := strings.Split(conn.LocalAddr().String(), ":")[0]

	return endr
}

func manipularConexao(cliente net.Conn) {
	//fechar conexao no fim da operacao
	defer cliente.Close()

	//Manipulando dados [Ler/Escrever]

	//indetificando o usuario
	id_porta := cliente.RemoteAddr().String()
	indetificador := strings.Split(id_porta, ":")
	ip := indetificador[0]
	porta := indetificador[1]

	fmt.Println("Usuario Ip:", ip, "porta:", porta)

	mens_env := "1" //Teste de envio de ID ao cliente
	_, erro := cliente.Write([]byte(mens_env))
	if erro != nil {
		fmt.Println("Erro ao enviar mensagem:", erro)
		return
	}

	buffer := make([]byte, 1024)
	tam_bytes, erro := cliente.Read(buffer)
	if erro != nil {
		fmt.Println("Erro ao receber mensagem:", erro)
		return
	}
	mens_receb := string(buffer[:tam_bytes])
	if (mens_receb != "ID_ok"){
		fmt.Println("Falha na identificação do cliente")
		return
	}

	fmt.Println("Cliente identificado com sucesso!")


	for{
		//Lendo dados
		//Buffer de 1 KB
		buffer = make([]byte, 1024)
		//Tamanho da mensagem recebida
		tam_bytes, erro = cliente.Read(buffer)

		if erro != nil {
			fmt.Println("Erro ao receber mensagem:", erro)
			return
		}

		//Guardando a mensagem
		mens_receb = string(buffer[:tam_bytes])

		//exibindo mensagem recebida
		fmt.Println(mens_receb)

		//Tratando a mensagem resposta
		if (mens_receb == "exit") { 
			mens_env = "exit_ok"
		} else {
			mens_env = "OK!"
		}
		_, erro = cliente.Write([]byte(mens_env))

		if erro != nil {
			fmt.Println("Erro ao enviar a mensagem:", erro)
			return
		}
		fmt.Println("Mensagem enviada com sucesso!")
		if (mens_receb == "exit") {
			fmt.Println("Encerramento confirmado!")
			return
		}
	}

}

func main() {
	/*
		    *Criando o servidor
			* A função Listen cria servidores
	*/
	server, erro := net.Listen("tcp", ":8088")

	if erro != nil {
		fmt.Println("Erro ao criar o servidor:", erro)
		return
	}

	//fecha a porta
	defer server.Close()

	// endereco := endereco_local()
	// fmt.Println("Servidor funcionando no endereço", endereco)
	// se funcionar
	fmt.Println("Servidor funcionando na porta 8088...")

	//Loop infinito do servidor
	for {
		conexao, erro := server.Accept()

		if erro != nil {
			fmt.Println("Erro ao aceitar conexão:", erro)
			continue
		}
		go manipularConexao(conexao)
	}
}
