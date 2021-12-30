FROM golang:1.17.1-alpine

WORKDIR /
COPY . .

RUN go mod download

EXPOSE 5000

CMD ["go", "run", "main.go"]