FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build

EXPOSE 1323

CMD ["go", "run", "main.go"]