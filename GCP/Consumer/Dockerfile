FROM golang:1.19-alpine
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY . .

RUN go mod download

CMD ["go", "run", "main.go"]