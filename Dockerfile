FROM golang:1.22 as builder
WORKDIR /app
COPY go.* ./
RUN go mod download

RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' -o /chanif-go-api .
FROM alpine:3.14
EXPOSE 3000
CMD ["/chanif-go-api"]
