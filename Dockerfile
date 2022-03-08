FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /belajardocker

COPY . .

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/belajardocker/binary"]