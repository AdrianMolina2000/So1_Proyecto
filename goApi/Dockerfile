FROM golang:1.19-alpine

WORKDIR /
COPY . .

RUN go mod download

EXPOSE 8000

CMD ["go", "run", "main.go"]