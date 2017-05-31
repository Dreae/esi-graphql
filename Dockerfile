FROM golang:1.8

RUN apt-get update && apt-get -y upgrade

RUN go get -u github.com/Masterminds/glide
RUN go get -u github.com/jteeuwen/go-bindata/go-bindata

COPY . /go/src/github.com/dreae/esi-graphql

WORKDIR /go/src/github.com/dreae/esi-graphql

RUN glide install
RUN go-bindata assets assets/schema
RUN go build

CMD ./esi-graphql
