# Usa a imagem oficial do Golang como base
FROM golang:1.21-alpine

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia o arquivo go.mod para o contêiner
COPY go.mod ./

# Baixa as dependências do projeto e gera o arquivo go.sum
RUN go mod download

# Copia o restante dos arquivos da aplicação para o contêiner
COPY app/ .

# Compila o código do servidor Go e gera o binário "server"
RUN go build -o server .

# Expõe a porta 8088 que o servidor vai utilizar
EXPOSE 8088

# Comando para rodar o binário do servidor
CMD ["./server"]