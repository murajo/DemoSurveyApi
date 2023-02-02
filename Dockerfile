FROM golang:latest

RUN go mod init main \
  && go mod tidy \
  && go build

EXPOSE 1323


CMD ["go", "run", "main.go"]