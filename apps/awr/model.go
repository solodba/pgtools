package awr

// 系统信息结构体
type SystemInfo struct {
	// 主机名
	Hostname string `json:"hostname"`
	// 运行时间
	RunTime string `json:"run_time"`
	// CPU信息
	Cpu string `json:"cpu"`
	// 平均负载
	LoadAvg string `json:"load_avg"`
	// 内存信息
	Memory string `json:"memory"`
	// 交换分区信息
	Swap string `json:"swap"`
}

// 系统信息结构体构造函数
func NewSystemInfo() *SystemInfo {
	return &SystemInfo{}
}

// PgClusterInfo结构体
type PgClusterInfo struct {
	// pg版本号
	ServerVersion string `json:"server_version"`
	// pg启动时间
	ServerStartTime string `json:"server_start_time"`
	// pg标识符
	SystemIdentifier string `json:"system_identifier"`
	// pg时间线
	TimeLine string `json:"time_line"`
	// pg最后一次checkpoint时间
	LastCheckpointTime string `json:"last_checkpoint_time"`
	// pg的WAL日志lsn
	RedoLsn string `json:"redo_lsn"`
	// pg的checkpoint的lsn
	CheckpointLsn string `json:"checkpoint_lsn"`
	// pg的事务信息
	TransactionId string `json:"transaction_id"`
	// pg模式
	RecoveryMode string `json:"recovery_mode"`
	// pg cluster参数集合
	ParamSet *ParamSet `json:"param_set"`
}

// 查询参数结构体
type Param struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// 查询参数集合结构体
type ParamSet struct {
	Total      int      `json:"total"`
	ParamItems []*Param `json:"param_item"`
}

// 查询参数结构体构造函数
func NewParam() *Param {
	return &Param{}
}

// 查询参数集合结构体构造函数
func NewParamSet() *ParamSet {
	return &ParamSet{
		ParamItems: make([]*Param, 0),
	}
}

// 查询参数集合结构体添加方法
func (p *ParamSet) AddItems(item ...*Param) {
	p.ParamItems = append(p.ParamItems, item...)
}

// PgClusterInfo构造函数
func NewPgClusterInfo() *PgClusterInfo {
	return &PgClusterInfo{
		ParamSet: NewParamSet(),
	}
}

// AwrData结构体
type AwrData struct {
	SystemInfo    *SystemInfo    `json:"system_info"`
	PgClusterInfo *PgClusterInfo `json:"pg_cluster_info"`
}

// AwrData结构体初始化函数
func NewAwrData() *AwrData {
	return &AwrData{
		SystemInfo:    NewSystemInfo(),
		PgClusterInfo: NewPgClusterInfo(),
	}
}

// 取模运算函数
func Mod(x, y int) int {
	return x % y
}
