# Etapa de build
FROM golang:1.23-alpine

# Diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar todos os arquivos do diretório atual para o diretório de trabalho no contêiner
COPY go.mod ./

# Compilar o código Go
RUN go mod download

# Definir diretório de trabalho no contêiner
COPY app .

# Copiar o binário da etapa de build para a etapa de execução
RUN go build -o client .

# Expor a porta que o aplicativo usa (caso necessário)
EXPOSE 8088

# Comando de entrada ao iniciar o contêiner
CMD ["./client"]
