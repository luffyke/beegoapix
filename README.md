## goxapi
goxapi is beego extension api framework, to develop more faster api service.

## Install
```
go get github.com/luffyke/goxapi
```

## Design
### BaseController(base.go)
1. accept all client http request, reflect and call sub-controller to handle request(v0.1)
2. log request and response(v0.1)
3. error handling(v0.1)
4. version control(v0.2)
5. combine controller
6. cache(etag)
7. priviledge

## Demo
#### new api project
```
bee api helloworld
```

#### edit router.go
```
package routers

import (
	"github.com/luffyke/goxapi"
)

func init() {
	goxapi.Router()
	// add your business path mapping
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

## Api version control
router.go, combine controller name(app) with version(v2) to a new controller(AppV2Controller)
```
goxapi.RegController("appv2", controllers.AppV2Controller{})
```

```
type AppV2Controller struct {
}
func (this *AppV2Controller) CheckVersion(request api.ApiRequest) api.ApiResponse {
	logs.Debug(request.Id)
	logs.Debug(request.Data["versionCode"])
	var response api.ApiResponse
	response.Data = make(map[string]interface{})
	response.Data["versionName"] = "version name 2.0"
	return response
}
```

post request
```
http://localhost:8080/v2/app/check-version
```