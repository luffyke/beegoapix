package controllers

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/luffyke/goxapi/models/api"
)

type BaseController struct {
	beego.Controller
}

var regControllers map[string]interface{} = make(map[string]interface{})

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
		// valid request

		// get Data

		// get controller and get method
		controller, method := this.Ctx.Input.Param(":controller"), this.Ctx.Input.Param(":method")
		controllerName := regControllers[controller]
		if controllerName == nil {
			logs.Error("controller is not registered:", controller)
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
		}
	}
	this.Data["json"] = response
	this.ServeJSON()
}

func RegController(name string, controller beego.Controller) {
	regControllers[name] = controller
}

func formatMethod(method string) string {
	strs := strings.Split(method, "-")
	var result string
	for _, s := range strs {
		result = result + strings.ToUpper(s[0:1]) + s[1:len(s)]
	}
	return result
}
