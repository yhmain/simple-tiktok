
[Github地址](https://github.com/yhmain/simple-tiktok)  

# 项目名称：simple-tiktok
main 分支
最终展示所用，慎改 合并

## 抖音项目服务端简单示例

具体功能内容参考飞书说明文档
https://bytedance.feishu.cn/docx/doxcnbgkMy2J0Y3E6ihqrvtHXPg

## 项目部署（主要针对Windows系统，Linux类似吧）
1. 首先通过ipconfig获取本机或者服务器的IP，修改main.go里面的常量SERVER_IP="Your IP:Port"
2. 修改service目录下的config.go里面的常量，举例如下：
即设置静态资源的路径
```
PREFIX_VIDEOS string = "http://192.168.1.108:8080/static/videos/"
PREFIX_COVERS string = "http://192.168.1.108:8080/static/covers/"
```
3. MySQL配置
在本地新建MySQL的数据库后，（运行tiktok.sql即可）
修改dao目录下的config.go文件各个相关属性

```
package dao

const (
	FEED_VIDEOS_NUM = 30 //视频数由服务端控制，单次最多30个

	//MySQL配置如下
	USER          = "root"
	PASSWORD      = "123456"
	SERVERIP      = "127.0.0.1"
	PORT          = "3306"
	DATABASE_NAME = "tiktok"
)
```

## 项目运行
工程无其他依赖，直接编译运行即可。  
Windows系统 运行方式（2种）：
1. 在Git Bash上输入命令再回车：`go build && ./simple-tiktok`
2. 在终端或者命令行上依次输入：
```
go build
simple-tiktok
```
- 注意app里面配置前缀url为服务器（自己电脑）的 http://ip:port  
- 关于.apk文件，可以安装在手机上，也可在电脑上安装一个安卓模拟器。  
这里我使用的是[MuMu模拟器](https://mumu.163.com/)
（注意：其他方式可能会在发布视频的时候出现上传错误）

Linux系统运行方式：`go build && ./simple-tiktok`  

## 项目架构
![cedd6c1e71611846cc2c1ea1179e7c4.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/394373ff35f94df1a97343901f4d554d~tplv-k3u1fbpfcp-watermark.image?)

controller：控制层，主函数中的路由调用  
dao：连接数据库，对数据库进行操作  
model：对应数据库中的每个实体  
service：逻辑层，处理核心业务的逻辑输出  

public：存储静态资源  

main.go：主函数
router.go：初始化路由  
tiktok.sql：数据库设计文件  

## 技术栈
语言：Go  
底层存储：MySQL
ORM框架：GORM
HTTP框架：Gin


## 功能说明

接口功能不完善，仅作为示例

* 用户登录数据保存在内存中，单次运行过程中有效
* 视频上传后会保存到本地 public 目录中，访问时用 
* 注意app里面配置前缀url为服务器（自己电脑）的 http://ip:port

## 测试数据

//测试数据写在 demo_data.go 中，用于列表接口的 mock 测试

## 问题难点：
1. Go语言如何从一个mp4文件中抽取某一帧作为封面？  
	ffmpeg  
2. 池化技术的应用
3. 