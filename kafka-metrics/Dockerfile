FROM neunnsy/kafka
MAINTAINER wang-x@neunn.com
ADD jmx_prometheus_javaagent-0.10.jar /opt/kafka/bin/jmx_prometheus_javaagent-0.10.jar
ADD config.yaml /opt/kafka/bin/config.yaml
ADD kafka-run.sh /opt/kafka/bin/kafka-run.sh
ADD kafka-server-start.sh /opt/kafka/bin/kafka-server-start.sh

RUN chmod 755 /opt/kafka/bin/kafka-run.sh
RUN chmod 755 /opt/kafka/bin/kafka-server-start.sh

EXPOSE 9000