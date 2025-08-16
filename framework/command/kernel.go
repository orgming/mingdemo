package command

import "github.com/orgming/ming/framework/cobra"

// AddKernelCommands 绑定框架命令
func AddKernelCommands(root *cobra.Command) {
	//root.AddCommand(DemoCmd)

	root.AddCommand(initAppCmd())
	root.AddCommand(initCronCmd())
	root.AddCommand(initEnvCmd())
	root.AddCommand(initBuildCmd())
	root.AddCommand(initDevCmd())
	root.AddCommand(initProviderCmd())
	root.AddCommand(initCmdCmd())
	root.AddCommand(initMiddlewareCmd())
}
