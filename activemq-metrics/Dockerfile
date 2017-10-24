FROM tianctrl/activemq:latest
MAINTAINER wang-x@neunn.com
ADD jmx_prometheus_javaagent-0.10.jar /opt/activemq/bin/jmx_prometheus_javaagent-0.10.jar
ADD config.yaml /opt/activemq/bin/config.yaml
ADD env /opt/activemq/bin/env
EXPOSE 9000
