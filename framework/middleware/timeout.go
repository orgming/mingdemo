package middleware

import (
	"context"
	"fmt"
	"log"
	"mingdemo/framework"
	"time"
)

func Timeout(d time.Duration) framework.ControllerHandler {

	return func(c *framework.Context) error {
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
			c.SetStatus(500).Json("time out")
			log.Println(p)

		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			c.SetHasTimeout()
			c.SetStatus(500).Json("time out")
		}
		return nil
	}
}
