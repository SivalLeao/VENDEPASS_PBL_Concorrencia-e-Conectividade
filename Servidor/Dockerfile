# Use a imagem oficial do Golang como base
FROM golang:1.23-alpine

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia go.mod para o contêiner
COPY go.mod ./

# Faz o download das dependências (isso criará o go.sum)
RUN go mod download

# Copia o restante dos arquivos do projeto para o contêiner
COPY app .

# Compila o servidor Go
RUN go build -o server .

# Expõe a porta que o servidor vai usar
EXPOSE 8088

# Comando para rodar o servidor compilado
CMD ["./server"]
