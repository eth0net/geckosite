FROM golang:1.15-alpine AS builder
WORKDIR /build
COPY . .
RUN go build -o main

FROM alpine:3.11.3
COPY --from=BUILD /build/main /app/main
CMD ["/app/main"]
