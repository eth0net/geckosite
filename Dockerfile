FROM golang:1.15-alpine AS builder
WORKDIR /build
COPY go.mod .
RUN go mod download
RUN go mod verify
COPY . .
RUN go build -o main

FROM alpine:3.11.3
COPY --from=builder /build/main /app/main
EXPOSE 8080
CMD ["/app/main"]
