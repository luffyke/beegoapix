## 简介
beegoapix是一个beego API开发框架，用于快速开发API服务。

## 安装
```
go get github.com/luffyke/beegoapix
```

## 功能
1. 统一API请求处理和返回(v0.1)
2. 统一日志处理(v0.1)
3. 异常控制(v0.1)
4. [API版本控制(v0.2)](https://github.com/luffyke/beegoapix/wiki/API-version-control)
5. API权限控制(v0.2)
6. 组合接口
7. 接口缓存

## 示例
#### 设置环境
```bash
go get github.com/astaxie/beego
go get github.com/beego/bee
go get github.com/luffyke/beegoapix
```

#### 新建API项目
```
bee api hello
```

#### 编辑router.go，添加api路由
```
package routers

import (
	"hello/controllers"
	"github.com/luffyke/beegoapix"
)

func init() {
	beegoapix.Router()
	// add your business path mapping
	beegoapix.RegController("app", controllers.AppController{})
}
```

#### 实现业务controller
```
package controllers

import (
	"github.com/luffyke/beegoapix/api"

	"github.com/astaxie/beego/logs"
)

type AppController struct {
}

func (this *AppController) CheckVersion(request api.ApiRequest) (response api.ApiResponse) {
	logs.Debug(request.Id)
	logs.Debug(request.Data["versionCode"])
	response.Data = make(map[string]interface{})
	response.Data["versionName"] = "version name 1.0"
	return response
}
```

#### 运行服务
```
bee run
```

#### 测试请求
```
http://localhost:8080/v1/app/check-version
```

#### 请求示例
```
{
  "id":"12345678",
  "sign":"abc",
  "client":{
    "caller":"app",
    "os":"android",
    "ver":"1.0",
    "platform":"android",
    "ch":"offical",
    "ex":{
      "imei":"1a2b3c"
    }
  },
  "page":{
  	"page":1,
  	"size":10
  },
  "user":{
    "uid":"123",
    "sid":"abc"
  },
  "data":{
    "versionCode":"v1.0.0"
  }
}
```

##### 返回
```
{
    "state": {
        "code": 0,
        "msg": ""
    },
    "data": {
        "versionName": "version name 1.0"
    }
}
```

## 部署
1. 设置golang环境
2. 安装 beego 和 beegoapix
```
go get github.com/astaxie/beego
go get github.com/luffyke/beegoapix
```
3. 下载和运行[beego部署脚本](https://gist.github.com/luffyke/790154ec5142abd9fd6245a5fd8b9427)

## Android 示例
参考项目 [beegoapix-android-demo](https://github.com/luffyke/beegoapix-android-demo)