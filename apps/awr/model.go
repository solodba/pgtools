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
	QueryId           string `json:"query_id"`
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
	LongQuery         string `json:"long_query"`
}

// ComsumeTopSqlSet结构体
type ComsumeTopSqlSet struct {
	Total              int              `json:"total"`
	DbName             string           `json:"dbname"`
	Type               string           `json:"type"`
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

// ComsumeTopSqlTotalSet结构体
type ComsumeTopSqlTotalSet struct {
	ComsumeTopSqlSetItems []*ComsumeTopSqlSet `json:"comsume_top_sql_set_items"`
}

// ComsumeTopSqlTotalSet结构体构造函数
func NewComsumeTopSqlTotalSet() *ComsumeTopSqlTotalSet {
	return &ComsumeTopSqlTotalSet{
		ComsumeTopSqlSetItems: make([]*ComsumeTopSqlSet, 0),
	}
}

// ComsumeTopSqlTotalSet结构体添加方法
func (c *ComsumeTopSqlTotalSet) AddItems(items ...*ComsumeTopSqlSet) {
	c.ComsumeTopSqlSetItems = append(c.ComsumeTopSqlSetItems, items...)
}

// ComsumeAllSqlSet结构体
type ComsumeAllSqlSet struct {
	Total              int              `json:"total"`
	ComsumeAllSqlItems []*ComsumeAllSql `json:"comsume_all_sql"`
}

// ComsumeAllSql结构体
type ComsumeAllSql struct {
	QueryId   string `json:"query_id"`
	QueryText string `json:"query_text"`
}

// ComsumeAllSql结构体构造函数
func NewComsumeAllSql() *ComsumeAllSql {
	return &ComsumeAllSql{}
}

// ComsumeAllSqlSet结构体构造函数
func NewComsumeAllSqlSet() *ComsumeAllSqlSet {
	return &ComsumeAllSqlSet{
		ComsumeAllSqlItems: make([]*ComsumeAllSql, 0),
	}
}

// ComsumeAllSqlSet结构体添加方法
func (c *ComsumeAllSqlSet) AddItems(items ...*ComsumeAllSql) {
	c.ComsumeAllSqlItems = append(c.ComsumeAllSqlItems, items...)
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

// BgWriterInfo结构体
type BgWriterInfo struct {
	// 强制执行checkpoint的百分比
	ForceCp string `json:"force_cp"`
	// 两次检查点中间的平均时间间隔(分钟)
	AvgMinCp string `json:"avg_min_cp"`
	// 平均checkpoint写入时间(秒)
	AvgCpWriteTime string `json:"avg_cp_write_time"`
	// 平均checkpoint检查点执行sync同步时间（秒）
	AvgCpSyncTime string `json:"avg_cp_sync_time"`
	// 总写入的数据量（MB）
	TotalWrite string `json:"total_write"`
	// 每个checkpoint检查点写入脏块的平均数据量（MB）
	MbPerCp string `json:"mb_per_cp"`
	// checkpoint检查点每秒写入脏块的速率（MBps）
	CpMbps string `json:"cp_mbps"`
	// bgwriter后台写每秒的写入块的速率（MBps）
	BgWriterMbps string `json:"bg_writer_mbps"`
	// 后端进程每秒写入速率（MBps）
	BackendMbps string `json:"backend_mbps"`
	//  每秒总写入速率（MBps）
	TotalMbps string `json:"total_mbps"`
	//  新分配的缓冲区比例
	NewBufferRatio string `json:"new_buffer_ratio"`
	// 由检查点清理的缓冲区百分比
	CleanByCp string `json:"clean_by_cp"`
	// 由后台写入器清理的缓冲区百分比
	CleanByBgWriter string `json:"clean_by_bg_writer"`
	// 由后端进程清理的缓冲区百分比
	CleanByBackend string `json:"clean_by_backend"`
	//  bgwriter后台暂停的百分比
	BgWriterHaltsPerRuns string `json:"bg_writer_halts_per_runs"`
	// bgwriter后台由于LRU命中导致暂停的百分比
	BgWriterHaltDueToLruHit string `json:"bg_writer_halt_due_to_lru_hit"`
	// BgWriter相关参数集合
	ParamSet *ParamSet `json:"param_set"`
}

// BgWriterInfo结构体构造函数
func NewBgWriterInfo() *BgWriterInfo {
	return &BgWriterInfo{
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

// DbInfo结构体
type DbInfo struct {
	Name                  string                 `json:"name"`
	Owner                 string                 `json:"owner"`
	Tablespace            string                 `json:"tablespace"`
	Connections           string                 `json:"connections"`
	FrozenXidAge          string                 `json:"frozen_xid_age"`
	Transactions          string                 `json:"transactions"`
	CacheHits             string                 `json:"cache_hits"`
	RowsChanged           string                 `json:"rows_changed"`
	TotalTemp             string                 `json:"total_temp"`
	Problems              string                 `json:"problems"`
	Size                  string                 `json:"size"`
	InstalledExtensionSet *InstalledExtensionSet `json:"installed_extension_set"`
}

// InstalledExtension结构体
type InstalledExtension struct {
	Name             string `json:"name"`
	DefaultVersion   string `json:"default_version"`
	InstalledVersion string `json:"installed_version"`
	Comment          string `json:"comment"`
}

// InstalledExtensionSet结构体
type InstalledExtensionSet struct {
	Total                   int                   `json:"total"`
	InstalledExtensionItems []*InstalledExtension `json:"installed_extension_items"`
}

// InstalledExtension结构体构造函数
func NewInstalledExtension() *InstalledExtension {
	return &InstalledExtension{}
}

// InstalledExtensionSet结构体构造函数
func NewInstalledExtensionSet() *InstalledExtensionSet {
	return &InstalledExtensionSet{
		InstalledExtensionItems: make([]*InstalledExtension, 0),
	}
}

// InstalledExtensionSet结构体添加方法
func (i *InstalledExtensionSet) AddItems(items ...*InstalledExtension) {
	i.InstalledExtensionItems = append(i.InstalledExtensionItems, items...)
}

// DbInfoSet结构体
type DbInfoSet struct {
	Total       int       `json:"total"`
	DbInfoItems []*DbInfo `json:"db_info_items"`
}

// DbInfo结构体构造函数
func NewDbInfo() *DbInfo {
	return &DbInfo{
		InstalledExtensionSet: NewInstalledExtensionSet(),
	}
}

// DbInfoSet结构体构造函数
func NewDbInfoSet() *DbInfoSet {
	return &DbInfoSet{
		DbInfoItems: make([]*DbInfo, 0),
	}
}

// DbInfoSet结构体添加方法
func (d *DbInfoSet) AddItems(items ...*DbInfo) {
	d.DbInfoItems = append(d.DbInfoItems, items...)
}

// AwrData结构体
type AwrData struct {
	SystemInfo            *SystemInfo            `json:"system_info"`
	PgClusterInfo         *PgClusterInfo         `json:"pg_cluster_info"`
	ComsumeTopSqlTotalSet *ComsumeTopSqlTotalSet `json:"comsume_top_sql_total_set"`
	ComsumeAllSqlSet      *ComsumeAllSqlSet      `json:"comsume_all_sql_set"`
	WalFileInfo           *WalFileInfo           `json:"wal_file_info"`
	LockInfoSet           *LockInfoSet           `json:"lock_info_set"`
	VacuumInfoSet         *VacuumInfoSet         `json:"vacuum_info_set"`
	RoleInfoSet           *RoleInfoSet           `json:"role_info_set"`
	BackendInfo           *BackendInfo           `json:"backend_info"`
	TablespaceInfoSet     *TablespaceInfoSet     `json:"tablespace_info_set"`
	DbInfoSet             *DbInfoSet             `json:"db_info_set"`
	BgWriterInfo          *BgWriterInfo          `json:"bg_writer_info"`
}

// AwrData结构体初始化函数
func NewAwrData() *AwrData {
	return &AwrData{
		SystemInfo:            NewSystemInfo(),
		PgClusterInfo:         NewPgClusterInfo(),
		ComsumeTopSqlTotalSet: NewComsumeTopSqlTotalSet(),
		ComsumeAllSqlSet:      NewComsumeAllSqlSet(),
		WalFileInfo:           NewWalFileInfo(),
		LockInfoSet:           NewLockInfoSet(),
		VacuumInfoSet:         NewVacuumInfoSet(),
		RoleInfoSet:           NewRoleInfoSet(),
		BackendInfo:           NewBackendInfo(),
		TablespaceInfoSet:     NewTablespaceInfoSet(),
		DbInfoSet:             NewDbInfoSet(),
		BgWriterInfo:          NewBgWriterInfo(),
	}
}

// 取模运算函数
func Mod(x, y int) int {
	return x % y
}
