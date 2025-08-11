package command

import (
	"fmt"
	"github.com/erikdubbelboer/gspt"
	"github.com/orgming/ming/framework/cobra"
	"github.com/orgming/ming/framework/contract"
	"github.com/orgming/ming/framework/util"
	"github.com/sevlyar/go-daemon"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

var cronDaemon = false

func initCronCmd() *cobra.Command {
	cronStartCmd.Flags().BoolVarP(&cronDaemon, "daemon", "d", false, "start serve daemon")
	cronCmd.AddCommand(cronRestartCmd)
	cronCmd.AddCommand(cronStateCmd)
	cronCmd.AddCommand(cronStopCmd)
	cronCmd.AddCommand(cronListCmd)
	cronCmd.AddCommand(cronStartCmd)
	return cronCmd
}

var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "定时任务相关命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
		}
		return nil
	},
}

var cronListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有定时任务",
	RunE: func(cmd *cobra.Command, args []string) error {
		cronSpecs := cmd.Root().CronSpecs
		ps := [][]string{}
		for _, cronSpec := range cronSpecs {
			ps = append(ps, []string{cronSpec.Type, cronSpec.Spec, cronSpec.Cmd.Use, cronSpec.Cmd.Short, cronSpec.ServiceName})
		}
		util.PrettyPrint(ps)
		return nil
	},
}

var cronStartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动cron常驻进程",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)

		folder := appService.RuntimeFolder()
		serverPIDFile := filepath.Join(folder, "cron.pid")
		logFolder := appService.LogFolder()
		serverLogFile := filepath.Join(logFolder, "cron.log")
		curFolder := appService.BaseFolder()

		if cronDaemon {
			ctx := &daemon.Context{
				PidFileName: serverPIDFile,
				PidFilePerm: 0644,
				LogFileName: serverLogFile,
				LogFilePerm: 0640,
				WorkDir:     curFolder,
				Umask:       027,
				// 子进程的参数，按照这个参数设置，子进程的命令为 ./hade cron start --daemon=true
				Args: []string{"", "cron", "start", "--daemon=true"},
			}
			// 启动子进程，p不为空表示当前是父进程，p为空表示当前是子进程
			p, err := ctx.Reborn()
			if err != nil {
				return err
			}
			if p != nil {
				// 父进程就直接打印成功信息
				fmt.Println("cron serve started, pid: ", p.Pid)
				fmt.Println("log file: ", serverLogFile)
				return nil
			}
			defer ctx.Release()

			fmt.Println("daemon started")
			gspt.SetProcTitle("ming cron")
			cmd.Root().Cron.Run()
			return nil
		}

		fmt.Println("start cron job")
		content := strconv.Itoa(os.Getpid())
		fmt.Println("[PID]", content)
		err := os.WriteFile(serverPIDFile, []byte(content), 0664)
		if err != nil {
			return err
		}

		gspt.SetProcTitle("ming cron")
		cmd.Root().Cron.Run()
		return nil
	},
}

var cronRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "重启cron常驻进程",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)

		serverPidFile := filepath.Join(appService.RuntimeFolder(), "cron.pid")

		content, err := os.ReadFile(serverPidFile)
		if err != nil {
			return err
		}

		if content != nil && len(content) > 0 {
			pid, err := strconv.Atoi(string(content))
			if err != nil {
				return err
			}
			if util.CheckProcessExist(pid) {
				if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
					return err
				}
				// check process closed
				for i := 0; i < 10; i++ {
					if util.CheckProcessExist(pid) == false {
						break
					}
					time.Sleep(1 * time.Second)
				}
				fmt.Println("kill process:" + strconv.Itoa(pid))
			}
		}
		cronDaemon = true
		return cronStartCmd.RunE(cmd, args)
	},
}

var cronStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "停止cron常驻进程",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)

		serverPidFile := filepath.Join(appService.RuntimeFolder(), "cron.pid")

		content, err := os.ReadFile(serverPidFile)
		if err != nil {
			return err
		}

		if content != nil && len(content) > 0 {
			pid, err := strconv.Atoi(string(content))
			if err != nil {
				return err
			}
			if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
				return err
			}
			if err := os.WriteFile(serverPidFile, []byte{}, 0644); err != nil {
				return err
			}
			fmt.Println("stop pid:", pid)
		}
		return nil
	},
}

var cronStateCmd = &cobra.Command{
	Use:   "state",
	Short: "cron常驻进程状态",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)

		serverPidFile := filepath.Join(appService.RuntimeFolder(), "cron.pid")

		content, err := os.ReadFile(serverPidFile)
		if err != nil {
			return err
		}

		if content != nil && len(content) > 0 {
			pid, err := strconv.Atoi(string(content))
			if err != nil {
				return err
			}
			if util.CheckProcessExist(pid) {
				fmt.Println("cron server started, pid:", pid)
				return nil
			}
		}
		fmt.Println("no cron server start")
		return nil
	},
}
