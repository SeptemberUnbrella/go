package main

import (
	"github.com/go-programming-tour-book/blog-service/internal/routers"
	"net/http"
	"time"
)

func main() {
	//r := gin.Default()
	//r.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"message": "pong"})
	//})
	//r.Run()
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8008",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}
