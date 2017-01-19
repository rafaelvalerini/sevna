FROM golang

WORKDIR /go/src/agregador

ADD . /go/src/agregador

RUN apt-get update -y

RUN apt-get install tzdata -y

ENV TZ=America/Sao_Paulo

RUN rm -rf /var/cache/apk/*

RUN go get

RUN go install agregador

RUN apt-get install uuid-runtime -y

ENTRYPOINT ["/go/bin/agregador"]

EXPOSE 8080

CMD []


