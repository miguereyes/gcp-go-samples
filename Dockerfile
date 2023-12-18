FROM golang:1.19-buster

WORKDIR /app

COPY . .

RUN go build -o helloworld

EXPOSE 8080

CMD ["./helloworld"]
