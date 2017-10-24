# db2-service-broker
db2 service broker for kubernetes


## Start Service Broker
```
bee run -gendoc=true -downdoc=false
```

## 证书名称和证书内容示例
```
connect_uri
DATABASE=sample; HOSTNAME=192.168.21.87; PORT=50000; PROTOCOL=TCPIP; UID=db2admin; PWD=12345678; FILE_ROOT=C:\Tablespaces\
```


## DB2 的关键操作示例
```
# 创建bufferpool和tablespace
create bufferpool bp32k all nodes size -1 pagesize 32k
create regular tablespace  tablespace1 pagesize 32k managed by database using(file 'C:\Tablespaces\tablespace1' 5g) bufferpool bp32k
create regular tablespace  tablespace2 pagesize 32k managed by database using(file 'C:\Tablespaces\tablespace2' 5g) bufferpool bp32k
grant use of tablespace  tablespace1 to user db2admin

# 删除bufferpool和tablespace
drop bufferpool bp32k
drop tablespace tablespace1

# 查看tablespace
list tablespaces
list tablespaces show detail

# 根据名称查询bufferpool
select bpname,npages,pagesize from syscat.bufferpools where bpname='BP81C0373AF932K'

# broker的实现里类似于下面这种:
create bufferpool bpac6a817f3532k all nodes size -1 pagesize 32k
create regular tablespace tp7f2630994c pagesize 32k managed by database using(file 'C:\Tablespaces\tp7f2630994c' 5g) bufferpool bpac6a817f3532k
```