package main

import (
	"os"

	"path/filepath"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// absolute path for website
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	// start gin website and listen to port
	r := gin.Default()
	r.StaticFS("/", http.Dir(dir+"/website/build/"))
	r.Run(":8080")
}
