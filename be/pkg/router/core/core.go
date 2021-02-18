package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandlerFunc core封装后的handler
type HandlerFunc func(c *Context)

// Handle 将core.HandlerFunc转换为gin.HandlerFunc
func Handle(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{
			c,
		}
		h(ctx)
	}
}

// Context core封装后的Context
type Context struct {
	*gin.Context
}

const (
	// CodeOK 表示响应成功
	CodeOK = 0
	// CodeErr 表示响应失败
	CodeErr = 1
)

// Res 标准JSON输出格式
type Res struct {
	// Code 响应的业务错误码。0表示业务执行成功，非0表示业务执行失败。
	Code int `json:"code"`
	// Msg 响应的参考消息。前端可使用msg来做提示
	Msg string `json:"msg"`
	// Data 响应的具体数据
	Data interface{} `json:"data"`
}

// ResPage 带分页的标准JSON输出格式
type ResPage struct {
	Res
	// Current 总记录数
	Current int `json:"current"`
	// PageSize 每页记录数
	PageSize int `json:"pageSize"`
	// Total 总页数
	Total int `json:"total"`
}

// JSON 输出响应JSON
// 形如 {"code":<code>, "msg":<msg>, "data":<data>}
func (c *Context) JSON(httpStatus int, res Res) {
	c.Context.JSON(httpStatus, res)
}

// JSONOK 输出响应成功JSON，如果data不为零值，则输出data
// 形如 {"code":0, "msg":"成功", "data":<data>}
func (c *Context) JSONOK(data ...interface{}) {
	j := new(Res)
	j.Code = CodeOK
	j.Msg = "成功"
	if len(data) > 0 {
		j.Data = data[0]
	} else {
		j.Data = ""
	}
	c.Context.JSON(http.StatusOK, j)
	return
}

// JSONE 输出失败响应
func (c *Context) JSONE(code int, msg string, data error) {
	j := new(Res)
	j.Code = code
	j.Msg = msg
	//if data != nil {
	//	j.Data = data.Error()
	//}
	c.Context.JSON(http.StatusOK, j)
	return
}

// JSONPage 输出带分页信息的JSON
func (c *Context) JSONPage(data interface{}, current, pageSize, total int) {
	j := new(ResPage)
	j.Code = CodeOK
	j.Data = data
	j.Current = current
	j.PageSize = pageSize
	j.Total = total
	c.Context.JSON(http.StatusOK, j)
}
