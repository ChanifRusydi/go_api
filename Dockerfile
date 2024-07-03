FROM golang:1.22
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go get -d -v ./...
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' -o /chanif-go-api .
EXPOSE 3000
CMD ["/chanif-go-api"]