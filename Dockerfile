# build stage
FROM golang:1.13.0-alpine3.10 as builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN apk add --no-cache make curl
WORKDIR /home/app
COPY . /home/app
RUN make build

# runtime stage
FROM alpine:3.10.2
WORKDIR /app
COPY --from=builder /home/app/bin/ /app/
RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*
