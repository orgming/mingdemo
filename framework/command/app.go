package command

import (
	"context"
	"github.com/orgming/ming/framework/cobra"
	"github.com/orgming/ming/framework/contract"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// app启动地址 类似localhost:8888 或者 :8888
var appAddress = ""

func initAppCmd() *cobra.Command {
	appStartCmd.Flags().StringVar(&appAddress, "address", ":8888", "设置app启动的地址，默认为：8888")

	appCmd.AddCommand(appStartCmd)
	return appCmd
}

// 命令行参数第一级为app的命令，它没有实际功能，只是打印帮助文档
var appCmd = &cobra.Command{
	Use:   "app",
	Short: "业务应用控制命令",
	Long:  "业务应用控制命令，其包含业务启动，关闭，重启，查询等功能",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Help()
		return nil
	},
}

// 启动一个web服务
var appStartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动一个Web服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		// 从服务容器中获取kernel的服务实例
		kernelService := container.MustMake(contract.KernelKey).(contract.Kernel)
		// 从kernel服务实例中获取引擎
		engine := kernelService.HttpEngine()

		server := &http.Server{
			Addr:    appAddress,
			Handler: engine,
		}

		go func() {
			server.ListenAndServe()
		}()

		// 当前的goroutine等待信号量
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		// 阻塞当前goroutine等待信号
		<-quit

		timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()

		if err := server.Shutdown(timeoutCtx); err != nil {
			log.Fatal("Server Shutdown: ", err)
		}

		return nil
	},
}
