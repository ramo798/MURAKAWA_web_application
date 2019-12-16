FROM golang:latest

WORKDIR /go-app
COPY . .

COPY go.mod .
COPY go.sum .

RUN go mod download

CMD ["go", "run", "main.go"]