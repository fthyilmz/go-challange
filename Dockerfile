FROM golang:1.18 AS builder

RUN apt-get update \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY .. .

RUN go test -v ./...

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o main .

FROM scratch

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/.env ./.env
COPY --from=builder /app/main ./main


ENTRYPOINT ["/app/main"]