FROM golang:latest

RUN mkdir /app

ADD . /app/

WORKDIR /app

EXPOSE 8080

RUN GOOS=linux go build -o main ./main.go

RUN ["chmod", "+x", "/app/main"]

CMD ["/app/main"]
