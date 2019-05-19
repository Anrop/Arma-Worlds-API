FROM golang:1.12-alpine AS builder
RUN apk --no-cache add git
WORKDIR /build/
COPY . /build/
RUN go build

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /build/Arma-Worlds-API /app/Arma-Worlds-API
CMD /app/Arma-Worlds-API
