package gee

import (
	"encoding/json"
	"net/http"
)

type Object map[string]interface{}

type Context struct {
	Req    *http.Request
	Res    http.ResponseWriter
	prefix string
	Params map[string]string
	handlers []HandleFunc
	index int
}

func newContext(
	req *http.Request,
	res http.ResponseWriter,
	prefix string,
	params map[string]string,
	handlers []HandleFunc) *Context {
	return &Context{
		Req: req,
		Res: res,
		prefix: prefix,
		Params: params,
		handlers: handlers,
		index: -1,
	}
}

func (ctx *Context) GetFormValue(key string) string {
	return ctx.Req.FormValue(key)
}

func (ctx *Context) GetQueryValue(key string) string {
	return ctx.Req.URL.Query().Get(key)
}

func (ctx *Context) GetParamValue(key string) string {
	val, _ := ctx.Params[key]
	return val
}

func (ctx *Context) SetHeader(key string, val string) {
	ctx.Res.Header().Set(key, val)
}

func (ctx *Context) SetStatus(code int) {
	ctx.Res.WriteHeader(code)
}

func (ctx *Context) JSON(code int, obj interface{}) {
	ctx.SetHeader("Context-Type", "application/json")
	ctx.SetStatus(code)
	encoder := json.NewEncoder(ctx.Res)
	if err := encoder.Encode(obj); err != nil {
		http.Error(ctx.Res, err.Error(), 500)
	}
}

func (ctx *Context) Fail(code int, err string) {
	ctx.index = len(ctx.handlers)
	ctx.SetStatus(code)
	ctx.JSON(code, Object{"message": err})
}

func (ctx *Context) Next() {
	ctx.index++

	max := len(ctx.handlers)

	for ; ctx.index < max; ctx.index++ {
		ctx.handlers[ctx.index](ctx)
	}
}
