package main

import (
	"fmt"
	"gee"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"
)

func logger() gee.HandleFunc {
	return func(ctx *gee.Context) {
		// Start timer
		t := time.Now()
		// Process request
		ctx.Next()
		// Calculate resolution time
		log.Printf("%s in %v", ctx.Req.RequestURI, time.Since(t))
	}
}

func auth() gee.HandleFunc {
	return func(ctx *gee.Context) {
		log.Printf("%s authenticated", ctx.Req.RequestURI)
	}
}

func handleTrigger(ctx *gee.Context) {
	obj := make(gee.Object)
	obj["path"] = ctx.Req.URL.Path
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

// print stack trace for debug
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func recovery() gee.HandleFunc {
	return func(ctx *gee.Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				ctx.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		ctx.Next()
	}
}
