package main

import (
	//"fmt"
	"gee"
	"log"
	"time"
)

func Logger() gee.HandleFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("%s in %v", c.Req.RequestURI, time.Since(t))
	}
}

func onlyForV2() gee.HandleFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		c.Next()
		// Calculate resolution time
		log.Printf("%s in %v for group v2", c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	web := gee.NewWeb()
	web.Use(Logger())
	web.GET("/", handleRoot)
	v1 := gee.NewGroup(web.Pipeline, "/v1")
	v1.Use(onlyForV2())
	v1.GET("/trigger", handleTrigger)
	v1.GET("/:name", handleParams)
	v1.Static("/log", web.SetupFileSvrHandler(v1, "/log", "./pub"))
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
