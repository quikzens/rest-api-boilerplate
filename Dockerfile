# Build stage
FROM golang:1.18-alpine3.16 AS builder
WORKDIR /app 
COPY . .
RUN go build -o main main.go
       
# Run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .

EXPOSE 3000
CMD [ "/app/main" ]