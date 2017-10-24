FROM golang:1.7.5

MAINTAINER wang-x@neunn.com

RUN mkdir -p /go/src/suyan-service-catalog/

COPY . /go/src/suyan-service-catalog/
WORKDIR /go/src/suyan-service-catalog/

RUN cp -r /go/src/suyan-service-catalog/vendor/* /go/src/ && \
    go install -a -v github.com/beego/bee

EXPOSE 8000

ENTRYPOINT ["/go/bin/bee", "run", "-gendoc=true", "-downdoc=false"]