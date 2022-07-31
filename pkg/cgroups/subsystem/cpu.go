package subsystem

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type CpuSubSystem struct {
	apply bool
}

func (c *CpuSubSystem) Name() string {
	return "cpu"
}

func (c *CpuSubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	subSystemCgroupsPath, err := GetCgroupPath(c.Name(), cgroupPath, true)
	if err != nil {
		logrus.Errorf("get %s path, err: %v", cgroupPath, err)
		return err
	}

	if res.CpuShare != "" {
		c.apply = true
		if err = ioutil.WriteFile(path.Join(subSystemCgroupsPath, "cpu.shares"), []byte(res.CpuShare), 0644); err != nil {
			logrus.Errorf("failed to write file cpu.shares, err: %+v", err)
			return err
		}
	}

	return nil
}

func (c *CpuSubSystem) Remove(cgroupPath string) error {
	subSystemCgroupsPath, err := GetCgroupPath(c.Name(), cgroupPath, false)
	if err != nil {
		return err
	}
	return os.RemoveAll(subSystemCgroupsPath)
}

func (c *CpuSubSystem) Apply(cgroupPath string, pid int) error {
	if !c.apply {
		return nil
	}

	subSystemCgroupsPath, err := GetCgroupPath(c.Name(), cgroupPath, false)
	if err != nil {
		return err
	}

	taskPath := path.Join(subSystemCgroupsPath, "tasks")
	if err = ioutil.WriteFile(taskPath, []byte(strconv.Itoa(pid)), 0644); err != nil {
		return err
	}
	return nil
}
