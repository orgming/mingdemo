package command

import (
	"fmt"
	"github.com/orgming/ming/framework/cobra"
	"github.com/orgming/ming/framework/contract"
)

var DemoCmd = &cobra.Command{
	Use:   "demo",
	Short: "demo for framework",
	Run: func(cmd *cobra.Command, args []string) {
		container := cmd.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)
		fmt.Println("app base folder:", appService.BaseFolder())
	},
}
