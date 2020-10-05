FROM golang:1.15-alpine
WORKDIR /app
COPY go.mod .
RUN go mod download
RUN go mod verify
COPY . .
RUN go build -o main
EXPOSE 80
CMD ["/app/main"]
