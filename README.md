# star-edge-cloud

star-edge-cloud是一个边缘计算（edge computing）-云计算的开源软件平台，可以为监测类项目提供一个可靠、简便的解决方案。

软件架构包括边缘端和云端两个部分，边缘端负责接入采集卡、智能设备和各类传感器，可以进行数据的压缩、过滤和缓存，集成算法和业务处理，将数据传到云端。

云端负责数据收集整理，并将数据存储，弹性扩展业务功能。

## 边缘端
边缘端目标是采集数据，集成算法。需要完成以下功能：
> * 多传感器情况下的大数据量(包括数据的复制问题）
> * 可以扩展算法模块和其他模块
> * 智能网关，断路器的容错机制--健康检查
> * 调度功能
> * 规则引擎
> * 消息总线机制--传输层
> * 考虑视频业务情况
> * 日志功能，记录系统、传感器的状态，数据可回溯
> * 应具有数据存储（或缓存）功能，考虑使用一种特别数据库
> * 多协议的支持

### 云端
基于docker的容器云平台，目标是汇集数据，进行计算。需要完成以下功能：
> * 数据存储，便于以后大数据分析
> * 具有RestAPI的数据接口
> * 具有历史数据导入功能
> * docker云应具有网络隔离功能

### 监控界面
可视化呈现数据。需要完成以下功能:

1. Web界面
> * 基于角色权限功能
> * 统计设备状态信息（不同维度）
> * 设备信息

2. 工具包
> * 实现一些类似诊断分析，故障修复等功能的工具

3. 移动App
> * 留待2.0开发

## 部署教程
在Linux--Deepin15.5下,进入deploy目录，执行编译脚本。

### edge端

1. 执行部署命令：
```
cd deploy/
sudo chmod +x edge.sh
./edge.sh
```
2. 运行系统：
```
sudo ./core
```
3. 访问：[http://localhost:21000/html/index.html](http://localhost:21000/html/index.html)
![edge](./images/edge.png)
4. 运行log服务
5. 运行store服务
6. 添加设备设备，选择compile目录下编译好的文件
7. 添加扩展设备，选择compile目录下编译好的文件

### cloud端

测试坏境搭建（安装docker,docker-compose略)：

1. 打开2375端口（后面将使用tls访问，开启2376）：
```
vi /lib/systemd/system/docker.service
```
2. 找到Execstart=/usr/bin/dockerd后加上
```
-H tcp://0.0.0.0:2375 -H unix://var/run/docker.sock  
```
保存并且退出

```
systemctl daemon-reload
service docker restart//重启启动docker
systemctl stats docker//可以查看相关内容，看看2375是否已经设置好
```

3. 访问和验证：
[http://localhost:2375/info](http://localhost:2375/info)

4. 拉取hbase容器 
```
docker pull harisekhon/hbase 
```
5. 启动容器 
```
docker run -d -h myhbase -p 2181:2181 -p 8080:8080 -p 8085:8085 -p 9090:9090 -p 9095:9095 -p 16000:16000 -p 16010:16010 -p 16201:16201 -p 16301:16301 -p 16020:16020 -p 16030:16030 --name hbase1.3 harisekhon/hbase
```

6. 访问及验证hbase
[http://localhost:16010/master-status](http://localhost:16010/master-status)

7. 执行命令：
```
cd deploy/
sudo chmod +x cloud.sh
mvn clean package
./cloud.sh
```
8. 运行：
```
java -jar caas*.jar
#这种方法还没有尝试:nohup java -jar ***.jar &
```
9. 部署tomcat
下载：[http://mirrors.shu.edu.cn/apache/tomcat/tomcat-8/v8.5.37/bin/apache-tomcat-8.5.37.tar.gz](http://mirrors.shu.edu.cn/apache/tomcat/tomcat-8/v8.5.37/bin/apache-tomcat-8.5.37.tar.gz)
10. 进入目录后运行：
```
./startup.sh
```
11. 拷贝display下web项目到webapps之中
```
cp -r */web */webapps/
```
12. 访问及验证：
[http://localhost:8080/web/index.html](http://localhost:8080/web/index.html)
![cloud](./images/cloud.png)
> 注：
> 1.目前仅仅是验证版本，尚有很多很多功能没有完成，部分功能还有Bug，但这只是开始   
> 2.查看sqlite数据，可以使用SQLiteStudio


QQ交流群：590749338