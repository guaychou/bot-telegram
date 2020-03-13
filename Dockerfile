FROM golang:alpine as builder
RUN apk add git
RUN mkdir /app
WORKDIR /app
RUN go get github.com/gomodule/redigo/redis
RUN go get github.com/go-telegram-bot-api/telegram-bot-api
RUN go get github.com/guaychou/openweatherapi
COPY main.go /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bot-telegram-chou
RUN adduser -S -D -H -h /app appuser

FROM alpine:3.11
LABEL maintainer="Kevinchou kevin.harnanta@gmail.com"
# Spesifik timezone
ENV TZ="Asia/Jakarta"
RUN apk add tzdata
# Import from builder.
COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable
COPY --from=builder /app/bot-telegram-chou /app/bot-telegram-chou
# Use an unprivileged user.
USER appuser
# Run the binary.
ENTRYPOINT ["/app/bot-telegram-chou"]