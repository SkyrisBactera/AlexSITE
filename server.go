package main

import (
	"fmt"

	"github.com/NYTimes/gziphandler"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/turtlemonvh/gin-wraphh"
	//"os"
)

var (
	//Domain cannot be an IP Address, unless you are willing to sacrifice HTTPS
	domain       = "localhost"
	subdirectory = ""
)

func main() {
	if domain == "" || domain == "localhost" {
		fmt.Println("Using HTTP only")
		http := gin.Default()
		http.Static("/alex", "./public")
		http.Use(wraphh.WrapHH(gziphandler.GzipHandler))
		http.Run(":80")
	} else {
		fmt.Println("Using HTTPS only")
		http := gin.Default()
		https := gin.Default()
		https.Static("/alex", "./public")
		http.Use(wraphh.WrapHH(gziphandler.GzipHandler))
		https.Use(wraphh.WrapHH(gziphandler.GzipHandler))
		http.GET("/*path", func(c *gin.Context) {
			c.Redirect(302, "https://"+domain+subdirectory+c.Param("variable"))
		})
		go autotls.Run(https, domain)
		http.Run(":80")
	}
}
