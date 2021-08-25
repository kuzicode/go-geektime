
##  Week04 作业题目：

1. 按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。


## Ask

```
.
├── api                       // api pkg
│   └── user
│       └── v1
│           ├── user.xx1.go
│           └── user.xx2.go
├── bin
│   └── app                   // bin file
├── cmd                       // main dir, compile here
│   ├── myapp1
│   │   ├── base.go
│   │   └── main.go
│   └── myapp2
│       └── main.go
├── go.mod
├── internal                  // private pkg
│   ├── demo
│   │   ├── biz
│   │   │   ├── biz.go
│   │   │   └── user.go
│   │   ├── data
│   │   │   └── user.go
│   │   ├── pkg
│   │   │   └── chainTool
│   │   └── service
│   └── pkg
├── Makefile
├── pkg                       // public pkg
│   ├── cache
│   │   ├── memcache
│   │   └── redis
│   └── conf
│       ├── dsn
│       └── env
├── README.md
└── test                      // test module
    └── test.go
```
