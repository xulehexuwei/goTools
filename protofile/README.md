准备工作

Mac 安装 protoc 编译器
$ brew install protobuf

安装 protoc 编译器插件
$ go get -u github.com/golang/protobuf/protoc-gen-go

定义服务
hello.proto

生成proto go 包
需要之前安装的 protoc 以及 protoc-gen-go 工具了
$ protoc --go_out=plugins=grpc:. helloworld.proto
--go_out 表示指定最终生成文件的输出路径，. 表示生成在当前路径下

## 注意这边如果使用 protoc --go_out=grpc:. hello.proto 命令生成的结果和指定了 plugins 生成的结果有不同，需要加上 plugins 去生成

hello.proto 需有下面的参数，代表生成的go文件的位置
option go_package = "../proto";
