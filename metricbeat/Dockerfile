FROM golang:1.7.1
MAINTAINER Nicolas Ruflin <ruflin@elastic.co>

RUN set -x && \
    apt-get update && \
    apt-get install -y netcat python-virtualenv python-pip && \
    apt-get clean

# Setup work environment
ENV METRICBEAT_PATH /go/src/github.com/ashilokhvostov/beats/metricbeat

RUN mkdir -p $METRICBEAT_PATH/build/coverage
WORKDIR $METRICBEAT_PATH
