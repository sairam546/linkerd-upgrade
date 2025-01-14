FROM golang:1.20-alpine as builder
WORKDIR /app
COPY go.mod ./
RUN go mod tidy
COPY . .
RUN go build -o app2 .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app2 .
EXPOSE 8081
CMD ["./app2"]
