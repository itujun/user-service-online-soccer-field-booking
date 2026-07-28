FROM golang:alpine AS builder

RUN apk update
RUN apk add git openssh tzdata build-base python3 net-tools

WORKDIR /app

COPY .env.example .env
COPY . .

RUN go install github.com/buu700/gin@latest
RUN go mod tidy

RUN make build

FROM alpine:latest
RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    apk --no-cache add curl && \
    mkdir /app

WORKDIR /app

EXPOSE 8001

# PERBAIKAN: Hanya salin file yang benar-benar dibutuhkan untuk production
COPY --from=builder /app/user-service /app/
COPY --from=builder /app/.env /app/

ENTRYPOINT ["/app/user-service"]
