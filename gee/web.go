package gee

import (
	"fmt"
	"net/http"
	"html/template"
	"strings"
	"path"
	"os"
)

type Web struct {
	*Pipeline
	engine *template.Template
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

func (web *Web) SetupFileSvrHandler(key string, local string) HandleFunc {
	fileServer := http.FileServer(http.Dir(local))

	return func (ctx *Context) {
		fpath := ctx.GetParamValue(key)
		if _, err := os.Stat(path.Join(local, fpath)); nil != err {
			ctx.Fail(http.StatusNotFound, err.Error())
			return
		}

		stripped := http.StripPrefix(ctx.prefix, fileServer)
		stripped.ServeHTTP(ctx.Res, ctx.Req)
	}
}

func (web *Web) SetupTemplateEngine(pattern string, FormatAsDate interface{}) {
	funcMap := template.FuncMap{
		"FormatAsDate": FormatAsDate,
	}
	web.engine = template.Must(template.New("").Funcs(funcMap).ParseGlob(pattern))
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
	prefix := web.getMiddlewares(p, web.children, middlewares)
	middlewares = append(middlewares, handler)
	ctx := newContext(req, res, prefix, params, middlewares)
	ctx.Next()
}

func (web *Web) HTML(name string, data interface{}) HandleFunc  {
	engine := web.engine
	return func (ctx *Context) {
		url := ctx.GetParamValue("url")
		if strings.HasSuffix(url, "/") {
			url = url + name
		} else {
			url = strings.Replace(url, "html", "tmpl", -1)
		}
		ctx.SetHeader("Content-Type", "text/html")
		ctx.SetStatus(http.StatusOK)
		if err := engine.ExecuteTemplate(ctx.Res, name, data); err != nil {
			ctx.Fail(500, err.Error())
		}
	}
}

func (web *Web) Start(port string) (err error) {
	return http.ListenAndServe(port, web)
}
