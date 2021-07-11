package main

import (
	//"fmt"
	"log"
	"gee"
)

func main(){
	router := gee.NewRouter()
	router.GET("/", handleRoot)
	router.GET("/trigger", handleTrigger)
	log.Fatal(router.Start(":8080"))
}

func handleRoot(ctx *gee.Context){
	ctx.JSON(200, gee.Object{
		"path": ctx.Req.URL.Path,
	})
	//fmt.Fprintf(res, "From %q\n", ctx.Req.URL.Path)
}

func handleTrigger(ctx *gee.Context){
	obj := make(gee.Object)
	for k, v := range ctx.Req.Header {
		obj[k] = v
		//fmt.Fprintf(res, "Key[%q] [%q]\n", k, v)
	}
	ctx.JSON(200, obj)
}
