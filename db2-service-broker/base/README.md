## Build Docker Image Note
```
go1.7.5.linux-amd64.tar.gz 需要从官网下载,然后放在base目录下,然后才能做基础镜像的构建
```

## Build Base Docker Image
```
docker build -t neunnsy/db2-service-broker-base:v1.0.0 .
```

## Pull Base Docker Image From Docker Hub
```
docker pull neunnsy/db2-service-broker-base:v1.0.0
```