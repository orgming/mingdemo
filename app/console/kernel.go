package console

import (
	"github.com/orgming/ming/app/console/command/demo"
	"github.com/orgming/ming/framework"
	"github.com/orgming/ming/framework/cobra"
	"github.com/orgming/ming/framework/command"
)

// RunCommand 初始化根Command并运行
func RunCommand(container framework.Container) error {
	var rootCmd = &cobra.Command{
		Use:   "ming",
		Short: "ming 命令",
		Long:  "ming框架提供的命令行工具，使用这个命令行工具能很方便执行框架自带命令，也能很方便编写业务命令",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		// 不需要出现 cobra 默认的 completion 子命令
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
	rootCmd.SetContainer(container)
	// 绑定框架的命令
	command.AddKernelCommands(rootCmd)
	// 绑定业务的命令
	AddAppCommand(rootCmd)
	// 执行RootCommand
	return rootCmd.Execute()
}

func AddAppCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(demo.InitFoo())
}
