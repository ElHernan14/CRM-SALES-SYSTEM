# ==========================
# Stage 1 - Build
# ==========================
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Instalar certificados (útil para conexiones TLS como Supabase)
RUN apk add --no-cache ca-certificates git

# Descargar dependencias primero (mejor cache)
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del proyecto
COPY . .

# Compilar el binario
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o nexora-api main.go

# ==========================
# Stage 2 - Runtime
# ==========================
FROM alpine:3.22

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/nexora-api .
COPY --from=builder /app/docs ./docs

EXPOSE 8081

CMD ["./nexora-api"]