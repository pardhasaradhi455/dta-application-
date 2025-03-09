FROM golang:1.23 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-example

FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /go-example /go-example

EXPOSE 7071

USER nonroot:nonroot

ENTRYPOINT ["/go-example"]