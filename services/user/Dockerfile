FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o user-service

EXPOSE 8081

CMD ["./user-service"]
