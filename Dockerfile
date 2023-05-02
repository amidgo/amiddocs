FROM golang:1.19-alpine3.17 as builder
WORKDIR /app
COPY go.mod /app/
RUN go mod download
ADD . /app/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/fiber /app/cmd/main/main.go

FROM alpine:3.17
COPY --from=builder /app/files /files
COPY --from=builder /app/config /config
COPY --from=builder /app/fiber /fiber
EXPOSE 10101
ENTRYPOINT [ "/fiber" ]
