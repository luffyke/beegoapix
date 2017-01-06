## beegoapix
beegoapix is beego api extension framework, to develop more faster api service.

## Install
```
go get github.com/luffyke/beegoapix
```

## Function
1. accept all client http request, reflect and call sub-controller to handle request(v0.1)
2. log request and response(v0.1)
3. error handling(v0.1)
4. [version control(v0.2)](https://github.com/luffyke/beegoapix/wiki/API-version-control)
5. priviledge(v0.2)
6. combine controller
7. cache(etag)

## Demo
make sure have installed beego and bee tool.
#### new api project
```
bee api hello
```

#### edit router.go
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

#### write your business controller
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

#### run the server
```
bee run
```

#### post the request
```
http://localhost:8080/v1/app/check-version
```

#### request
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

##### response
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

## Android demo
please reference to project [beegoapix-android-demo](https://github.com/luffyke/beegoapix-android-demo)