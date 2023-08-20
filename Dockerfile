FROM golang:1.19 AS builder
WORKDIR /app
ENV GOPATH=/
COPY . .
RUN apt-get update &&\
 go mod download &&\
 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o wbl0 ./cmd/main.go


FROM alpine:latest as production
WORKDIR /app
RUN apk add --no-cache postgresql-client
ENV GIN_MODE release
COPY .env wait-postgres.sh ./
COPY ./internal/template ./internal/template
RUN chmod +x wait-postgres.sh
COPY --from=builder /app/wbl0 .
CMD ["./wbl0"]