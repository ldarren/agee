package main

import (
	//"fmt"
	"gee"
	"log"
)

func main() {
	web := gee.NewWeb()
	v1 := gee.NewRouterGroup(web, "/v1")
	v1.GET("/", handleRoot)
	v1.GET("/trigger", handleTrigger)
	v1.GET("/:name", handleParams)
	log.Fatal(web.Start(":8080"))
}

func handleRoot(ctx *gee.Context) {
	ctx.JSON(200, gee.Object{
		"path": ctx.Req.URL.Path,
	})
	//fmt.Fprintf(res, "From %q\n", ctx.Req.URL.Path)
}

func handleTrigger(ctx *gee.Context) {
	obj := make(gee.Object)
	for k, v := range ctx.Req.Header {
		obj[k] = v
		//fmt.Fprintf(res, "Key[%q] [%q]\n", k, v)
	}
	ctx.JSON(200, obj)
}

func handleParams(ctx *gee.Context) {
	ctx.JSON(200, gee.Object{
		"name": ctx.GetParamValue("name"),
	})
}
