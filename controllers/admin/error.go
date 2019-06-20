package admin

type ErrorController struct {
	baseController
}

func (c *ErrorController) Error404()  {
	c.Data["content"] = "page not found"
	c.TplName = "admin/common/404.html"
}
