package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.StaticFS("/", http.Dir("web/build"))
	r.Run(":8080")
}
