FROM golang:1.21

ADD . /home

WORKDIR /home

RUN go get .

EXPOSE 8000

CMD ["go","run","."]
