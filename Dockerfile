# Build go binary
FROM golang:1.17.1 AS builder
ENV GO111MODULE=on

RUN groupadd -g 10000 app && useradd -m -u 10001 -g app app

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY pkg ./pkg

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix nocgo -o main cmd/*.go

# Build application image
FROM alpine:3
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/main ./

USER app

ENTRYPOINT ["./main"]
