package command

import (
	"fmt"
	"github.com/orgming/ming/framework/cobra"
	"github.com/orgming/ming/framework/contract"
	"github.com/orgming/ming/framework/util"
)

func initEnvCmd() *cobra.Command {
	envCmd.AddCommand(envListCmd)
	return envCmd
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "获取当前的App环境",
	Run: func(cmd *cobra.Command, args []string) {
		container := cmd.GetContainer()
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		fmt.Println("environment:", envService.AppEnv())
	},
}

var envListCmd = &cobra.Command{
	Use:   "list",
	Short: "获取所有的环境变量",
	Run: func(cmd *cobra.Command, args []string) {
		container := cmd.GetContainer()
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		envs := envService.All()
		outs := [][]string{}
		for k, v := range envs {
			outs = append(outs, []string{k, v})
		}
		util.PrettyPrint(outs)
	},
}
