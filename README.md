[中文文档](./README.zh-cn.md)

## beegoapix
beegoapix is beego api extension framework, to develop more faster api service.

## Install
```
go get github.com/luffyke/beegoapix
```

## Function
1. Http API handling, accept all client http request, reflect and call sub-controller to handle request(v0.1)
2. Logging, log request and response(v0.1)
3. Error handling(v0.1)
4. [API version control(v0.2)](https://github.com/luffyke/beegoapix/wiki/API-version-control)
5. API authority control(v0.2)
6. Combine controller
7. Cache(etag)

## Demo
#### Setup environment
```bash
go get github.com/astaxie/beego
go get github.com/luffyke/beegoapix
go get github.com/beego/bee
```

#### New api project
```
bee api hello
```

#### Edit router.go
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

#### Business controller
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

#### Run the server
```
bee run
```

#### Post the request
```
http://localhost:8080/v1/app/check-version
```

#### Request
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

##### Response
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

## Deploy
1. setup go environment
2. install beego and beegoapix
```
go get github.com/astaxie/beego
go get github.com/luffyke/beegoapix
```
3. download and run [beego deploy shell](https://gist.github.com/luffyke/790154ec5142abd9fd6245a5fd8b9427)

## Android demo
please reference to project [beegoapix-android-demo](https://github.com/luffyke/beegoapix-android-demo)