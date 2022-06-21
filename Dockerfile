# 编译镜像
FROM harbor.devops.hpay/base/golang:1.17-alpine AS build

WORKDIR /project/
COPY ./  /project
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go build -o /project/work/k8s-node-manager cmd/main.go

# 运行镜像
FROM harbor.devops.hpay/base/golang:alpine
ENV TZ Asia/Shanghai
RUN apk add tzdata \
    &&  cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && apk del tzdata
COPY ./config.yaml /work/
COPY --from=build /project/work/ /work/

# 定义工作目录为work
WORKDIR /work
# 开放http 9999端口
EXPOSE 9999
# 启动http服务
ENTRYPOINT ["./k8s-node-manager"]