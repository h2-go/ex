package main

import (
  "github.com/gin-gonic/gin"
)

func router() *gin.Engine {
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    c.String(200, "Hello")
  })

  return r
}

func main() {
  router().Run()
}
