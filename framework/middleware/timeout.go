package middleware

import (
	"context"
	"fmt"
	"github.com/orgming/ming/framework/gin"
	"log"
	"time"
)

func Timeout(d time.Duration) gin.HandlerFunc {

	return func(c *gin.Context) {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		durationCtx, cancelFunc := context.WithTimeout(c.BaseContext(), d)
		defer cancelFunc()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			// 执行具体的业务逻辑
			c.Next()

			finish <- struct{}{}
		}()

		// 业务逻辑后的操作
		select {
		case p := <-panicChan:
			c.ISetStatus(500).IJson("time out")
			log.Println(p)

		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			c.ISetStatus(500).IJson("time out")
		}
	}
}
