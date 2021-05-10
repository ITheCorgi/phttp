FROM golang:1.16.2-alpine3.13 AS builder
RUN apk add --no-cache git
WORKDIR /github.com/ITheCorgi/phttp/

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go test -v
RUN GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /github.com/ITheCorgi/phttp/.bin/app .
EXPOSE 8080
CMD [ "./app" ]