FROM golang:1.24.4-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o backend-brighted ./cmd/main.go
# RUN go install github.com/cosmtrek/air@latest

EXPOSE 3000

CMD ["./backend-brighted"]