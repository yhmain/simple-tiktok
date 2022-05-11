# simple-demo

main 分支
最终展示所用，慎改 合并

## 抖音项目服务端简单示例

具体功能内容参考飞书说明文档

工程无其他依赖，直接编译运行即可

```shell
对于windows系统
下面命令在git bash上运行即可
go build && ./simple-tiktok
或者go build之后，会生成一个simple-tiktok.exe，点击运行也行
```

### 功能说明

接口功能不完善，仅作为示例

* 用户登录数据保存在内存中，单次运行过程中有效
* 视频上传后会保存到本地 public 目录中，访问时用 127.0.0.1:8080/static/video_name 即可
* 注意app里面配置前缀url为服务器（自己电脑）的 http://ip:port

### 测试数据

测试数据写在 demo_data.go 中，用于列表接口的 mock 测试