package beegoapix

import (
	"github.com/astaxie/beego"
)

const (
	VERSION = "0.1.0"
)

func Router() {
	beego.Router("/:version/:controller/:method", &BaseController{}, "*:Post")
}
