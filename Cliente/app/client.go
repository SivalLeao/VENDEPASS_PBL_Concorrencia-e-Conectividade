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
	"time"
)

// Função para limpar o terminal
func limpar_terminal() {
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

// Função para exibir o cabeçalho com o endereço do servidor para conexão
func cabecalho(endereco string) {
	limpar_terminal()
	
	tamanho := len(endereco)
	espacamento := ""
	if tamanho < 33 {
		espacamento = strings.Repeat(" ", 33-tamanho)
	}
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println("|\033[32m             VENDEPASS: Venda de Passagens         	 \033[0m|")
	fmt.Println("|--------------------------------------------------------|")
	fmt.Println("|\033[34m            Conectado:", endereco+espacamento+"\033[0m|")
	fmt.Print("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-\n\n")
}

func enviar(server net.Conn, dado []byte) error {
	_, erro := server.Write(dado)
	if erro != nil {
		return erro
	}
	return nil
}

func receber(server net.Conn) ([]byte, error) {
	buffer := make([]byte, 1024)
	tam_bytes, erro := server.Read(buffer)
	if erro != nil {
		return nil, erro
	}
	return buffer[:tam_bytes], nil
}

// Função para enviar mensagens
func enviar_mensagem(server net.Conn, mensagem string) {
	erro := enviar(server, []byte(mensagem))
	if erro != nil {
		limpar_terminal()
		
		if strings.Contains(erro.Error(), "Foi forçado o cancelamento de uma conexão existente pelo host remoto") {
			fmt.Println("\033[31m           Servidor com instabilidade...\033[0m")
			fmt.Print("Foi forçado o cancelamento de uma conexão \nexistente pelo host remoto\n")
			time.Sleep(5 * time.Second)

			main()
		}
		fmt.Println("--------------------------------------------")
		fmt.Println("Erro ao enviar mensagem:", erro)
		fmt.Println("--------------------------------------------")
		return
	}
}

// Função para receber mensagens
func receber_mensagem(server net.Conn) string {

	buffer, erro := receber(server)
	if erro != nil {
		fmt.Println("Erro ao receber mensagem:", erro)
		return ""
	}

	return string(buffer)
}

// Função para serializar dados
func serializar_dados[Tipo any](dados Tipo) ([]byte, error) {
	jsonData, erro := json.Marshal(dados)
	if erro != nil {
		return nil, erro
	}
	return jsonData, nil
}

// Função para desserializar dados
func desserializar_dados[Tipo any](jsonData []byte) (Tipo, error) {
	var dados Tipo
	erro := json.Unmarshal(jsonData, &dados)
	if erro != nil {
		return dados, erro
	}
	return dados, nil
}

// Função para enviar dados de um tipo desconhecido (como um slice ou um map)
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

// Função para receber dados de um tipo desconhecido (como um slice ou um map)
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

// Função para manipular a conexão com o servidor
func manipularConexao(server net.Conn, endereco string) {
	//fechar a conexao no fim da operacao
	defer server.Close()

	mens_receb := receber_mensagem(server) // Recebe a mensagem de identificação do cliente
	_, erro := strconv.Atoi(mens_receb)    // Testa se a mensagem recebida pode ser convertida em um inteiro

	if erro != nil { // Se a mensagem recebida for um inteiro, o cliente foi identificado
		fmt.Println("Erro ao identificar o cliente")
		return
	}

	id := mens_receb // Guarda o id que foi recebido e atribuído pelo servidor

	mens_env := "ID_ok"               // Envia a mensagem de confirmação de identificação
	enviar_mensagem(server, mens_env) // Envia a mensagem de confirmação de identificação

	fmt.Println(strings.TrimLeft("                 Identificação concluída", "\n\n"))

	var scan string                              // Variável para armazenar os comandos digitados pelo usuário
	var rotas_compradas []string                 // Lista de rotas compradas pelo cliente
	var rotas_disponiveis = make(map[string]int) // Mapa de rotas disponíveis

	// Loop para enviar e receber mensagens
	for {

		// Enviar mensagem
		fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
		fmt.Println("                            MENU")
		fmt.Print("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-\n\n")

		fmt.Println("1 - Comprar passagem")
		fmt.Println("2 - Consultar passagem")
		fmt.Print("3 - Sair\n\n")
		fmt.Print("==========================================================\n\n")

		fmt.Scanln(&scan) // Recebe o comando que se deseja realizar

		mens_env = id + ":" + scan        // Concatena o id do cliente com o comando que se deseja realizar. A mensagem a ser enviada ao servidor
		enviar_mensagem(server, mens_env) // Envia a mensagem ao servidor

		switch scan {
		case "1":
			cabecalho(endereco)
			fmt.Println("——————————————————————————————————————————————————————————")
			fmt.Println("                    Comprar passagem")
			fmt.Print("——————————————————————————————————————————————————————————\n\n")

			// Solicitar rotas disponíveis do servidor
			rotas_disponiveis, erro = receber_dados[map[string]int](server)

			// Lista de rotas (apenas nome)
			var rotas_list []string
			var tamanho_palavra = -1
			// Preenchendo a lista de rotas
			for chave := range rotas_disponiveis {
				rotas_list = append(rotas_list, chave)
			}

			for k := range rotas_list {
				if len(rotas_list[k]) > tamanho_palavra {
					tamanho_palavra = len(rotas_list[k])
				}
			}

			// Tamanho do mapa de rotas
			tamanho := len(rotas_disponiveis)
			var colunas = 3                // Quantidade de colunas
			var linhas = tamanho / colunas // Quantidade de linhas se for um numero multiplo de colunas
			var resto = tamanho % colunas  // Resto da divisão se for maior que zero o numero nao é multiplo de colunas
			if resto > 0 {
				linhas += resto // Se o resto for maior que zero, adiciona mais elementos para completar a ultima linha
			}

			// Criando a matriz de rotas para exibição
			var matriz = make([][]string, linhas)
			k := 0
			for i := 0; i < linhas; i++ {
				matriz[i] = make([]string, colunas)
				for j := 0; j < colunas; j++ {
					if k >= tamanho {
						matriz[i][j] = ""
					} else {
						if rotas_disponiveis[rotas_list[k]] == 1 {
							matriz[i][j] = "\033[31m" + rotas_list[k] + "\033[0m"
						} else {
							matriz[i][j] = "\033[32m" + rotas_list[k] + "\033[0m"
						}
						k++
					}
				}
			}

			// Exibindo a matriz de rotas
			fmt.Println("Rotas disponiveis:")
			fmt.Println("----------------------------------------------------------")

			for i := 0; i < linhas; i++ {
				for j := 0; j < colunas; j++ {
					tam_cidade := len(matriz[i][j])
					espacamento := strings.Repeat(" ", tamanho_palavra+1-tam_cidade+12)
					fmt.Print(matriz[i][j], espacamento)
					k++
				}
				fmt.Println()
			}
			fmt.Print("----------------------------------------------------------\n\n")

			fmt.Println("——————————————————————————————————————————————————————————")
			fmt.Println("                           OPCOES")
			fmt.Print("——————————————————————————————————————————————————————————\n\n")

			fmt.Println("Digite o nome da rota para comprar passagem")
			fmt.Print("Digite 3 para voltar ao menu\n\n")
			fmt.Print("==========================================================\n\n")

			var scan string

			fmt.Scanln(&scan)
			fmt.Println("Scan:", scan)

			cabecalho(endereco)
			mens_env = id + ":" + scan        // Concatena o id do cliente com o comando que se deseja realizar. A mensagem a ser enviada ao servidor
			enviar_mensagem(server, mens_env) // Envia a mensagem ao servidor

			if scan != "3" {
				mens_receb = receber_mensagem(server)
				if mens_receb == "ok" {
					fmt.Println("Operação concluída com sucesso!")
				} else {
					fmt.Println("Erro ao processar a operação!")
				}
			}
			continue

		case "2":
			cabecalho(endereco)
			fmt.Println("——————————————————————————————————————————————————————————")
			fmt.Println("                    Consultar passagem")
			fmt.Print("——————————————————————————————————————————————————————————\n\n")
			mens_receb = receber_mensagem(server)
			if mens_receb == "ok" {
				rotas_compradas, erro = receber_dados[[]string](server)
				if erro != nil {
					fmt.Println("Erro ao receber rotas compradas:", erro)
					mens_env = "erro"
					enviar_mensagem(server, mens_env)
					continue
				}

				// tamanho da maior palavra de uma rota
				tamanho_rota := -1
				fmt.Println("Rotas compradas:")
				fmt.Println("----------------------------------------------------------")
				for i := 0; i < len(rotas_compradas); i++ {
					if len(rotas_compradas[i]) > tamanho_rota {
						tamanho_rota = len(rotas_compradas[i])
					}
				}
				for i := 0; i < len(rotas_compradas); i++ {
					espacamento := strings.Repeat(" ", tamanho_rota+1-len(rotas_compradas[i]))
					fmt.Print(rotas_compradas[i], espacamento)
					if (i+1)%3 == 0 {
						fmt.Println()
					}
				}
				fmt.Println()
				fmt.Println("----------------------------------------------------------")

				fmt.Println("——————————————————————————————————————————————————————————")
				fmt.Println("                           OPCOES")
				fmt.Print("——————————————————————————————————————————————————————————\n\n")

				fmt.Println("Digite o nome da rota para cancelar passagem")
				fmt.Print("Digite 3 para voltar ao menu\n\n")
				fmt.Print("==========================================================\n\n")

				fmt.Scanln(&scan)
				cabecalho(endereco)
				mens_env = id + ":" + scan        // Concatena o id do cliente com o comando que se deseja realizar. A mensagem a ser enviada ao servidor
				enviar_mensagem(server, mens_env) // Envia a mensagem ao servidor
				if scan != "3" {
					mens_receb = receber_mensagem(server)
					if mens_receb == "ok" {
						fmt.Println("Operação concluída com sucesso!")
					} else {
						fmt.Println("Erro ao processar a operação!")
					}
				}
				continue
			} else {
				fmt.Println("Mensagem recebida:", mens_receb+"\n")
				fmt.Print("——————————————————————————————————————————————————————————\n\n")
				continue
			}

		case "3":
			cabecalho(endereco)
			fmt.Println("Preparando para encerrar conexão...")
			mens_receb = receber_mensagem(server)
			fmt.Println("Mensagem recebida:", mens_receb)
			if mens_receb == "exit_ok" {
				fmt.Println("Encerrando conexão...")
				return
			} else {
				fmt.Println("Erro ao encerrar conexão")
				continue
			}

		default:
			cabecalho(endereco)
			mens_receb = receber_mensagem(server)
			fmt.Println("Mensagem recebida:", mens_receb+"\n")
			fmt.Print("=============================================\n\n")
			continue
		}
	}
}

func main() {
	limpar_terminal()
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

	cabecalho(endereco_alvo)                 // Exibe o cabeçalho com o endereço do servidor para conexão
	manipularConexao(conexao, endereco_alvo) // Manipula a conexão com o servidor
}
