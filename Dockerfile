FROM golang as builder
WORKDIR /app
ADD . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -o goccapi

FROM alpine
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/goccapi .
CMD ["/app/goccapi"]