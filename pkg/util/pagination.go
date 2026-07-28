package util

import (
	"github.com/bianpeijiang/go-gin-learn/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
)

func GetPage(c *gin.Context) int {
	result := 0
	page := 1
	if getPage := c.Query("page"); getPage != "" {
		var err error
		page, err = com.StrTo(getPage).Int()
		if err != nil {
			log.Fatalf("GetPage err: %v", err)
		}
	}

	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
