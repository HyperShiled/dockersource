FROM neunnsy/db2-service-broker-base:v1.0.0

MAINTAINER wang-x@neunn.com

RUN mkdir -p /home/db2inst1/go/src/github.com/compassorg/db2-service-broker/

COPY . /home/db2inst1/go/src/github.com/compassorg/db2-service-broker/
WORKDIR /home/db2inst1/go/src/github.com/compassorg/db2-service-broker/

RUN cp -r /home/db2inst1/go/src/github.com/compassorg/db2-service-broker/vendor/* /home/db2inst1/go/src/ && \
    chown -R db2inst1:db2inst1 /home/db2inst1/go/src/ && \
    chmod 755 /home/db2inst1/go/src/github.com/compassorg/db2-service-broker/docker/run.sh && \
    su - db2inst1 -c "go install -a -v github.com/beego/bee"

EXPOSE 8000

ENTRYPOINT ["/home/db2inst1/go/src/github.com/compassorg/db2-service-broker/docker/run.sh"]