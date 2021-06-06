# go-admin 通用后台系统

#### [go-admin](https://github.com/big-pang/go-admin) v2.0.1版本，基于beego 2.0.1 框架和AdminLte前端框架，开发的go语言通用后台系统，在beego框架的基础上，封装了后台系统的分页功能，excel数据导出功能等丰富常用的扩展，基于MVC模式，html界面随心定义，相较于某些后台复杂代码生成的前端html元素，使用原生的html原生作为前端显示，更加的灵活自由。[go-admin](https://github.com/big-pang/go-admin)通用的后台系统真正的做到了开箱即用，欢迎大家使用。

## go-admin 安装

### 安装方式 (GO MOD方式安装,已移除 GOPATH方式安装说明，需要的请查看 tag v1.0.1)

#### 1、安装beego v2.0.1和bee v2.0.2
参考[Beego](https://beego.me/docs/install/) 和 [Bee](https://beego.me/docs/install/bee.md) 安装手册

#### 2、clone 项目到本地 GOPATH src目录之外的路径下
[**github**](https://github.com:big-pang/go-admin)
```
git clone git@github.com:big-pang/go-admin.git
```


#### 3、配置数据库
将目录中go-admin.sql文件导入mysql数据库
拷贝根目录下的config.yaml.example为config.yaml,修改文件内的数据库连接信息
```
cp config.yaml.example config.yaml
```

#### 4、安装项目依赖
```
go-admin目录下 go mod tidy 将自动下载依赖包
```

### 通过上面方式安装后,接下来

#### 运行系统
``` go
//在项目下运行，开始进行调试开发
bee run
```

#### 访问后台
访问`/admin/index/index`，默认超级管理员的账号密码都为`super_admin`。


## 补充
[go-admin](https://github.com/big-pang/go-admin) 项目在beego v2.0.1的框架基础上完善了很多丰富的常用后台功能，分页封装、excel数据一键导出等功能。
此外，如果有需要php语言的laravel版本的后台管理系统，可以使用[laravel-admin](https://github.com/big-pang/laravel-admin)。


#### [go-admin](https://github.com/big-pang/go-admin)通用后台系统效果图

