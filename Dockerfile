FROM golang:1.20 as builder

WORKDIR /build

COPY go.mod .
COPY go.sum .
COPY . .

RUN CGO_ENABLED=0 go build -buildvcs=false -o ./app .

FROM gcr.io/distroless/base-debian11

COPY --from=builder /build/app /app

ENTRYPOINT ["/app"]
