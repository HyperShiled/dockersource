FROM golang:1.7.5

MAINTAINER wang-x@neunn.com

RUN mkdir -p /go/src/github.com/compassorg/mongodb-service-broker/

COPY . /go/src/github.com/compassorg/mongodb-service-broker/
WORKDIR /go/src/github.com/compassorg/mongodb-service-broker/

RUN cp -r /go/src/github.com/compassorg/mongodb-service-broker/vendor/* /go/src/ && \
    chmod 755 /go/src/github.com/compassorg/mongodb-service-broker/docker/run.sh && \
    go install -a -v github.com/beego/bee

EXPOSE 8000

ENTRYPOINT ["/go/src/github.com/compassorg/mongodb-service-broker/docker/run.sh"]