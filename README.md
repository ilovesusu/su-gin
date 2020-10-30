- [x] GIN
- [ ] JWT
- [x] 存储
    - [x] Redis
    - [x] Mysql
        - [x] GORM
    - [ ] Mongodb
- [x] I18N
- [x] CONF
- [x] GO MOD
- [ ] ip 限流
- [ ] Https
- [ ] 请求限流
- [ ] 请求跨域
- [x] 支持优雅地重启或停止
- [x] 签名验证
    - [x] MD5 组合加密
    - [ ] AES 对称加密
    - [ ] RSA 非对称加密
- [ ] 请求日志记录
    - [ ] 请求日志记录
    - [ ] 访问日志记录
- [ ] 链路追踪（Jaeger）
- [ ] 自定义告警
- [ ] gRPC
- [ ] pprof
- [ ] CSRF
- [ ] 测试框架
- [ ] Swagger
- [ ] 消息队列
- [ ] 定时任务
- [ ] Docker
- [ ] docker-compose

说明文档：

```go
// 查看 CPU 信息
go tool pprof 127.0.0.1:9999/debug/pprof/profile
...
(pprof) 

//输入 web，生成 svg 文件。
//输入 png，生成 png 文件。
//输入 top，查看排名前 20 的信息。
//查看更多命令，请执行 pprof help。
```

其他同理，比如：

```go
// 查看 内存 信息
go tool pprof 127.0.0.1:9999/debug/pprof/heap

// 查看 协程 信息
go tool pprof 127.0.0.1:9999/debug/pprof/goroutine

// 查看 锁 信息
go tool pprof 127.0.0.1:9999/debug/pprof/mutex
```
如果还想查看火焰图，请执行如下命令：

```go
// 1.下载 pprof 工具
go get -u github.com/google/pprof

// 2.启动可视化界面
pprof -http=:9998 xxx.cpu.prof

// 3.查看可视化界面
http://127.0.0.1:9998/ui/
```