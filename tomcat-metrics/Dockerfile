FROM tomcat:7
MAINTAINER wang-x@neunn.com
ADD catalina.sh /usr/local/tomcat/bin/catalina.sh
ADD jmx_prometheus_javaagent-0.10.jar /usr/local/tomcat/bin/jmx_prometheus_javaagent-0.10.jar
ADD tomcat.yaml /usr/local/tomcat/bin/tomcat.yaml
EXPOSE 8000 9000
CMD ["catalina.sh", "run"]
