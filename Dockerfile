FROM node:22.14.0-alpine AS frontend-builder

WORKDIR /app/web

COPY web/package*.json ./

RUN npm install && \
    npm cache clean --force

COPY web/ .

RUN npm run generate

FROM golang:1.24.0-alpine AS backend-builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY --from=frontend-builder /app/web/dist /app/server/web

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/PowerKey .


FROM alpine:latest

WORKDIR /app

ENV TZ=Asia/Shanghai

RUN apk add --no-cache --update samba shadow net-tools libcap tzdata && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone && \
    addgroup -S power && \
    adduser -S power -G power -h /app

COPY --from=backend-builder --chown=power:power /app/PowerKey /app/PowerKey

USER power

EXPOSE 3000

CMD ["/app/PowerKey"]