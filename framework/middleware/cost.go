package middleware

import (
	"log"
	"mingdemo/framework"
	"time"
)

// Cost 请求时长统计
func Cost() framework.ControllerHandler {
	return func(c *framework.Context) error {
		start := time.Now()
		c.Next()
		end := time.Now()
		cost := end.Sub(start)

		log.Printf("api uri: %v, cost: %v", c.GetRequest().RequestURI, cost.Seconds())
		return nil
	}
}
