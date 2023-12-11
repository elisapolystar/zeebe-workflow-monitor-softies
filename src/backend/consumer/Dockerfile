FROM golang:1.21

ADD . /home

WORKDIR /home

RUN go build -o consumer .

EXPOSE 8000

CMD ["/home/consumer"]
