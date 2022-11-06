FROM golang:1.16-alpine
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY . .

RUN go mod download

EXPOSE 3030

CMD ["go", "run", "main.go"]