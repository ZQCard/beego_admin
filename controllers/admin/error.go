package admin

import (
	"fmt"
	"strings"
)

type ErrorController struct {
	baseController
}

func (c *ErrorController) Error404()  {
	if strings.HasPrefix(c.Ctx.Request.URL.String(), "/admin") {
		fmt.Println(22)
		c.Data["content"] = "page not found"
		c.TplName = "admin/common/404.html"
	}else {
		c.TplName = "company/404.html"
	}

}
