FROM golang:1.23-alpine
WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /run-api main.go

CMD ["/run-api"]
EXPOSE 80
