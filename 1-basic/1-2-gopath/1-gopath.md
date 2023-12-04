### GOPATH：
```text
存放源码的路径，go文件需要放在gopath的src目录下才能够编译运行，
因此不要设置全局gopath目录，会难以管理所有的golang项目。并且第
三方的依赖包也会拉取到src目录下，造成代码管理混乱。

gopath目录下文件
go
├── bin 编译生成的二进制文件
├── pkg 
|   ├── XX_amd64 XX是目标操作系统，mac->darwin_amd64, linux\->linux_amd64，存放的是.a结尾的文件。
|   ├── mod 当开启go Modules 模式下，go get命令缓存下依赖包存放的位置
|   └── sumdb go get命令缓存下载的checksum数据存放的位置
└── src 存放golang项目代码的位置
    ├── github.com
    ├── golang.org
    └── google.golang.org

缺点：
    1.gopath的配置导致项目需要存放在指定的目录下，不能随便移动
    2.go get 命令的时候，无法指定获取的版本
    3.引用第三方项目的时候，无法处理v1、v2、v3等不同版本的引用问题，
      因为在GOPATH 模式下项目路径都是 github.com/foo/project
    4.无法同步一致第三方版本号，在运行 Go 应用程序的时候，无法保证其
      它人与所期望依赖的第三方库是相同的版本。
      
针对以上问题，golang 推出了 goModule
```

### GOMODULE:
```text
go文件可以放在gomodule目录下而不需要放在src目录下也可以编译，在
引入了gomodule之后，gopath主要负责存放第三方依赖包，gomodule存
放自己的源码文件。当项目需要第三方依赖包，通过gomodule目录的
go.mod文件来引用gopath目录src包的第三方依赖即可，也就是说，
go1.11之后，gopath可以作为一个本地依赖包仓库使用.

在go 1.11 官方出手了推出了 Go Modules， 通过设置环境变量 GO111MODULE 进行开启或者关闭 go mod 模式
```
#### goModule的配置姿势：
**1.IDE** ide工具创建golang项目可配置，自行搜索网络资料

**2.命令行** 没什么必要，
```shell
# test_mod 为项目名称
go mod int test_mod

# 引入第三方包
go get github.com/robfig/cron/v3

# 删除没有引用的第三方包
go mod tidy
```

**go.mod**
```text
module go-practice

go 1.20

require github.com/robfig/cron/v3 v3.0.1 // indirect
```

**查看gopath**
```shell
go env GOPATH
```