FROM neunnsy/zookeeper
MAINTAINER wang-x@neunn.com
ADD jmx_prometheus_javaagent-0.10.jar /zookeeper-3.4.10/bin/jmx_prometheus_javaagent-0.10.jar
ADD config.yaml /zookeeper-3.4.10/bin/config.yaml
ADD zkServer.sh /zookeeper-3.4.10/bin/zkServer.sh
EXPOSE 9000
