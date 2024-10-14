FROM golang:1.23-alpine AS builder

COPY . /github.com/erikqwerty/chat-server/
WORKDIR /github.com/erikqwerty/chat-server/

RUN go mod download
RUN go build -o ./bin/chat-server cmd/chat-server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/erikqwerty/chat-server/bin/chat-server .

ADD .env .

CMD ["./chat-server"]