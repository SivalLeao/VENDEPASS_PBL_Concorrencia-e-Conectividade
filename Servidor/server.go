package main

import (
	"encoding/json"
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
func cabecalho() {
	lipar_terminal()
	endereco, porta := endereco_local()
	fmt.Println("=-=-=-=-=-=-==-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println("|  Servidor funcionando no endereço:\033[32m", endereco+":"+porta + "  \033[0m|")
	fmt.Print("=-=-=-=-=-=-==-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-\n\n")
}

//Função para obter o endereço IP local
func endereco_local() (string, string){
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println("Erro ao obter endereco local:", err)
		return "", ""
	}
	defer conn.Close()

	endr := strings.Split(conn.LocalAddr().String(), ":")[0]

	//endereco := endereco_local() //Obtendo o endereço IP local
	porta := "8088"

	return endr, porta
}

func enviar(cliente net.Conn, dado []byte) error{
	_, erro := cliente.Write(dado)
	if erro != nil {
		return erro
	}
	return nil
}

func receber(cliente net.Conn) ([]byte, error){
	buffer := make([]byte, 1024)
	tam_bytes, erro := cliente.Read(buffer)
	if erro != nil {
		return nil, erro
	}
	return buffer[:tam_bytes], nil
}

//Função para enviar mensagens
func enviar_mensagem(cliente net.Conn, mensagem string) {
	erro := enviar(cliente, []byte(mensagem))
	if erro != nil {
		fmt.Println("Erro ao enviar mensagem:", erro)
		return
	}
}

//Função para receber mensagens
func receber_mensagem(cliente net.Conn) string {
	buffer, erro := receber(cliente)
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
func enviar_dados[Tipo any](cliente net.Conn, dados Tipo) error {
	jsonData, erro := serializar_dados(dados)
	if erro != nil {
		return erro
	}
	erro = enviar(cliente, jsonData)
	if erro != nil {
		return erro
	}
	return nil
}

//Função para receber dados de um tipo desconhecido (como um slice ou um map)
func receber_dados[Tipo any](cliente net.Conn) (Tipo, error) {
	buffer, erro := receber(cliente)
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

func exib_mens_receb(mens_receb string, ip string, porta string){
	//exibindo mensagem recebida
	fmt.Println("Mensagem recebida!")
	fmt.Println("Cliente:\033[34m", ip, "\033[0m:\033[34m", porta + "\033[0m")
	fmt.Println("\n\033[33m",mens_receb +"\033[0m\n")
	return
}

//Função para manipular a conexão com o cliente
func manipularConexao(cliente net.Conn, id *int, cliente_id map[string]int, rotas map[string]int, cliente_rotas map[int][]string) {
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
	
	cabecalho()
	fmt.Println("Conexão estabelecida com o cliente!")
	fmt.Println("Ip:\033[34m", ip, "\033[0mPorta:\033[34m", porta + "\033[0m")

	var erro error //Variável para armazenar erros
	var mens_env string //Variável para armazenar a mensagem a ser enviada
	var mens_receb string //Variável para armazenar a mensagem recebida
	var comando []string //Variável para armazenar o comando recebido a partir do particionamento da mensagem recebida
	var id_receb int //Variável para armazenar o IP recebido a partir do particionamento da mensagem recebida
	var operacao string //Variável para armazenar a operação recebida a partir do particionamento da mensagem recebida

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

		cabecalho()
		exib_mens_receb(mens_receb, ip, porta)

		//Particionando a mensagem recebida
		comando = strings.Split(mens_receb, ":") //Particiona a mensagem recebida
		if len(comando) < 2 { //Se a mensagem não for particionada corretamente, exibe uma mensagem de erro
			fmt.Println("Mensagem inválida recebida")
			continue
		}
		id_receb, _= strconv.Atoi(comando[0]) //O ID é a primeira parte da mensagem recebida
		operacao = comando[1] //A operação é a segunda parte da mensagem recebida

		//Verificando a operação
		switch operacao {
			case "1": //Comprar passagem
				erro = enviar_dados(cliente, rotas) //Envia as rotas disponíveis
				if erro != nil {
					fmt.Println("Erro ao enviar rotas disponíveis:", erro)
					continue
				}
				mens_receb = receber_mensagem(cliente) //Recebe a rota escolhida pelo cliente
				comando = strings.Split(mens_receb, ":") //Particiona a mensagem recebida
				if len(comando) < 2 { //Se a mensagem não for particionada corretamente, exibe uma mensagem de erro
					fmt.Println("Mensagem inválida recebida")
					continue
				}
				id_receb, _= strconv.Atoi(comando[0]) //O ID é a primeira parte da mensagem recebida
				operacao = comando[1] //A operação é a segunda parte da mensagem recebida
				if operacao != "3" {
					pertence, existe := rotas[operacao] //Verifica se a rota existe
					if (!existe || pertence != 0) { //Se a rota não existe ou já foi comprada, exibe uma mensagem
						fmt.Println("Rota inválida!")
						fmt.Println("mensagem recebida:", mens_receb)
						fmt.Println("Comando[1]", comando[1])
						fmt.Println("Comando[0]", comando[0])
						fmt.Println("operacao", operacao)
						mens_env = "Rota inválida!"
						enviar_mensagem(cliente, mens_env)
						continue
					}
					rotas[operacao] = id_receb //A rota é comprada pelo cliente
					cliente_rotas[id_receb] = append(cliente_rotas[id_receb], operacao) //A rota é adicionada à lista de rotas compradas pelo cliente
					mens_env = "ok"
					enviar_mensagem(cliente, mens_env)
					fmt.Println("Operação concluída com sucesso!")
				}
				continue

			case "2": //Consultar passagem
				rotas_compradas, existe := cliente_rotas[id_receb] //Verifica se o cliente já comprou alguma passagem
				if (!existe || len(rotas_compradas) < 1) { //Se o cliente não comprou nenhuma passagem, exibe uma mensagem
					fmt.Println("Cliente", id_receb, "ainda não comprou nenhuma passagem!")
					mens_env = "Sem passagens compradas"
					enviar_mensagem(cliente, mens_env)
					continue
				} else { //Se o cliente comprou...
					mens_env = "ok"
					enviar_mensagem(cliente, mens_env)
					erro = enviar_dados(cliente, rotas_compradas) //Envia as rotas compradas pelo cliente
					if erro != nil {
						fmt.Println("Erro ao enviar rotas compradas:", erro)
						continue
					}
					mens_receb = receber_mensagem(cliente) //Recebe resposta do que fazer a seguir (Rota para cancelar ou 3 para voltar)
					comando = strings.Split(mens_receb, ":") //Particiona a mensagem recebida
					if len(comando) < 2 { //Se a mensagem não for particionada corretamente, exibe uma mensagem de erro
						fmt.Println("Mensagem inválida recebida")
						continue
					}
					id_receb, _= strconv.Atoi(comando[0]) //O ID é a primeira parte da mensagem recebida
					operacao = comando[1] //A operação é a segunda parte da mensagem recebida
					if operacao != "3" { //Se a operação não for 3, o cliente deseja cancelar uma rota
						pertence, existe := rotas[operacao] //Verifica se a rota existe
						if (!existe || pertence != id_receb) { //Se a rota não existe ou não pertence ao cliente, exibe uma mensagem
							fmt.Println("Rota inválida!")
							mens_env = "Rota inválida!"
							enviar_mensagem(cliente, mens_env)
							continue
						}
						rotas[operacao] = 0 //Libera a rota
						// Fazer algo pra remover a rota da lista de rotas compradas do cliente
						var rotas_compradas_atualizada []string
						for _, rota := range rotas_compradas {
							if rota != operacao {
								rotas_compradas_atualizada = append(rotas_compradas_atualizada, rota)
							}
						}
						cliente_rotas[id_receb] = rotas_compradas_atualizada
						//cliente_rotas[id_receb] = append(cliente_rotas[id_receb][:0], cliente_rotas[id_receb][1:]...) //Remove a rota do cliente
						mens_env = "ok"
						enviar_mensagem(cliente, mens_env)
						fmt.Println("Operação concluída com sucesso!")
					}
					continue
				}

			case "3": //Sair
				mens_env = "exit_ok"
				enviar_mensagem(cliente, mens_env)
				fmt.Println("Cliente", id_receb, "desconectado!")
				return

			default:
				fmt.Println("Operação inválida!")
				mens_env = "Operação inválida!"
				enviar_mensagem(cliente, mens_env)
				continue 
		}
	}
}

func main() {
	/*
		    *Criando o servidor
			* A função Listen cria servidores
	*/

	_, porta := endereco_local() //Obtendo o endereço IP local

	server, erro := net.Listen("tcp", ":"+porta)

	if erro != nil {
		fmt.Println("Erro ao criar o servidor:", erro)
		return
	}

	//fecha a porta
	defer server.Close()

	//endereco := endereco_local() //Obtendo o endereço IP local
	//porta := "8088"
	
	// se funcionar
	cabecalho() //Exibindo o endereço local para conexão

	//Variáveis do servidor
	var id *int = new(int) //Ponteiro de ID que será atualizado conforme clientes se conectarem
	*id = 1 //ID é inicializado com valor 1. Ou seja, não há nenhum cliente ID 0
	cliente_id := make(map[string]int) //"Dicionário" que armazena o ID de cada cliente conectado (Ex.: {"127.0.0.1": 1})

	//"Dicionário" das rotas disponíveis e quem as comprou (Ex.: {"Salvador": 1, "Feira de Santana": 3, "Xique-Xique": 4, "Aracaju": 2})
	rotas := map[string]int{
		"Salvador": 0, 
		"Feira de Santana": 0, 
		"Xique-Xique": 0, 
		"Aracaju": 0, 
		"Maceio": 0, 
		"Recife": 0, 
		"Fortaleza": 0}

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

		go manipularConexao(conexao, id, cliente_id, rotas, cliente_rotas) //Manipula a conexão em uma nova thread
	}
}
