# Etapa 1: construir o binário Go
FROM golang:1.25 AS builder

# Criar diretório de trabalho dentro do container
WORKDIR /app

# Copiar go.mod e go.sum e baixar dependências
COPY go.mod go.sum ./
#RUN go mod download
RUN go env && go mod download

# Copiar o restante do código
COPY . .

# Compilar o binário (CGO habilitado por causa do SQLite)
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# Etapa 2: criar imagem leve para rodar o binário
FROM debian:bookworm-slim

WORKDIR /root/

# Copiar o binário do builder
COPY --from=builder /app/main .

# Copiar o banco (opcional)
COPY --from=builder /app/tasks.db .

# Instalar dependências mínimas para SQLite
RUN apt-get update && apt-get install -y ca-certificates sqlite3 && rm -rf /var/lib/apt/lists/*

# Expõe a porta usada pelo Gin
EXPOSE 8080

# Comando de inicialização
CMD ["./main"]
