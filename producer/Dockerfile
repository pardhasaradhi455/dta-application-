FROM golang:1.23-alpine3.21 AS builder
RUN apk add --no-progress --no-cache gcc musl-dev
WORKDIR /build
COPY . .
RUN go mod download

RUN go build -tags musl -ldflags '-extldflags "-static"' -o /build/main

FROM scratch
WORKDIR /app
COPY --from=builder /build/main .
EXPOSE 8080
ENTRYPOINT ["/app/main"]