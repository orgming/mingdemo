package middleware

import (
	"github.com/orgming/ming/framework/gin"
	"log"
	"time"
)

// Cost 请求时长统计
func Cost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		log.Printf("api uri start: %v", c.Request.RequestURI)
		c.Next()
		end := time.Now()
		cost := end.Sub(start)
		log.Printf("api uri: %v, cost: %v", c.Request.RequestURI, cost.Seconds())
	}
}
