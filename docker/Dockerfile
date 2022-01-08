FROM golang:alpine as builder

WORKDIR /pro

# 把当前所有文件 拷贝到上面的工作目录下（包括配置文件）
COPY tmp /pro

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 设置go代理，加快拉包速度
RUN go env -w GOPROXY=https://goproxy.io,direct && \
    # 拉项目依赖
    go mod tidy && \
    go mod vendor

# 编译程序
RUN go build -o app .

##
## deploy
##

FROM alpine:3.7

# 第一行的as build，build是一个名称，这里使用
COPY --from=builder /pro /pro

RUN mkdir -p /pro/log

RUN chmod 777 /pro/log

WORKDIR /pro

CMD ["./app"]
