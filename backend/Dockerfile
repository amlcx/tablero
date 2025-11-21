FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o tablero_backend ./cmd/serve/main.go

# Runtime phase
FROM alpine:3.22

RUN adduser -D -H tablero_backend_user

USER tablero_backend_user

COPY --from=builder /app/tablero_backend /usr/local/bin/tablero_backend

WORKDIR /app

ENV TABLERO_SERVER_HOSTNAME=${TABLERO_SERVER_HOSTNAME}
ENV TABLERO_SERVER_PORT=${TABLERO_SERVER_PORT}
ENV TABLERO_JWKS_URL=${TABLERO_JWKS_URL}

CMD ["tablero_backend"]