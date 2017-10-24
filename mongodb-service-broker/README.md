# mongodb-service-broker
mongodb service broker for kubernetes


## Start Service Broker
```
bee run -gendoc=true -downdoc=false
```

## 证书名称和证书内容示例
```
connect_uri
super:12345678@192.168.21.87:27017
```

## 在Windows上的初始化步骤如下:
```
mongod --port 27017 --dbpath C:\Database\MongoDB\Server\3.4\data
mongo --port 27017
use admin
db.createUser(
  {
    user: "super",
    pwd: "12345678",
    roles: [ { role: "root", db: "admin" } ]
  }
)

mongod --auth --bind_ip 192.168.21.87 --port 27017 --dbpath C:\Database\MongoDB\Server\3.4\data
mongo --host 192.168.21.87 -u "super" -p "12345678" --authenticationDatabase "admin"
```

## 参考文献
```
https://docs.mongodb.com/manual/tutorial/enable-authentication/
https://docs.mongodb.com/manual/reference/built-in-roles/
http://blog.csdn.net/weiwangsisoftstone/article/details/39269373
```

## 验证工具 Robomongo 的网址
```
https://robomongo.org/download
```