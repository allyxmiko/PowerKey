FROM golang:1.24.0-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/PowerKey .

FROM alpine:latest

ENV TZ=Asia/Shanghai

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apk add --no-cache --update samba shadow net-tools libcap tzdata

RUN addgroup -S power && \
    adduser -S power -G power -h /app

COPY --from=builder --chown=power:power /app/PowerKey /app/PowerKey

WORKDIR /app

USER power

EXPOSE 3000

CMD ["/app/PowerKey"]