FROM golang:1.21

ADD . /home

WORKDIR /home

RUN go build -o dbclean .

CMD ["/home/dbclean"]
