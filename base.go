package beegoapix

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/luffyke/beegoapix/api"
)

type BaseController struct {
	beego.Controller
}

var regControllers = make(map[string]interface{})
var loginPaths = make(map[string]int)

func (this *BaseController) Post() {
	logs.Info("request:", string(this.Ctx.Input.RequestBody))
	var request api.ApiRequest
	var response api.ApiResponse
	defer func() {
		if err := recover(); err != nil {
			//logs.Debug(reflect.ValueOf(err).Kind())
			if reflect.Struct == reflect.ValueOf(err).Kind() {
				_, ok := err.(api.State)
				if ok {
					response.State = err.(api.State)
				}
			} else {
				logs.Error("server error!", err)
				response.State = api.Error
			}
			this.Data["json"] = response
			this.ServeJSON()
		}
	}()

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &request)
	if err != nil {
		logs.Error("json error:", err)
		//response.State = api.JsonError
		panic(api.JsonError)
	} else {
		response.Id = request.Id
		// valid request

		// valid session
		//logs.Info("url", this.Ctx.Input.URL())
		if _, ok := loginPaths[this.Ctx.Input.URL()]; ok {
			if request.User.Uid == "" || request.User.Sid == "" {
				panic(api.SessionError)
			}
		}
		// get controller and get method
		version, controller, method := this.Ctx.Input.Param(":version"), this.Ctx.Input.Param(":controller"), this.Ctx.Input.Param(":method")
		// get controller
		// default version v1
		if version != "v1" {
			controller = controller + version
		}
		controllerName := regControllers[controller]
		if controllerName == nil {
			logs.Error("controller not registered:", controller)
			//response.State = api.Error
			panic(api.Error)
		} else {
			method = formatMethod(method)
			// reflect
			t := reflect.TypeOf(controllerName)
			c := reflect.New(t)
			m := c.MethodByName(method)
			if !m.IsValid() {
				logs.Error("method not found:", method)
				//response.State = api.Error
				panic(api.Error)
			} else {
				in := make([]reflect.Value, 1)
				in[0] = reflect.ValueOf(request)
				out := make([]reflect.Value, 1)
				out = m.Call(in)
				response = out[0].Interface().(api.ApiResponse)
				//apiError := out[1].Interface().(api.ApiError)
			}
		}
	}
	b, err := json.Marshal(response)
	if err == nil {
		logs.Info("response:", string(b))
	}
	this.Data["json"] = response
	this.ServeJSON()
}

func RegController(name string, controller interface{}) {
	regControllers[name] = controller
}

func SetLoginPaths(paths ...string) {
	for _, path := range paths {
		loginPaths[path] = 1
	}
}

func formatMethod(method string) string {
	strs := strings.Split(method, "-")
	var result string
	for _, s := range strs {
		result = result + strings.ToUpper(s[0:1]) + s[1:len(s)]
	}
	return result
}
