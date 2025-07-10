FROM golang:1.24 AS base

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin/gateway cmd/app/main.go

ENV PORT=8000
EXPOSE 8000

# Start the application
CMD ["/build/bin/gateway"]
