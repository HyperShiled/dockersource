FROM ibmcom/websphere-traditional:8.5.5.10-profile
MAINTAINER wang-x@neunn.com
ADD jmx_prometheus_javaagent-0.10.jar /opt/IBM/WebSphere/AppServer/bin/jmx_prometheus_javaagent-0.10.jar
ADD config.yaml /opt/IBM/WebSphere/AppServer/bin/config.yaml
ADD server.xml /opt/IBM/WebSphere/AppServer/profiles/AppSrv01/config/cells/DefaultCell01/nodes/DefaultNode01/servers/server1/server.xml
USER root
RUN chown -R was:was /opt/IBM/WebSphere/AppServer/bin/jmx_prometheus_javaagent-0.10.jar && \
    chown -R was:was /opt/IBM/WebSphere/AppServer/bin/config.yaml && \
    chown -R was:was /opt/IBM/WebSphere/AppServer/profiles/AppSrv01/config/cells/DefaultCell01/nodes/DefaultNode01/servers/server1/server.xml
USER was
EXPOSE 9000
