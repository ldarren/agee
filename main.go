package main

import (
	"fmt"
	"gee"
	"log"
	"time"
)

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

type student struct {
	Name   string
	Age    int8
	Joined time.Time
}

func main() {
	web := gee.NewWeb()
	web.Use(recovery())
	web.Use(logger())

	stu1 := &student{Name: "Geektutu", Age: 20, Joined: time.Date(2021, 7, 17, 0, 0, 0, 0, time.UTC)}
	stu2 := &student{Name: "Jack", Age: 22, Joined: time.Date(2021, 7, 19, 0, 0, 0, 0, time.UTC)}

	web.SetupTemplateEngine("tmpl/*", FormatAsDate)
	web.GET("/*url", web.HTML("index.tmpl", gee.Object{
		"title":  "Agee",
		"stuArr": [2]*student{stu1, stu2},
	}))

	v1 := gee.NewGroup(web.Pipeline, "/v1")
	v1.Use(auth())
	v1.GET("/trigger", handleTrigger)
	v1.GET("/panic", func(ctx *gee.Context) {
		names := []int{1}
		ctx.JSON(200, gee.Object{
			"name": names[100],
		})
	})
	v1.GET("/:name", handleParams)

	assets := gee.NewGroup(web.Pipeline, "/assets")
	assets.GET("/*fpath", web.SetupFileSvrHandler("fpath", "./pub/upd"))

	log.Fatal(web.Start(":8080"))
}
