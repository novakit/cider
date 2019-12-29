# Cider

保姆型微服务 Sidecar

## 如何安装

源代码：

`go get -u github.com/novakit/cider`

Docker:

`docker pull novakit/cider`

## 如何使用

Cider 当前只考虑两种部署方式，单机模式 和 Kubernetes 模式

* Cider 接受的参数

    Cider 只接受 `--` 参数，后接要执行的命令
    
* Cider 接受的环境变量

    Cider 会从当前工作目录读取 `cider.env` 文件，作为环境变量的补充，不建议在生产环境使用

    * `PORT` Cider 监听的 HTTP 端口，默认为 `3000`
    * `CIDER_OUTLET_` 开头的环境变量，HTTP 依赖服务聚合配置
        * `tcp` 简单的 TCP 地址；格式 `tcp://HOST:PORT`；示例 `OUTLET_ACCOUNT_MS=/account/=tcp://10.10.10.10:8080`
        * `tcp+k8s` Kubernetes 环境下，使用 Kubernetes 自身的服务发现和健康检查机制；格式 `tcp+k8s://service.namespace:PORT`；示例 `OUTLET_SNOWFLAKE=/snowflake/=tcp+k8s://snowflake.common:sf-port`
    
* Cider 传递给子进程的环境变量

    * `PORT` 随机端口，替换原始值；子进程应该使用该环境变量作为 HTTP 服务监听端口
    * `CIDER_PORT_INLET` 原始 `PORT` 环境变量；一般情况下不需要使用
    * `CIDER_PORT_OUTLET` 随机端口；Cider 对所有依赖服务进行 HTTP 路径聚合，并在该端口提供服务；子进程对外所有调用均应该使用该端口
    * 所有其他环境变量

Docker 示例:

```
FROM novakit/cider AS cider

FROM golang:1.13 AS builder
COPY server.go /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server

FROM scratch
COPY --from cider /cider /
COPY --from builder /server /
CMD ["/cider", "--", "/server"]
```

## 开发者

Guo Y.K. <hi@guoyk.net>
