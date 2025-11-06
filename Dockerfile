FROM golang:1.24

WORKDIR /

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# --- For building static binary file (image size approximatly 35MB) ---
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /rest-api-server ../cmd/rest-api-server/main.go
# FROM scratch

RUN go build -o /rest-api-server ./cmd/rest-api-server/main.go

FROM ubuntu
COPY --from=0 /rest-api-server /bin/rest-api-server

CMD ["/bin/rest-api-server"]