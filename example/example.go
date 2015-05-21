package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/iyee/ginpongo2"
)

func main() {
	r := gin.Default()
	r.Use(ginpongo2.Pongo2())

	r.GET("/", func(c *gin.Context) {
		c.Set("template", "index.html")
		c.Set("data", map[string]interface{}{"message": "Hello World!"})
	})

	r.GET("/none", func(c *gin.Context) {
		c.Set("template", "none.html")
	})

	r.GET("/invalidTemplate", func(c *gin.Context) {
		c.Set("template", 3)
	})

	r.GET("/invalidData", func(c *gin.Context) {
		c.Set("template", "index.html")
		c.Set("data", 3)
	})

	r.GET("/emptyData", func(c *gin.Context) {
		c.Set("template", "index.html")
		c.Set("data", nil)
	})

	log.Println(r.Run(":8080"))
}
