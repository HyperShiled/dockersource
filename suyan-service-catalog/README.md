## 在Kubernetes Cluster中创建Secret, 供Catalog与Broker之间通信使用
```
kubectl create -f conf/broker-secret.yaml
kubectl get secrets service-broker-secret -o yaml
```

## 启动API服务
```
bee run -gendoc=true -downdoc=false
```

## 使用源码构建Docker镜像
```
git clone http://10.20.0.196/hor/suyan-service-catalog.git
cd suyan-service-catalog/
docker build -t neunnsy/service-catalog-api:v1.1.0 .
docker run -d -p 8003:8000 -v `pwd`/conf/suyan-service-catalog.conf:/go/src/suyan-service-catalog/conf/suyan-service-catalog.conf --name suyan-service-catalog-api neunnsy/service-catalog-api:v1.1.0
```