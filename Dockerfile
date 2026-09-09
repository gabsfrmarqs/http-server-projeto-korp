FROM golang:1.27.0

COPY httpserver.go .
COPY go.mod .

RUN go build -o httpserver httpserver.go

COPY httpserver .

EXPOSE 8080

CMD ["./httpserver"]

