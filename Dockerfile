FROM golang:latest

add agregador /tmp

run chmod +x /tmp/agregador

expose 8080

WORKDIR /tmp

cmd ["/tmp/agregador"]