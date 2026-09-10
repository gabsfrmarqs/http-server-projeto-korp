FROM golang:1.27.0

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o httpserver httpserver.go

EXPOSE 8080

CMD ["./httpserver"]

