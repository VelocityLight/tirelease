package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.StaticFS("/", http.Dir("./website/build"))
	r.Run(":8080")
}
