FROM golang:alpine

WORKDIR /app

RUN go mod init main

COPY test.go /app

RUN go mod tidy

RUN go mod download

RUN go mod vendor

RUN go build -o pullimage .

ENTRYPOINT ["./pullimage"]