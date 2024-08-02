package awr

// 系统信息结构体
type SystemInfo struct {
	Hostname string `json:"hostname"`
	RunTime  string `json:"run_time"`
	Cpu      string `json:"cpu"`
	LoadAvg  string `json:"load_avg"`
	Memory   string `json:"memory"`
	Swap     string `json:"swap"`
}

// 系统信息结构体构造函数
func NewSystemInfo() *SystemInfo {
	return &SystemInfo{}
}
