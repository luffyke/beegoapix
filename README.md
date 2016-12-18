## goxapi
goxapi is beego extension api framework, to development more faster api service.

## Install
```
go get github.com/luffyke/goxapi
```

## Design
### BaseController
1. accepte all client http request, reflect and call sub-controller to handle request.
2. log request and response
3. error handle

### apiRequest

### apiResponse

## Demo
#### new api project
```
bee new helloworld
```

#### edit router.go
```
package routers

import (
	"github.com/luffyke/goxapi"
)

func init() {
	goxapi.Router()
	goxapi.RegController("app", controllers.AppController{})
}
```

#### write your business controller
```
package controllers

import (
	"github.com/luffyke/goxapi/api"

	"github.com/astaxie/beego/logs"
)

type AppController struct {
}

func (this *AppController) CheckVersion(request api.ApiRequest) api.ApiResponse {
	logs.Debug(request.Id)
	logs.Debug(request.Data["versionCode"])
	var response api.ApiResponse
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
  	"size":10,
  	"totalSize":20
  },
  "user":{
    "uid":123,
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