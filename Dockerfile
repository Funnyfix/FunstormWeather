FROM golang:alpine AS builder
WORKDIR /build
ADD go.mod .
COPY . .
RUN go build -o funstorm main.go
FROM alpine
WORKDIR /build
COPY --from=builder /build/funstorm /build/funstorm
RUN ls -all
ENTRYPOINT [ "/build/funstorm" ]