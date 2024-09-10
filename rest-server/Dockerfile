FROM golang:alpine3.20 AS builder

RUN sed 's/https/http/g' -i /etc/apk/repositories
RUN apk update && apk add --no-cache ca-certificates git upx

WORKDIR /build
ADD . .

WORKDIR /build/hello

ENV GOPRIVATE=proxy.golang.org/**,github.com/**
RUN go mod download
RUN go build
RUN upx -7 hello

FROM alpine:3.20 AS runner

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

WORKDIR /home/appuser

COPY --chown=appuser --from=builder /build/hello/hello ./hello

ENTRYPOINT ["./hello"]