FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY ./ .

RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine:latest AS runtime

RUN apk add tzdata

WORKDIR /app

COPY --from=builder /app/app .
COPY config.yaml .

EXPOSE 8080

CMD ["./app", "http"]
