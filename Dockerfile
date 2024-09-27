# Step 1: Modules caching
FROM golang:1.22-alpine as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.22-alpine as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOARCH=amd64 \
    go build -o /bin/app ./cmd/main.go

FROM alpine:latest
COPY --from=builder /app/configs /configs
COPY --from=builder /bin/app /app
COPY wait-for-db.sh /wait-for-db.sh

# Делаем скрипт исполняемым
RUN chmod +x wait-for-db.sh

# Указываем скрипт ожидания для запуска
CMD ["./wait-for-db.sh", "./app"]