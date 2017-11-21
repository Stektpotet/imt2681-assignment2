FROM golang:latest

MAINTAINER Halvor Smedås <Stektpotet@gmail.com>

ENV PROJECT_ROOT=$GOPATH/src/github.com/Stektpotet/imt2681-assignment2

WORKDIR $PROJECT_ROOT

# Root -> Workdir
COPY . .
RUN env
RUN ls -al ./*
RUN go build -v ./cmd/currencytrackr/

CMD ["./currencytrackr"]
