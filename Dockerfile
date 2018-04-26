FROM golang

WORKDIR /go/src/sevna

ADD . /go/src/sevna

RUN apt-get update -y

RUN echo "America/Sao_Paulo" > /etc/timezone

RUN dpkg-reconfigure -f noninteractive tzdata

RUN go get

RUN go install sevna

RUN apt-get install uuid-runtime -y

ENTRYPOINT ["/go/bin/sevna"]

EXPOSE 8080

CMD []


