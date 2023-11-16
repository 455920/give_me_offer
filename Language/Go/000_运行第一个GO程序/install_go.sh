# 1、执行脚本前置操作
# sudo vi /etc/yum.repos.d/go.repo
# 文件中添加如下内容
# [go]
# name=Go Repository
# baseurl=https://download.go.dev/goroot/x86_64
# enabled=1
# gpgcheck=0

# 2、运行脚本 执行sh install_go.sh

sudo yum update

# 安装go编译器
sudo yum install golang

# 查看go版本
go version