Go Modules是Go语言自1.11版本引入的官方包管理工具，用于管理和版本控制Go项目的依赖关系。它提供了一种更现代化、灵活和可靠的方式来管理包，取代了之前的GOPATH方式。

以下是Go Modules的基本用法：

初始化模块：在你的项目根目录下，使用命令go mod init来初始化一个新的模块。例如，go mod init example.com/myproject。这将创建一个go.mod文件，用于记录项目的依赖关系和版本信息。

添加依赖：使用go get命令来添加依赖项。例如，go get github.com/gorilla/mux。这将下载并安装gorilla/mux包，并将其添加到go.mod文件中。

构建和运行：使用go build或go run命令来构建和运行你的项目。Go Modules会自动处理依赖项的下载和版本管理。

版本控制：Go Modules使用语义化版本控制（Semantic Versioning）来管理依赖项的版本。你可以使用go get命令来指定特定的版本，例如，go get github.com/gorilla/mux@v1.8.0。

更新依赖：使用go get -u命令来更新依赖项到最新版本。例如，go get -u github.com/gorilla/mux。

清理依赖：使用go mod tidy命令来清理不再使用的依赖项。它会根据你的代码和go.mod文件中的实际使用情况，移除不需要的依赖项。

Vendor目录：Go Modules还支持Vendor模式，可以将依赖项复制到项目的vendor目录中，以便离线构建和版本控制。可以使用go mod vendor命令来生成vendor目录。

需要注意的是，Go Modules不再依赖于GOPATH，你可以在任意目录下进行Go开发。Go Modules会自动管理依赖项的下载和版本控制，无需手动设置GOPATH。

- 1、go mod init 初始化模块，会在根项目下执行后，会在的根目录下生成一个go.mod来管理项目的依赖
- 2、go.mod会自动识别依赖关系并且自动设置gopath目录为go.mod所在目录, 只需要go build即可
- 3、go get xxx 添加依赖项，会添加到go.mod中
- 4、在Go语言中，一个目录下只能定义一个包。这是Go语言的约定和规范。
- 5、目录名和包名最好一致