
### 基于 mq 的 API 网关 / API Gateway
---
主要特性 / main features:
- golang
- api-gateway
- mq
- json_rpc(by rabbitmq)
- micro_service


配置文件 / config
```cfg
# config.cfg
[gateway]
listen=0.0.0.0:9000
debug=true
signKey=SIGN_KEY
timeout=20

[mq]
url=amqp://test:pw123@127.0.0.1:5672/

```

快速安装RabbitMQ / quickly install rabbitmq

```bash
docker run -d --hostname rabbit_host --name mq -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=pw123 -p 15672:15672 -p 5672:5672 rabbitmq:management
```
启动代码 / run

```bash
# 启动网关
go run gateway/main.go

# 启动微服务端
go run micro_service/main.go

```



测试连接 / test
```
http://127.0.0.1/gateway/?module=examples&version=1.0&method=Examples.Echo&bizContent={"Body":"6666"}
```