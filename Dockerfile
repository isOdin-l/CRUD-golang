FROM golang:1.24-alpine

WORKDIR /

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./rest-api-server ./cmd/app/

CMD ["./rest-api-server"]