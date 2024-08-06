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

// ComsumeTopSql结构体
type ComsumeTopSql struct {
	UserId            string `json:"user_id"`
	DbId              string `json:"db_id"`
	Calls             string `json:"calls"`
	MinExecTime       string `json:"min_exec_time"`
	MaxExecTime       string `json:"max_exec_time"`
	MeanExecTime      string `json:"mean_exec_time"`
	TotalExecTime     string `json:"total_exec_time"`
	SharedBlksHit     string `json:"shared_blk_hint"`
	SharedBlksRead    string `json:"shared_blks_read"`
	SharedBlksDirtied string `json:"shared_blks_dirtied"`
	SharedBlksWritten string `json:"shared_blks_written"`
	TempBlksRead      string `json:"temp_blks_read"`
	TempBlksWritten   string `json:"temp_blks_written"`
	BlkReadTime       string `json:"blk_read_time"`
	BlkWriteTime      string `json:"blk_write_time"`
	Query             string `json:"query"`
}

// ComsumeTopSqlSet结构体
type ComsumeTopSqlSet struct {
	Total              int              `json:"total"`
	ComsumeTopSqlItems []*ComsumeTopSql `json:"comsume_top_sql_items"`
}

// ComsumeTopSql构造函数
func NewComsumeTopSql() *ComsumeTopSql {
	return &ComsumeTopSql{}
}

// ComsumeTopSqlSet构造函数
func NewComsumeTopSqlSet() *ComsumeTopSqlSet {
	return &ComsumeTopSqlSet{
		ComsumeTopSqlItems: make([]*ComsumeTopSql, 0),
	}
}

// ComsumeTopSqlSet结构体添加方法
func (c *ComsumeTopSqlSet) AddItems(items ...*ComsumeTopSql) {
	c.ComsumeTopSqlItems = append(c.ComsumeTopSqlItems, items...)
}

// WalFileInfo结构体
type WalFileInfo struct {
	// 是否开启归档
	ArchiveMode string `json:"archive_mode"`
	// 当前WAL文件总数
	WalFileCount string `json:"wal_file_count"`
	// 归档WAL文件总数
	ArchivedFileCount string `json:"archived_file_count"`
	// WAL文件归档速率
	ArchiveRate string `json:"archive_rate"`
	// 最后归档WAL文件名
	LastArchived string `json:"last_archived"`
	// 最后归档失败WAL文件名
	LastFailure string `json:"last_failure"`
	// 归档失败文件总数量
	ArchivedFailCount string `json:"archived_fail_count"`
	// 总体信息
	Total string `json:"total"`
	// walfile相关参数集合
	ParamSet *ParamSet `json:"param_set"`
}

// WalFileInfo结构体构造函数
func NewWalFileInfo() *WalFileInfo {
	return &WalFileInfo{
		ParamSet: NewParamSet(),
	}
}

// LockInfo结构体
type LockInfo struct {
	LockType string `json:"lock_type"`
	Granted  string `json:"granted"`
	Total    string `json:"total"`
}

// LockInfoSet结构体
type LockInfoSet struct {
	Total         int         `json:"total"`
	LockInfoItems []*LockInfo `json:"lock_info_items"`
}

// LockInfo结构体构造函数
func NewLockInfo() *LockInfo {
	return &LockInfo{}
}

// LockInfoSet结构体构造函数
func NewLockInfoSet() *LockInfoSet {
	return &LockInfoSet{
		LockInfoItems: make([]*LockInfo, 0),
	}
}

// LockInfoSet结构体添加方法
func (l *LockInfoSet) AddItems(items ...*LockInfo) {
	l.LockInfoItems = append(l.LockInfoItems, items...)
}

// VacuumInfo结构体
type VacuumInfo struct {
	Pid              string `json:"pid"`
	Datname          string `json:"datname"`
	TableName        string `json:"table_name"`
	Phase            string `json:"phase"`
	HeapBlksTotal    string `json:"heap_blks_total"`
	HeapBlksScanned  string `json:"heap_blks_scanned"`
	HeapBlksVacuumed string `json:"heap_blks_vacuumed"`
	IndexVacuumCount string `json:"index_vacuum_count"`
	MaxDeadTuples    string `json:"max_dead_tuples"`
	NumDeadTuples    string `json:"num_dead_tuples"`
}

// VacuumInfoSet结构体
type VacuumInfoSet struct {
	Total           int           `json:"total"`
	ParamSet        *ParamSet     `json:"param_set"`
	VacuumInfoItems []*VacuumInfo `json:"vacuum_info_items"`
}

// VacuumInfo结构体构造函数
func NewVacuumInfo() *VacuumInfo {
	return &VacuumInfo{}
}

// VacuumInfoSet结构体构造函数
func NewVacuumInfoSet() *VacuumInfoSet {
	return &VacuumInfoSet{
		VacuumInfoItems: make([]*VacuumInfo, 0),
		ParamSet:        NewParamSet(),
	}
}

// VacuumInfoSet结构体添加方法
func (v *VacuumInfoSet) AddItems(items ...*VacuumInfo) {
	v.VacuumInfoItems = append(v.VacuumInfoItems, items...)
}

// RoleInfo结构体
type RoleInfo struct {
	Name      string `json:"name"`
	Login     string `json:"login"`
	Repl      string `json:"repl"`
	Super     string `json:"super"`
	CreatRol  string `json:"creat_rol"`
	CreatDb   string `json:"creat_db"`
	BypassRls string `json:"bypass_rls"`
	Inherit   string `json:"inherit"`
	ConnLimit string `json:"conn_limit"`
	Expires   string `json:"expires"`
	MemberOf  string `json:"member_of"`
}

// RoleInfoSet结构体
type RoleInfoSet struct {
	Total         int         `json:"total"`
	RoleInfoItems []*RoleInfo `json:"role_info_items"`
}

// RoleInfo结构体构造函数
func NewRoleInfo() *RoleInfo {
	return &RoleInfo{}
}

// RoleInfoSet结构体构造函数
func NewRoleInfoSet() *RoleInfoSet {
	return &RoleInfoSet{
		RoleInfoItems: make([]*RoleInfo, 0),
	}
}

// RoleInfoSet结构体添加方法
func (r *RoleInfoSet) AddItems(items ...*RoleInfo) {
	r.RoleInfoItems = append(r.RoleInfoItems, items...)
}

// BackendInfo结构体
type BackendInfo struct {
	MaxConnect    string `json:"max_connect"`
	TotalBackends string `json:"total_backends"`
	WaitOnLocks   string `json:"wait_on_locks"`
	LongXact      string `json:"long_xact"`
	IdleInXact    string `json:"idle_in_xact"`
}

// BackendInfo结构体构造函数
func NewBackendInfo() *BackendInfo {
	return &BackendInfo{}
}

// TablespaceInfo结构体
type TablespaceInfo struct {
	Name      string `json:"name"`
	Owner     string `json:"owner"`
	Location  string `json:"location"`
	Size      string `json:"size"`
	DiskUsed  string `json:"disk_used"`
	InodeUsed string `json:"inode_used"`
}

// TablespaceInfoSet结构体
type TablespaceInfoSet struct {
	Total               int               `json:"total"`
	TablespaceInfoItems []*TablespaceInfo `json:"tablespace_info_items"`
}

// TablespaceInfo结构体构造函数
func NewTablespaceInfo() *TablespaceInfo {
	return &TablespaceInfo{}
}

// TablespaceInfoSet结构体构造函数
func NewTablespaceInfoSet() *TablespaceInfoSet {
	return &TablespaceInfoSet{
		TablespaceInfoItems: make([]*TablespaceInfo, 0),
	}
}

// TablespaceInfoSet结构体添加方法
func (t *TablespaceInfoSet) AddItems(items ...*TablespaceInfo) {
	t.TablespaceInfoItems = append(t.TablespaceInfoItems, items...)
}

// AwrData结构体
type AwrData struct {
	SystemInfo          *SystemInfo        `json:"system_info"`
	PgClusterInfo       *PgClusterInfo     `json:"pg_cluster_info"`
	ComsumeIoSqlSet     *ComsumeTopSqlSet  `json:"comsume_io_sql_set"`
	ComsumeTimeSqlSet   *ComsumeTopSqlSet  `json:"comsume_time_sql_set"`
	ComsumeBufferSqlSet *ComsumeTopSqlSet  `json:"comsume_buffer_sql_set"`
	ComsumeTempSqlSet   *ComsumeTopSqlSet  `json:"comsume_temp_sql_set"`
	WalFileInfo         *WalFileInfo       `json:"wal_file_info"`
	LockInfoSet         *LockInfoSet       `json:"lock_info_set"`
	VacuumInfoSet       *VacuumInfoSet     `json:"vacuum_info_set"`
	RoleInfoSet         *RoleInfoSet       `json:"role_info_set"`
	BackendInfo         *BackendInfo       `json:"backend_info"`
	TablespaceInfoSet   *TablespaceInfoSet `json:"tablespace_info_set"`
}

// AwrData结构体初始化函数
func NewAwrData() *AwrData {
	return &AwrData{
		SystemInfo:          NewSystemInfo(),
		PgClusterInfo:       NewPgClusterInfo(),
		ComsumeIoSqlSet:     NewComsumeTopSqlSet(),
		ComsumeTimeSqlSet:   NewComsumeTopSqlSet(),
		ComsumeBufferSqlSet: NewComsumeTopSqlSet(),
		ComsumeTempSqlSet:   NewComsumeTopSqlSet(),
		WalFileInfo:         NewWalFileInfo(),
		LockInfoSet:         NewLockInfoSet(),
		VacuumInfoSet:       NewVacuumInfoSet(),
		RoleInfoSet:         NewRoleInfoSet(),
		BackendInfo:         NewBackendInfo(),
		TablespaceInfoSet:   NewTablespaceInfoSet(),
	}
}

// 取模运算函数
func Mod(x, y int) int {
	return x % y
}
