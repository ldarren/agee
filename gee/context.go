package gee

import (
	"net/http"
	"encoding/json"
)

type Object map[string]interface{}

type Context struct{
	Req *http.Request
	Res http.ResponseWriter
	Params map[string]string
}

func newContext(req *http.Request, res http.ResponseWriter, params map[string]string) *Context{
	return &Context{Req: req, Res: res, Params: params}
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

func (ctx *Context) SetHeader(key string, val string){
	ctx.Res.Header().Set(key, val)
}

func (ctx *Context) SetStatus(code int){
	ctx.Res.WriteHeader(code)
}

func (ctx *Context) JSON(code int, obj interface{}){
	ctx.SetHeader("Context-Type", "application/json")
	ctx.SetStatus(code)
	encoder := json.NewEncoder(ctx.Res)
	if err := encoder.Encode(obj); err != nil {
		http.Error(ctx.Res, err.Error(), 500)
	}
}
