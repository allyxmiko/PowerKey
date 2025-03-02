FROM node:22.14.0-alpine AS frontend-builder

WORKDIR /app/web_src

COPY web_src/package*.json ./

RUN npm install && \
    npm cache clean --force

COPY web_src/ .

RUN npm run generate

FROM golang:1.24.0-alpine AS backend-builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY --from=frontend-builder /app/web_src/dist /app/server/static

RUN apk add --no-cache musl-dev gcc && \
    CGO_ENABLED=1 GOOS=linux go build -o /app/PowerKey .


FROM alpine:latest

WORKDIR /app

ENV TZ=Asia/Shanghai

RUN apk add --no-cache --update samba shadow net-tools libcap tzdata && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY --from=backend-builder /app/PowerKey /app/PowerKey

EXPOSE 3000

CMD ["/app/PowerKey"]