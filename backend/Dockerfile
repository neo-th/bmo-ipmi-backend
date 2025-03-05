FROM golang:latest

WORKDIR /app

COPY . /app

RUN go build -o bmo-ipmi .

EXPOSE 8080

CMD ["./bmo-ipmi"]