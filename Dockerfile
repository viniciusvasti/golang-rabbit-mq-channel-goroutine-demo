FROM golang:1.21.6-alpine3.19 AS base
RUN apk update
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN go build -o consumer ./cmd/consumer/main.go

FROM alpine:3.16.5 AS binary
ENV DOCKERIZE_VERSION v0.7.0
COPY --from=base /app/consumer .
EXPOSE 3000

RUN apk update --no-cache \
    && apk add --no-cache wget openssl \
    && wget -O - https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz | tar xzf - -C /usr/local/bin \
    && apk del wget

# dockerize wait for rabbitmq
ENTRYPOINT ["dockerize", "-wait", "tcp://rabbitmq:5672", "-timeout", "60s"]

CMD ["./consumer"]