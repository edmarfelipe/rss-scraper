FROM golang:latest AS builder
WORKDIR /app

COPY go.mod go.sum ./
COPY cmd ./cmd
COPY docs ./docs
COPY internal ./internal
COPY sql ./sql

RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/apiserver/main.go

FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=builder /app/main ./
EXPOSE 8080
ENTRYPOINT [ "/app/main" ]