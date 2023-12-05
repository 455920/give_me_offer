go mod init example

# 安装protobuf插件：在终端中运行以下命令来安装Go语言的protobuf插件：
go get google.golang.org/protobuf/cmd/protoc-gen-go

# 安装protobuf库：运行以下命令来安装Go语言的protobuf库
go get google.golang.org/protobuf/cmd/protoc-gen-go


# 验证安装：在终端中运行以下命令来验证安装是否成功
protoc-gen-go --version