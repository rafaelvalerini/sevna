FROM golang

WORKDIR /go/src/agregador

ADD . /go/src/agregador

RUN apt-get update -y

RUN echo "America/Sao_Paulo" > /etc/timezone

RUN dpkg-reconfigure -f noninteractive tzdata

RUN go get

RUN go install agregador

RUN apt-get install uuid-runtime -y

ENTRYPOINT ["/go/bin/agregador"]

EXPOSE 8080

CMD []


