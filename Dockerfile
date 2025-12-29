# syntax=docker/dockerfile:1

# --- Estágio de Build ---
FROM golang:1.23-alpine AS builder

# Instala certificados CA para chamadas HTTPS
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Cache de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código fonte
COPY . .

# Build otimizado:
# CGO_ENABLED=0 para linkagem estática (roda no scratch)
# -ldflags="-s -w" reduz o tamanho do binário (remove debug symbols)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /gopportunities ./main.go

# --- Estágio Final ---
FROM scratch

# Copia certificados de segurança do builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copia o binário compilado
COPY --from=builder /gopportunities /gopportunities

# Porta que o Gin costuma usar
EXPOSE 8080

ENTRYPOINT ["/gopportunities"]