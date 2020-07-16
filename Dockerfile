FROM golang:1.13

RUN mkdir /app

ADD . /app/

WORKDIR /app

EXPOSE 8080

CMD ["go","run","main.go","client.go"]