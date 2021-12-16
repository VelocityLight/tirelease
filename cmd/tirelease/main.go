package main

import (
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// absolute path for website
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// start gin website and listen to port
	r := gin.Default()
	r.StaticFS("/", http.Dir(dir+"/website/build/"))
	r.Run(":8080")
}
