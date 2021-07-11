package gee

import (
	"net/http"
	"encoding/json"
)

type Object map[string]interface{}

type Context struct{
	Req *http.Request
	Res http.ResponseWriter
}

func newContext(req *http.Request, res http.ResponseWriter) *Context{
	return &Context{Req: req, Res: res}
}

func (ctx *Context) GetFormValue(key string) string {
	return ctx.Req.FormValue(key)
}

func (ctx *Context) GetQueryValue(key string) string {
	return ctx.Req.URL.Query().Get(key)
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
