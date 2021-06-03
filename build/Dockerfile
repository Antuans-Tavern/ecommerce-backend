FROM golang:1.16 AS builder
COPY . .
RUN go mod tidy
RUN go build -o /ecommerce ./cmd/ecommerce

FROM scratch
COPY --from=builder /ecommerce /ecommerce
CMD ["/ecommerce"]
