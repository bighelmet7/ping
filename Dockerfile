FROM golang

RUN mkdir /sample

ADD . /sample

WORKDIR /sample

RUN go build simple.go

CMD ["./simple"]
