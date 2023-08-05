FROM golang:latest

WORKDIR /app

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/server/main.go
RUN GOOS=linux CGO_ENABLED=1 go build -ldflags="-w -s" -o server ./cmd/server/main.go

CMD ["./server"]
