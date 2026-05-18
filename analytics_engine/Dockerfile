FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/analytics-api ./cmd/api

FROM gcr.io/distroless/static-debian12
COPY --from=builder /out/analytics-api /analytics-api
EXPOSE 8080
ENTRYPOINT ["/analytics-api"]
