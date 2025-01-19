FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o product-service

EXPOSE 8080

CMD ["./product-service"]