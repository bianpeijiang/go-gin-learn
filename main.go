package main

import (
	"fmt"
	"github.com/bianpeijiang/go-gin-learn/pkg/setting"
	"github.com/bianpeijiang/go-gin-learn/routers"
	"net/http"
)

func main() {
	/*router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})*/
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
