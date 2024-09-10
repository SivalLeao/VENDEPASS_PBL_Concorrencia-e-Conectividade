package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
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

func cabecalho(endereco string, porta string) {
	lipar_terminal()
	fmt.Println("=-=-=-=-=-=-==-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println("|  Servidor funcionando no endereço:\033[32m", endereco+":"+porta + "  \033[0m|")
	fmt.Println("=-=-=-=-=-=-==-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-\n")
}

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

func manipularConexao(cliente net.Conn, ip_server string, porta_server string, id *int, cliente_id map[string]int, rotas map[string]int) {
	//fechar conexao no fim da operacao
	defer cliente.Close()

	//Manipulando dados [Ler/Escrever]

	//indetificando o usuario
	id_porta := cliente.RemoteAddr().String()
	indetificador := strings.Split(id_porta, ":")
	ip := indetificador[0]
	porta := indetificador[1]

	//Verificando se o cliente já foi identificado
	id_antigo, existe := cliente_id[ip]
	
	cabecalho(ip_server, porta_server)
	fmt.Println("Conexão estabelecida com o cliente!")
	fmt.Println("Ip:\033[34m", ip, "\033[0mPorta:\033[34m", porta + "\033[0m")

	var mens_env string
	var comando []string

	if existe {
		mens_env = strconv.Itoa(id_antigo)
	} else {
		mens_env = strconv.Itoa(*id)
	}

	_, erro := cliente.Write([]byte(mens_env))
	if erro != nil {
		fmt.Println("Erro ao enviar mensagem:", erro)
		return
	}

	buffer := make([]byte, 1024)
	tam_bytes, erro := cliente.Read(buffer)
	if erro != nil {
		cabecalho(ip_server, porta_server)
		fmt.Println("Erro ao receber mensagem:", erro)
		return
	}
	mens_receb := string(buffer[:tam_bytes])
	if (mens_receb != "ID_ok"){
		fmt.Println("Falha na identificação do cliente")
		return
	}

	if existe {
		fmt.Println("Cliente", ip, "ID:", id_antigo, "reconectado!")
	} else {
		cliente_id[ip] = *id
		*id++
		fmt.Println("Cliente", ip, "ID:", cliente_id[ip], "registrado!")
	}

	fmt.Println("Cliente identificado com sucesso!")


	for{
		//Lendo dados
		//Buffer de 1 KB
		buffer = make([]byte, 1024)
		//Tamanho da mensagem recebida
		tam_bytes, erro = cliente.Read(buffer)

		if erro != nil {
			cabecalho(ip_server, porta_server)
			fmt.Println("Erro ao receber mensagem:", erro)
			return
		}

		//Guardando a mensagem
		mens_receb = string(buffer[:tam_bytes])

		cabecalho(ip_server, porta_server)
		//exibindo mensagem recebida
		fmt.Println("Mensagem recebida!")
		fmt.Println("Cliente:\033[34m", ip, "\033[0m:\033[34m", porta + "\033[0m")
		fmt.Println("\n\033[33m",mens_receb +"\033[0m\n")

		comando = strings.Split(mens_receb, ":")
		if len(comando) != 2 {
			fmt.Println("Comando inválido recebido!")
			mens_env = "Invalido"
			_, erro = cliente.Write([]byte(mens_env))

			if erro != nil {
				fmt.Println("Erro ao enviar a mensagem:", erro)
				return
			}
			fmt.Println("Resposta enviada com sucesso!")
			continue
		}
		id_receb, _ := strconv.Atoi(comando[0])

		pertence, existe := rotas[comando[1]]

		if !existe {
			fmt.Println("Rota não encontrada!")
			mens_env = "Rota não encontrada!"
		} else {
			if pertence == 0 {
				fmt.Println("Rota", comando[1], "disponível!")
				mens_env = "Rota comprada!"
				rotas[comando[1]] = id_receb
			} else {
				fmt.Println("Rota", comando[1], "indisponível!")
				mens_env = "Rota indisponível!"
			}
		}

		//Tratando a mensagem resposta
		if (comando[1] == "exit") { 
			mens_env = "exit_ok"
		}
		_, erro = cliente.Write([]byte(mens_env))

		if erro != nil {
			fmt.Println("Erro ao enviar a mensagem:", erro)
			return
		}
		fmt.Println("Resposta enviada com sucesso!")

		if (comando[1] == "exit") {
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

	endereco := endereco_local()
	porta := "8088"
	
	// se funcionar
	cabecalho(endereco, porta)

	var id *int = new(int)
	*id = 1
	cliente_id := make(map[string]int)
	rotas := map[string]int{"Salvador": 0, "Feira de Santana": 0, "Xique-Xique": 0, "Aracaju": 0}

	//Loop infinito do servidor
	for {
		conexao, erro := server.Accept()

		if erro != nil {
			fmt.Println("Erro ao aceitar conexão:", erro)
			continue
		}
		go manipularConexao(conexao, endereco, porta, id, cliente_id, rotas)
	}
}
