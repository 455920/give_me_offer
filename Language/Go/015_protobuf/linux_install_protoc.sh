# 更新系统：打开终端，并使用以下命令更新系统软件包：
sudo yum update

# 安装依赖项：使用以下命令安装构建和编译工具所需的依赖项：
sudo yum install -y autoconf automake libtool curl make glibc-devel gcc-c++

# 下载编译器：你可以从Protocol Buffers的GitHub存储库下载编译器。使用以下命令下载最新版本的源代码：
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.18.0/protobuf-all-3.18.0.tar.gz

# 解压缩文件：使用以下命令解压缩下载的压缩文件：
tar -xvzf protobuf-all-3.18.0.tar.gz

# 进入目录：进入解压缩后的目录：
cd protobuf-3.18.0

# 配置和编译：使用以下命令配置和编译编译器
./configure
make

# 安装编译器：使用以下命令将编译器安装到系统中：
sudo make install

# 验证安装：运行以下命令验证安装是否成功：
protoc --version