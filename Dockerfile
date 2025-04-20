FROM golang:alpine AS builder

WORKDIR /app

COPY ./go.mod ./go.sum /app/

RUN go mod download

COPY . .

RUN go build -o tskrx ./cmd/main.go

EXPOSE 8080

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/tskrx .

CMD [ "./tskrx -c config.yml" ]