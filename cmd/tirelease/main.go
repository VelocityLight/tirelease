package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// start gin website and listen to port
	r := gin.Default()
	r.StaticFS("/", http.Dir("website/build/"))
	r.Run(":8080")
}
