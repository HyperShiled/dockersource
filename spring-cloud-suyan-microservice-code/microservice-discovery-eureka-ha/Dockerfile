FROM java:8
VOLUME /tmp
ADD target/microservice-discovery-eureka-ha-0.0.1-SNAPSHOT.jar app.jar
RUN bash -c 'touch /app.jar'
EXPOSE 8761 8762
ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom","-jar","/app.jar"]