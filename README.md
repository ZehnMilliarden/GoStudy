# GoStudy

## 开发环境配置

### 安装包下载
	1) https://golang.google.cn/dl/
	2) https://golang.org/dl/
	
### 环境变量配置
	go env , 查看环境变量配置

#### GOPROXY
	GOPROXY=Go镜像代理, 设置命令 go env -w GOPROXY=https://goproxy.cn,direct

#### GO111MODULE
	GO111MODULE=off，go命令行将不会支持module功能，寻找依赖包的方式将会沿用旧版本那种通过vendor目录或者GOPATH模式来查找。
	GO111MODULE=on，go命令行会使用modules，而一点也不会去GOPATH目录下查找。
	GO111MODULE=auto，默认值，go命令行将会根据当前目录来决定是否启用module功能。这种情况下可以分为两种情形：当前目录在GOPATH/src之外且该目录包含go.mod文件，或者当前文件在包含go.mod文件的目录下面。当module功能启用时，GOPATH在项目构建过程中不再担当import的角色，但它仍然存储下载的依赖包，具体位置在$GOPATH/pkg/mod。
	设置方法
		1) go env -w GO111MODULE=xxx
		2) vscode 配置文件中设置 settings.json -> "go.toolsEnvVars": { "GO111MODULE": "xxx"} , 注意其他环境变量也可以参照设置
	
### VSCODE配置
	1) GO基础扩展
	2) VS命令行 (ctrl+shift+p) , 搜索 Go Install/Update Tools 更新下载所有选项

### Go.Mod
	1) 在工程的源码路径处 执行 go mod init xxx(模块名)，初始化模块
	2) go mod tidy ，加载依赖模块


### Go.Work
	示例
	go work init .\src\main .\src\example
	