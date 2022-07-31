package main

import (
	"fmt"

	"docker-go/pkg/cgroups/subsystem"

	"github.com/urfave/cli"
)

// 启动一个namespace隔离的容器进程
var runCommand = cli.Command{
	Name:  "run",
	Usage: "Create a container with namespace and cgroups limit",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "it",
			Usage: "enable tty",
		},
		cli.StringFlag{
			Name:  "m",
			Usage: "memory limit",
		},
		cli.StringFlag{
			Name:  "cpushare",
			Usage: "cpushare limit",
		},
		cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpuset limit",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("missing container args")
		}
		tty := context.Bool("it")

		resourceConfig := &subsystem.ResourceConfig{
			MemoryLimit: context.String("m"),
			CpuShare:    context.String("cpuset"),
			CpuSet:      context.String("cpushare"),
		}

		// 运行容器后的第一个命令，本例中即top命令
		var cmdArray []string
		for _, arg := range context.Args() {
			cmdArray = append(cmdArray, arg)
		}
		Run(cmdArray, tty, resourceConfig)
		return nil
	},
}

// initCommand is executed in container, do mount operation and run user's process
var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init run user's process in container. Do not call it outside",
}
