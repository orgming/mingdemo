package demo

import (
	"github.com/orgming/ming/framework/cobra"
	"log"
)

func InitFoo() *cobra.Command {
	FooCmd.AddCommand(Foo2Cmd)
	return FooCmd
}

var FooCmd = &cobra.Command{
	Use:     "foo",
	Short:   "foo 命令简要说明",
	Long:    "foo 命令长说明",
	Aliases: []string{"fo", "f"},
	Example: "foo命令例子",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		log.Println(container)
		return nil
	},
}

var Foo2Cmd = &cobra.Command{
	Use:     "foo2",
	Short:   "foo2 命令简要说明",
	Long:    "foo2 命令长说明",
	Aliases: []string{"fo2", "f2"},
	Example: "foo2命令例子",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		log.Println(container)
		return nil
	},
}
