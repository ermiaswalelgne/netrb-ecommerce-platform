FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o order-service

EXPOSE 8082

CMD ["./order-service"]
