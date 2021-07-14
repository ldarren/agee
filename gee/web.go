package gee

import (
	"fmt"
	"net/http"
	"path"
	"os"
)

type Web struct {
	*Pipeline
}

func NewWeb() *Web {
	return &Web{
		Pipeline: &Pipeline{
			prefix: "",
			middlewares: make([]HandleFunc, 0),
			router: newRouter(),
			children: make([]*Pipeline, 0),
		},
	}
}

func (web *Web) SetupFileSvrHandler(p *Pipeline, prefix string, local string) HandleFunc {
	abs := path.Join(p.prefix, prefix)
	fileServer := http.StripPrefix(abs, http.FileServer(http.Dir(local)))

	return func (ctx *Context) {
		fpath := ctx.GetParamValue("fpath")
		if _, err := os.Stat(fpath); nil != err {
			ctx.SetStatus(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(ctx.Res, ctx.Req)
	}
}

func (web *Web) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	p := req.URL.Path
	handler, params := web.router.get(req.Method, p)
	if nil == handler {
		fmt.Fprintf(res, "404 Not Found: %s\n", req.URL)
		return
	}
	middlewares := make([]HandleFunc, 0)
	middlewares = append(middlewares, web.middlewares...)
	web.getMiddlewares(p, web.children, middlewares)
	middlewares = append(middlewares, handler)
	ctx := newContext(req, res, params, middlewares)
	ctx.Next()
}

func (web *Web) Start(port string) (err error) {
	return http.ListenAndServe(port, web)
}
