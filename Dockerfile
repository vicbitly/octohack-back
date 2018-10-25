FROM golang:latest

ADD . /go/src/app/

WORKDIR /go/src/app

EXPOSE 8080

RUN go get
RUN GOOS=linux go build -o main ./main.go

RUN ["chmod", "+x", "/go/src/app/main"]

CMD ["/go/src/app/main"]