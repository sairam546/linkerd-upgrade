FROM golang:1.20 as builder
WORKDIR /app
COPY app-b/ .
COPY proto/ ./proto
RUN go mod init && go mod tidy && go build -o app-b

FROM gcr.io/distroless/base-debian11
WORKDIR /app
COPY --from=builder /app/app-b .
COPY --from=builder /app/proto ./proto
CMD ["/app/app-b"]