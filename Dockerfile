FROM golang

WORKDIR /go/src/agregador

ADD . /go/src/agregador

RUN go get

RUN go install agregador

RUN apt-get update

RUN apt-get install uuid-runtime -y

ENTRYPOINT ["/go/bin/agregador"]

EXPOSE 8080

CMD []


