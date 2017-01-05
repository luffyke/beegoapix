package beegoapix

import (
	"github.com/astaxie/beego"
)

const (
	VERSION = "0.2.0"
)

func Router() {
	beego.Router("/:version/:controller/:method", &BaseController{}, "*:Post")
}
