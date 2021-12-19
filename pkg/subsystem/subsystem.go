package subsystem

// ResourceConfig is the config of container
type ResourceConfig struct {
	MemoryLimit string
	CpuShare    string
	CpuSet      string
}

type SubSystem interface {
	// Name return subsystem name, like cpu,memory
	Name() string
	// Set sets specified resource limit
	Set(cgroupPath string, resourceConfig *ResourceConfig)
	// Remove removes cgroup
	Remove(cgroupPath string) error
	// Apply a task to cgroup
	Apply(cgroupPath string, pid int) error
}

var Subsystems = []SubSystem{

}