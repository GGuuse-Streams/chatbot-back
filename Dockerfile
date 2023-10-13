FROM node:18-alpine as builder
COPY --from=golang:1.21.0-alpine /usr/local/go/ /usr/local/go/
ENV PATH="$PATH:/usr/local/go/bin"
ENV PATH="$PATH:/root/go/bin"

WORKDIR /app

RUN apk add upx

COPY . .

FROM builder as bot_builder
RUN cd apps/bot && \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ./out ./cmd/main.go && upx -9 -k ./out

FROM alpine:latest as bot
WORKDIR /app
COPY --from=bot_builder /app/apps/bot/out /bin/bot
COPY --from=bot_builder /app/config/config.yml /config/config.yml
CMD ["/bin/bot"]

FROM builder as commands_builder
RUN cd apps/commands && \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ./out ./cmd/main.go && upx -9 -k ./out

FROM alpine:latest as commands
WORKDIR /app
COPY --from=commands_builder /app/apps/commands/out /bin/commands
CMD ["/bin/commands"]