FROM golang:alpine as builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64\
    GOPROXY=https://goproxy.cn

WORKDIR /pro

# 把当前所有文件 拷贝到上面的工作目录下（包括配置文件）
COPY go.mod go.mod
COPY go.sum go.sum

# 把COPY tmp /pro是为了，go.mod go.sum文件内容不变的情况下，不重新下载
RUN go mod download

COPY tmp /pro
RUN go mod tidy && go mod vendor

# 编译程序
RUN go build -o app .

FROM alpine:3.7

# 第一行的as build，build是一个名称，这里使用
COPY --from=builder /pro /pro

RUN mkdir -p /pro/log

RUN chmod 777 /pro/log

WORKDIR /pro

CMD ["./app"]
