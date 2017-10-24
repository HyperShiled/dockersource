## Docker Image List
```
neunnsy/microservice-provider-user:0.0.1-SNAPSHOT
neunnsy/microservice-consumer-movie-ribbon-hystrix:0.0.1-SNAPSHOT
neunnsy/microservice-discovery-eureka-ha:0.0.1-SNAPSHOT
neunnsy/microservice-gateway-zuul:0.0.1-SNAPSHOT
neunnsy/microservice-hystrix-turbine:0.0.1-SNAPSHOT
neunnsy/microservice-hystrix-dashboard:0.0.1-SNAPSHOT
```

## Example Info (其中 192.168.250.49 为 Node 的IP地址)
```
kubectl create namespace spring-cloud
helm install charts/spring-cloud-suyan/ --name demo --namespace=spring-cloud

kubectl patch service peer1-demo-spring-cloud-suyan -p '{"spec": {"type": "NodePort"}}' -n spring-cloud
kubectl patch service peer2-demo-spring-cloud-suyan -p '{"spec": {"type": "NodePort"}}' -n spring-cloud
kubectl patch service consumer-service-demo-spring-cloud-suyan -p '{"spec": {"type": "NodePort"}}' -n spring-cloud
kubectl patch service gateway-service-demo-spring-cloud-suyan -p '{"spec": {"type": "NodePort"}}' -n spring-cloud
kubectl patch service hystrix-dashboard-service-demo-spring-cloud-suyan -p '{"spec": {"type": "NodePort"}}' -n spring-cloud
kubectl patch service turbine-service-demo-spring-cloud-suyan -p '{"spec": {"type": "NodePort"}}' -n spring-cloud

http://192.168.250.49:32508
http://192.168.250.49:32499
http://192.168.250.49:30995/user/1
http://192.168.250.49:32562/microservice-consumer-movie/user/1
http://192.168.250.49:32624/hystrix
http://192.168.250.49:32624/hystrix/monitor?stream=http%3A%2F%2F192.168.250.49%3A30995%2Fhystrix.stream
http://192.168.250.49:32624/hystrix/monitor?stream=http%3A%2F%2F192.168.250.49%3A30457%2Fturbine.stream
kubectl get deployments -n spring-cloud
kubectl scale --replicas=3 deployments/provider-demo-spring-cloud-suyan -n spring-cloud
kubectl get pods -n spring-cloud
kubectl get services -n spring-cloud
```