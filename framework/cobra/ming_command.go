package cobra

import (
	"github.com/orgming/ming/framework"
	"github.com/robfig/cron/v3"
	"log"
)

func (c *Command) GetContainer() framework.Container {
	return c.container
}

func (c *Command) SetContainer(container framework.Container) {
	c.container = container
}

// CronSpec Cron命令的信息，用于展示
type CronSpec struct {
	Type        string
	Cmd         *Command
	Spec        string
	ServiceName string
}

func (c *Command) SetParentNull() {
	c.parent = nil
}

// AddCronCommand 创建一个Cron任务
func (c *Command) AddCronCommand(spec string, cmd *Command) {
	// cron是挂载在根Command上
	root := c.Root()
	if root.Cron == nil {
		root.Cron = cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
		root.CronSpecs = []CronSpec{}
	}
	// 增加说明信息
	root.CronSpecs = append(root.CronSpecs, CronSpec{
		Type: "normal-cron",
		Cmd:  cmd,
		Spec: spec,
	})

	var cronCmd Command
	ctx := root.Context()
	cronCmd = *cmd
	cronCmd.args = []string{}
	cronCmd.SetParentNull()
	cronCmd.SetContainer(root.GetContainer())

	root.Cron.AddFunc(spec, func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()

		err := cronCmd.ExecuteContext(ctx)
		if err != nil {
			log.Println(err)
		}
	})

}
