# syntax=docker/dockerfile:1
FROM golang:1.24-alpine AS build
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o server ./cmd

FROM alpine:3.20
WORKDIR /app
COPY --from=build /app/server .
CMD ["./server"]
