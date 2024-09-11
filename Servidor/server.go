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
func cabecalho(endereco string, porta string) {
	lipar_terminal()
	fmt.Println("=-=-=-=-=-=-==-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println("|  Servidor funcionando no endereço:\033[32m", endereco+":"+porta + "  \033[0m|")
	fmt.Println("=-=-=-=-=-=-==-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-\n")
}

//Função para obter o endereço IP local
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

//Função para enviar mensagens
func enviar_mensagem(cliente net.Conn, mensagem string) {
	_, erro := cliente.Write([]byte(mensagem))
	if erro != nil {
		fmt.Println("Erro ao enviar mensagem:", erro)
		return
	}
}

//Função para receber mensagens
func receber_mensagem(cliente net.Conn) string {
	buffer := make([]byte, 1024)
	tam_bytes, erro := cliente.Read(buffer)
	if erro != nil {
		fmt.Println("Erro ao receber mensagem:", erro)
		return ""
	}

	mensagem := string(buffer[:tam_bytes])
	return mensagem
}

//Função para manipular a conexão com o cliente
func manipularConexao(cliente net.Conn, ip_server string, porta_server string, id *int, cliente_id map[string]int, rotas map[string]int, cliente_rotas map[int][]string) {
	//fechar conexao no fim da operacao
	defer cliente.Close()

	//Manipulando dados [Ler/Escrever]

	//indetificando o usuario
	id_porta := cliente.RemoteAddr().String()
	indetificador := strings.Split(id_porta, ":") //Obtendo o IP e a porta do cliente
	ip := indetificador[0]
	porta := indetificador[1]

	//Verificando se o cliente já foi identificado
	id_antigo, existe := cliente_id[ip] //Armazena o ID do cliente caso ele já exista no dicionário
	
	cabecalho(ip_server, porta_server)
	fmt.Println("Conexão estabelecida com o cliente!")
	fmt.Println("Ip:\033[34m", ip, "\033[0mPorta:\033[34m", porta + "\033[0m")

	var mens_env string //Variável para armazenar a mensagem a ser enviada
	var mens_receb string //Variável para armazenar a mensagem recebida
	var comando []string //Variável para armazenar o comando recebido a partir do particionamento da mensagem recebida

	if existe { //Se o cliente já foi identificado, envia o ID antigo
		mens_env = strconv.Itoa(id_antigo)
	} else { //Se o cliente não foi identificado, envia o ID atual e futuramente incrementa o ID
		mens_env = strconv.Itoa(*id)
	}

	enviar_mensagem(cliente, mens_env) //Envia a mensagem de identificação (ID do cliente)

	mens_receb = receber_mensagem(cliente) //Recebe a mensagem de confirmação de identificação do cliente
	if (mens_receb != "ID_ok"){ //Se a mensagem recebida não for "ID_ok", houve um erro na identificação do cliente
		fmt.Println("Falha na identificação do cliente")
		return
	}

	if existe { //Se o cliente já foi identificado, não incrementa o ID
		fmt.Println("Cliente", ip, "ID:", id_antigo, "reconectado!")
	} else { //Se o cliente não foi identificado, incrementa o ID e armazena o ID do cliente no dicionário
		cliente_id[ip] = *id
		*id++
		fmt.Println("Cliente", ip, "ID:", cliente_id[ip], "registrado!")
	}

	fmt.Println("Cliente identificado com sucesso!")


	for{

		//Guardando a mensagem
		mens_receb = receber_mensagem(cliente) //Recebe a mensagem do cliente com seus comandos

		cabecalho(ip_server, porta_server)
		//exibindo mensagem recebida
		fmt.Println("Mensagem recebida!")
		fmt.Println("Cliente:\033[34m", ip, "\033[0m:\033[34m", porta + "\033[0m")
		fmt.Println("\n\033[33m",mens_receb +"\033[0m\n")

		//Início particionando a mensagem recebida
		comando = strings.Split(mens_receb, ":")
		if len(comando) != 2 {
			fmt.Println("Comando inválido recebido!")
			mens_env = "Invalido"
			enviar_mensagem(cliente, mens_env)
			continue
		}
		id_receb, _ := strconv.Atoi(comando[0])
		//Fim particionando a mensagem recebida

		pertence, existe := rotas[comando[1]] //Verifica se a rota existe e se ela já foi comprada

		if !existe { //Se a rota não existe, envia uma mensagem de erro
			fmt.Println("Rota não encontrada!")
			mens_env = "Rota não encontrada!"
		} else { //Se a rota existe, verifica se ela já foi comprada
			if pertence == 0 { //Se a rota está disponível, envia uma mensagem de sucesso e atualiza o dicionário de rotas
				fmt.Println("Rota", comando[1], "disponível!")
				mens_env = "Rota comprada!"
				rotas[comando[1]] = id_receb
			} else { //Se a rota já foi comprada, envia uma mensagem de erro
				fmt.Println("Rota", comando[1], "indisponível!")
				mens_env = "Rota indisponível!"
			}
		}

		//Tratando a mensagem resposta
		if (comando[1] == "exit") { //Se o comando for "exit", encerra a conexão com o cliente
			mens_env = "exit_ok" //Envia uma mensagem de confirmação de encerramento
		}

		enviar_mensagem(cliente, mens_env) //Envia a mensagem de resposta ao cliente

		if (comando[1] == "exit") { //Se o comando for "exit", encerra a conexão com o cliente
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

	endereco := endereco_local() //Obtendo o endereço IP local
	porta := "8088"
	
	// se funcionar
	cabecalho(endereco, porta) //Exibindo o endereço local para conexão

	//Variáveis do servidor
	var id *int = new(int) //Ponteiro de ID que será atualizado conforme clientes se conectarem
	*id = 1 //ID é inicializado com valor 1. Ou seja, não há nenhum cliente ID 0
	cliente_id := make(map[string]int) //"Dicionário" que armazena o ID de cada cliente conectado (Ex.: {"127.0.0.1": 1})

	//"Dicionário" das rotas disponíveis e quem as comprou (Ex.: {"Salvador": 1, "Feira de Santana": 3, "Xique-Xique": 4, "Aracaju": 2})
	rotas := map[string]int{"Salvador": 0, "Feira de Santana": 0, "Xique-Xique": 0, "Aracaju": 0}

	//"Dicionário" que armazena as rotas compradas por cada cliente 
	//(Ex.: {1: ["Salvador", "Feira de Santana"], 2: ["Aracaju"], 3: ["Feira de Santana"], 4: ["Xique-Xique"]})
	cliente_rotas := make(map[int][]string)

	//Loop infinito do servidor
	for {
		conexao, erro := server.Accept() //Aceita conexões

		if erro != nil {
			fmt.Println("Erro ao aceitar conexão:", erro)
			continue
		}

		go manipularConexao(conexao, endereco, porta, id, cliente_id, rotas, cliente_rotas) //Manipula a conexão em uma nova thread
	}
}
