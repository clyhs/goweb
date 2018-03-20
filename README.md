# helloweb
* 用GO写的一个WEB应用。
### 数据库
* 数据库采用mysql
### 用到的第三方技术
* xorm
* goconfig
### 编译
clean:
	@rm -rf controller/controller

build:
	@cd controller && godep go build -a -tags "netgo static_build" -installsuffix netgo
### 启动
* 默认端口 8070
* 启动命令 ./controller server