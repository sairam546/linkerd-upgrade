FROM golang:1.20-alpine as builder
WORKDIR /app
COPY go.mod ./
RUN go mod tidy
COPY . .
RUN go build -o app1 .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app1 .
ENV APP2_HOST=app2
ENV APP2_PORT=8081
ENV APP1_PORT=8080
EXPOSE 8080
CMD ["./app1"]
