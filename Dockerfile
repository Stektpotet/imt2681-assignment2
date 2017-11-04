FROM alpine:latest

MAINTAINER Halvor Smedås <Stektpotet@gmail.com>

WORKDIR "/opt"

ADD .docker_build/currencytrackr /opt/bin/currencytrackr
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["currencytrackr"]
