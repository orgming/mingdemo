package command

import "github.com/orgming/ming/framework/cobra"

func AddKernelCommands(root *cobra.Command) {
	root.AddCommand(DemoCmd)

	root.AddCommand(initAppCmd())
}
