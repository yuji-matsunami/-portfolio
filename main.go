package main

import (
	"net/http"

	"github.com/gin-gonic/gin")


func main() {
	server := gin.Default()
	server.Static("/static", "./static")
	server.StaticFS("/templates", http.Dir("templates"))
	server.GET("/", func(c *gin.Context) {
		c.Redirect(302, "templates/main.tmpl")
	})
	server.Run()
}