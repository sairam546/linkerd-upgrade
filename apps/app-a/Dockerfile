FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN go mod init && go mod tidy && go build -o app-a

FROM gcr.io/distroless/base-debian11
WORKDIR /app
COPY --from=builder /app/app-a .
COPY --from=builder /app/proto ./proto
CMD ["/app/app-a"]