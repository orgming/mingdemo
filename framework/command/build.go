package command

import (
	"fmt"
	"github.com/orgming/ming/framework/cobra"
	"log"
	"os/exec"
)

func initBuildCmd() *cobra.Command {
	buildCmd.AddCommand(buildFrontendCmd)
	buildCmd.AddCommand(buildSelfCmd)
	buildCmd.AddCommand(buildBackendCmd)
	buildCmd.AddCommand(buildAllCmd)
	return buildCmd
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "编译相关命令",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			c.Help()
		}
		return nil
	},
}

var buildFrontendCmd = &cobra.Command{
	Use:   "frontend",
	Short: "使用npm编译前端",
	RunE: func(c *cobra.Command, args []string) error {
		// 获取path路径下的npm命令
		path, err := exec.LookPath("npm")
		if err != nil {
			log.Fatalln("npm命令不存在")
		}

		// 执行npm run build命令
		cmd := exec.Command(path, "run", "build")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("============= 前端编译失败 ============")
			fmt.Println(string(output))
			fmt.Println("============= 前端编译失败 ============")
			return err
		}
		fmt.Print(string(output))
		fmt.Println("============= 前端编译成功 ============")
		return nil
	},
}

var buildSelfCmd = &cobra.Command{
	Use:   "self",
	Short: "编译ming命令",
	RunE: func(c *cobra.Command, args []string) error {
		path, err := exec.LookPath("go")
		if err != nil {
			log.Fatalln("ming go: please install go in path first")
		}

		cmd := exec.Command(path, "build", "-o", "ming", "./")
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("go build error:")
			fmt.Println(string(out))
			fmt.Println("--------------")
			return err
		}
		fmt.Println("build success please run ./ming direct")
		return nil
	},
}

var buildBackendCmd = &cobra.Command{
	Use:   "backend",
	Short: "使用go编译后端",
	RunE: func(c *cobra.Command, args []string) error {
		return buildSelfCmd.RunE(c, args)
	},
}

var buildAllCmd = &cobra.Command{
	Use:   "all",
	Short: "同时编译前端和后端",
	RunE: func(c *cobra.Command, args []string) error {
		err := buildFrontendCmd.RunE(c, args)
		if err != nil {
			return err
		}
		err = buildSelfCmd.RunE(c, args)
		if err != nil {
			return err
		}
		return nil
	},
}
