package main

import (
	"flag"
	"github.com/curt-labs/bencher/controllers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	port = flag.String("port", "8000", "Port for the application to start on")
)

func main() {
	flag.Parse()

	r := gin.Default()
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	r.LoadHTMLTemplates("templates/*")
	r.ServeFiles("/public/*filepath", http.Dir(wd+"/public"))
	r.GET("/run", bencher.Run)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusFound, "layout.html", nil)
	})

	s := &http.Server{
		Addr:           ":" + *port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Started on :%s\n", *port)
	s.ListenAndServe()

}
