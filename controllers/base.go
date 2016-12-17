package controllers

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"manu-number-server/models/api"
)

var regControllers map[string]interface{}

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Post() {
	logs.Info("request:", string(this.Ctx.Input.RequestBody))
	var request api.ApiRequest
	var response api.ApiResponse
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &request)
	if err != nil {
		logs.Error("json error:", err)
		response.State = *api.JsonError
	} else {
		response.Id = request.Id
		// 验证request

		// 获取Data

		// 获取controller和method
		controller, method := this.Ctx.Input.Param(":controller"), this.Ctx.Input.Param(":method")
		controllerName := regControllers[controller]
		if controllerName == nil {
			logs.Error("controller not found:", controller)
			response.State = *api.Error
		} else {
			method = formatMethod(method)
			// reflect
			t := reflect.TypeOf(controllerName)
			c := reflect.New(t)
			m := c.MethodByName(method)
			in := make([]reflect.Value, 1)
			in[0] = reflect.ValueOf(request)
			out := make([]reflect.Value, 1)
			out = m.Call(in)
			response = out[0].Interface().(api.ApiResponse)
			//this.Data["json"] = map[string]string{"id": request.Id}
		}
	}
	this.Data["json"] = response
	this.ServeJSON()
}

func init() {
	regControllers = make(map[string]interface{})
	regControllers["app"] = AppController{}
}

func formatMethod(method string) string {
	strs := strings.Split(method, "-")
	var result string
	for _, s := range strs {
		result = result + strings.ToUpper(s[0:1]) + s[1:len(s)]
	}
	return result
}
