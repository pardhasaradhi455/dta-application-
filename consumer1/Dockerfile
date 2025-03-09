FROM golang:1.23-alpine3.21 AS builder
RUN apk add --no-progress --no-cache gcc musl-dev
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -tags musl -ldflags '-extldflags "-static"' -o /build/main

FROM scratch
WORKDIR /app
COPY --from=builder /build/main .
COPY --from=builder /build/templates /app/templates
EXPOSE 8081
ENTRYPOINT ["/app/main"]