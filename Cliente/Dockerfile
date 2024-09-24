# Etapa de build: Utiliza uma imagem base do Go
FROM golang:1.23-alpine

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia o arquivo go.mod para o diretório de trabalho no contêiner
COPY go.mod ./

# Baixa as dependências do projeto definidas no go.mod
RUN go mod download

# Copia os arquivos da aplicação para o diretório de trabalho no contêiner
COPY app/ .

# Compila o código Go e gera o binário "client"
RUN go build -o client .

# Exponha a porta 8088 (caso o aplicativo utilize essa porta)
EXPOSE 8088

# Comando que será executado quando o contêiner iniciar
CMD ["./client"]