FROM neunnsy/weblogic:12.2.1-domain
MAINTAINER wang-x@neunn.com
ADD jmx_prometheus_javaagent-0.10.jar /u01/oracle/user_projects/domains/base_domain/bin/jmx_prometheus_javaagent-0.10.jar
ADD config.yaml /u01/oracle/user_projects/domains/base_domain/bin/config.yaml
ADD startWebLogic.sh /u01/oracle/user_projects/domains/base_domain/bin/startWebLogic.sh
USER root
RUN chown -R oracle:oracle /u01/oracle/user_projects/domains/base_domain/bin/jmx_prometheus_javaagent-0.10.jar && \
    chown -R oracle:oracle /u01/oracle/user_projects/domains/base_domain/bin/config.yaml && \
    chown -R oracle:oracle /u01/oracle/user_projects/domains/base_domain/bin/startWebLogic.sh
USER oracle
EXPOSE 9000
