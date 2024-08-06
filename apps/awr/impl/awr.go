package impl

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"

	"github.com/solodba/pgtools/apps/awr"
)

// 获取当前系统信息
func (i *impl) GetSystemInfo(ctx context.Context) (*awr.SystemInfo, error) {
	systemInfo := awr.NewSystemInfo()
	hostname, err := exec.Command("hostname").CombinedOutput()
	if err != nil {
		return nil, err
	}
	systemInfo.Hostname = strings.Trim(string(hostname), "\n")
	runTime, err := exec.Command("bash", "-c", `uptime | awk '{print $3}' | tr -d ','`).CombinedOutput()
	if err != nil {
		return nil, err
	}
	systemInfo.RunTime = strings.Trim(string(runTime), "\n")
	cpuType, err := exec.Command("bash", "-c", `lscpu | grep "^Model name:" | awk -F ':' '{print $2}' | awk '{$1=$1}1'`).CombinedOutput()
	if err != nil {
		return nil, err
	}
	cpuNum, err := exec.Command("bash", "-c", `lscpu | grep '^CPU(s)' | awk -F ':' '{print $2}' | awk '{$1=$1}1'`).CombinedOutput()
	if err != nil {
		return nil, err
	}
	systemInfo.Cpu = fmt.Sprintf("%s x %s", strings.Trim(string(cpuNum), "\n"), strings.Trim(string(cpuType), "\n"))
	loadAvg, err := exec.Command("bash", "-c", `uptime | awk -F 'load average:' '{print $2}' | awk '{$1=$1}1'`).CombinedOutput()
	if err != nil {
		return nil, err
	}
	systemInfo.LoadAvg = strings.Trim(string(loadAvg), "\n")
	memTotal, err := exec.Command("bash", "-c", `free -m | grep '^Mem:' | awk '{print $2}'`).CombinedOutput()
	if err != nil {
		return nil, err
	}
	memUsed, err := exec.Command("bash", "-c", `free -m | grep '^Mem:' | awk '{print $3}'`).CombinedOutput()
	if err != nil {
		return nil, err
	}
	memFree, err := exec.Command("bash", "-c", `free -m | grep '^Mem:' | awk '{print $4}'`).CombinedOutput()
	if err != nil {
		return nil, err
	}
	systemInfo.Memory = fmt.Sprintf("total=%s MiB, used=%s MiB, free=%s MiB",
		strings.Trim(string(memTotal), "\n"),
		strings.Trim(string(memUsed), "\n"),
		strings.Trim(string(memFree), "\n"))
	swapTotal, err := exec.Command("bash", "-c", `free -m | grep '^Swap:' | awk '{print $2}'`).CombinedOutput()
	if err != nil {
		return nil, err
	}
	swapUsed, err := exec.Command("bash", "-c", `free -m | grep '^Swap:' | awk '{print $3}'`).CombinedOutput()
	if err != nil {
		return nil, err
	}
	swapFree, err := exec.Command("bash", "-c", `free -m | grep '^Swap:' | awk '{print $4}'`).CombinedOutput()
	if err != nil {
		return nil, err
	}
	systemInfo.Swap = fmt.Sprintf("total=%s MiB, used=%s MiB, free=%s MiB",
		strings.Trim(string(swapTotal), "\n"),
		strings.Trim(string(swapUsed), "\n"),
		strings.Trim(string(swapFree), "\n"))
	return systemInfo, nil
}

// 获取当前数据库聚簇信息
func (i *impl) GetPgClusterInfo(ctx context.Context) (*awr.PgClusterInfo, error) {
	pgClusterInfo := awr.NewPgClusterInfo()
	serverVerion, err := exec.Command("bash", "-c", `su - postgres -c 'pg_controldata --version' | awk '{print $3}'`).CombinedOutput()
	if err != nil {
		return nil, err
	}
	pgClusterInfo.ServerVersion = strings.Trim(string(serverVerion), "\n")
	row := i.db.QueryRowContext(ctx, "select pg_postmaster_start_time()")
	err = row.Scan(&pgClusterInfo.ServerStartTime)
	if err != nil {
		return nil, err
	}
	systemIdentifier, err := exec.Command("bash", "-c", `su - postgres -c 'pg_controldata' | grep 'Database system identifier' | awk -F ':' '{print $2}' | awk '{$1=$1}1'`).CombinedOutput()
	if err != nil {
		return nil, err
	}
	pgClusterInfo.SystemIdentifier = strings.Trim(string(systemIdentifier), "\n")
	timeLine, err := exec.Command("bash", "-c", `su - postgres -c 'pg_controldata' | grep "Latest checkpoint's TimeLineID" | awk -F ':' '{print $2}' | awk '{$1=$1}1'`).CombinedOutput()
	if err != nil {
		return nil, err
	}
	pgClusterInfo.TimeLine = strings.Trim(string(timeLine), "\n")
	lastCheckpointTime, err := exec.Command("bash", "-c", `su - postgres -c 'pg_controldata' | grep 'Time of latest checkpoint' | awk -F 'Time of latest checkpoint:' '{print $2}' | awk '{$1=$1}1'`).CombinedOutput()
	if err != nil {
		return nil, err
	}
	pgClusterInfo.LastCheckpointTime = strings.Trim(string(lastCheckpointTime), "\n")
	redoLsn, err := exec.Command("bash", "-c", `su - postgres -c 'pg_controldata' | grep "Latest checkpoint's REDO location" | awk -F ':' '{print $2}' | awk '{$1=$1}1'`).CombinedOutput()
	if err != nil {
		return nil, err
	}
	pgClusterInfo.RedoLsn = strings.Trim(string(redoLsn), "\n")
	checkpointLsn, err := exec.Command("bash", "-c", `su - postgres -c 'pg_controldata' | grep "Latest checkpoint location" | awk -F ':' '{print $2}' | awk '{$1=$1}1'`).CombinedOutput()
	if err != nil {
		return nil, err
	}
	pgClusterInfo.CheckpointLsn = strings.Trim(string(checkpointLsn), "\n")
	oldestXid, err := exec.Command("bash", "-c", `su - postgres -c 'pg_controldata' | grep "Latest checkpoint's oldestXID:" | awk -F ':' '{print $2}' | awk '{$1=$1}1'`).CombinedOutput()
	if err != nil {
		return nil, err
	}
	nextXid, err := exec.Command("bash", "-c", `su - postgres -c 'pg_controldata' | grep "Latest checkpoint's NextXID" | awk -F ':' '{print $NF}' | awk '{$1=$1}1'`).CombinedOutput()
	if err != nil {
		return nil, err
	}
	pgClusterInfo.TransactionId = fmt.Sprintf("oldest = %s, next = %s", strings.Trim(string(oldestXid), "\n"), strings.Trim(string(nextXid), "\n"))
	var flag string
	row = i.db.QueryRowContext(ctx, "select pg_is_in_recovery()")
	err = row.Scan(&flag)
	if err != nil {
		return nil, err
	}
	switch flag {
	case "false":
		pgClusterInfo.RecoveryMode = "no"
	case "true":
		pgClusterInfo.RecoveryMode = "yes"
	default:
		pgClusterInfo.RecoveryMode = "N/S"
	}
	args := []string{"shared_buffers",
		"work_mem",
		"maintenance_work_mem",
		"temp_buffers",
		"autovacuum_work_mem",
		"temp_file_limit",
		"max_worker_processes",
		"autovacuum_max_workers",
		"max_parallel_workers_per_gather",
		"effective_io_concurrency"}
	paramSet := awr.NewParamSet()
	for _, arg := range args {
		param := awr.NewParam()
		row = i.db.QueryRowContext(ctx, "select setting from pg_settings where name=$1", arg)
		err = row.Scan(&param.Value)
		if err != nil {
			return nil, err
		}
		param.Name = arg
		paramSet.AddItems(param)
	}
	paramSet.Total = len(paramSet.ParamItems)
	pgClusterInfo.ParamSet = paramSet
	return pgClusterInfo, nil
}

// 获取当前IO消耗TOP 10的SQL
func (i *impl) GetComsumeIoSql(ctx context.Context) (*awr.ComsumeTopSqlSet, error) {
	sql := `
select 
userid::regrole,
dbid,
calls,
min_exec_time,
max_exec_time,
mean_exec_time,
total_exec_time,
shared_blks_hit,
shared_blks_read,
shared_blks_dirtied,
shared_blks_written,
temp_blks_read,
temp_blks_written,
blk_read_time,
blk_write_time,
query
from pg_stat_statements
order by (blk_read_time+blk_write_time)/calls desc limit 10
`
	rows, err := i.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comsumeTopSqlSet := awr.NewComsumeTopSqlSet()
	for rows.Next() {
		comsumeTopSql := awr.NewComsumeTopSql()
		err = rows.Scan(
			&comsumeTopSql.UserId,
			&comsumeTopSql.DbId,
			&comsumeTopSql.Calls,
			&comsumeTopSql.MinExecTime,
			&comsumeTopSql.MaxExecTime,
			&comsumeTopSql.MeanExecTime,
			&comsumeTopSql.TotalExecTime,
			&comsumeTopSql.SharedBlksHit,
			&comsumeTopSql.SharedBlksRead,
			&comsumeTopSql.SharedBlksDirtied,
			&comsumeTopSql.SharedBlksWritten,
			&comsumeTopSql.TempBlksRead,
			&comsumeTopSql.TempBlksWritten,
			&comsumeTopSql.BlkReadTime,
			&comsumeTopSql.BlkWriteTime,
			&comsumeTopSql.Query)
		if err != nil {
			return nil, err
		}
		if len(comsumeTopSql.Query) > 60 {
			comsumeTopSql.Query = comsumeTopSql.Query[0:61] + "..."
		}
		comsumeTopSqlSet.AddItems(comsumeTopSql)
	}
	comsumeTopSqlSet.Total = len(comsumeTopSqlSet.ComsumeTopSqlItems)
	return comsumeTopSqlSet, nil
}

// 获取当前耗时TOP 10的SQL
func (i *impl) GetComsumeTimeSql(ctx context.Context) (*awr.ComsumeTopSqlSet, error) {
	sql := `
select 
userid::regrole,
dbid,
calls,
min_exec_time,
max_exec_time,
mean_exec_time,
total_exec_time,
shared_blks_hit,
shared_blks_read,
shared_blks_dirtied,
shared_blks_written,
temp_blks_read,
temp_blks_written,
blk_read_time,
blk_write_time,
query
from pg_stat_statements
order by total_exec_time desc limit 10
`
	rows, err := i.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comsumeTopSqlSet := awr.NewComsumeTopSqlSet()
	for rows.Next() {
		comsumeTopSql := awr.NewComsumeTopSql()
		err = rows.Scan(
			&comsumeTopSql.UserId,
			&comsumeTopSql.DbId,
			&comsumeTopSql.Calls,
			&comsumeTopSql.MinExecTime,
			&comsumeTopSql.MaxExecTime,
			&comsumeTopSql.MeanExecTime,
			&comsumeTopSql.TotalExecTime,
			&comsumeTopSql.SharedBlksHit,
			&comsumeTopSql.SharedBlksRead,
			&comsumeTopSql.SharedBlksDirtied,
			&comsumeTopSql.SharedBlksWritten,
			&comsumeTopSql.TempBlksRead,
			&comsumeTopSql.TempBlksWritten,
			&comsumeTopSql.BlkReadTime,
			&comsumeTopSql.BlkWriteTime,
			&comsumeTopSql.Query)
		if err != nil {
			return nil, err
		}
		if len(comsumeTopSql.Query) > 60 {
			comsumeTopSql.Query = comsumeTopSql.Query[0:61] + "..."
		}
		comsumeTopSqlSet.AddItems(comsumeTopSql)
	}
	comsumeTopSqlSet.Total = len(comsumeTopSqlSet.ComsumeTopSqlItems)
	return comsumeTopSqlSet, nil
}

// 获取当前消耗Buffer TOP 10的SQL
func (i *impl) GetComsumeBufferSql(ctx context.Context) (*awr.ComsumeTopSqlSet, error) {
	sql := `
select 
userid::regrole,
dbid,
calls,
min_exec_time,
max_exec_time,
mean_exec_time,
total_exec_time,
shared_blks_hit,
shared_blks_read,
shared_blks_dirtied,
shared_blks_written,
temp_blks_read,
temp_blks_written,
blk_read_time,
blk_write_time,
query
from pg_stat_statements
order by (shared_blks_hit+shared_blks_dirtied) desc limit 10
`
	rows, err := i.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comsumeTopSqlSet := awr.NewComsumeTopSqlSet()
	for rows.Next() {
		comsumeTopSql := awr.NewComsumeTopSql()
		err = rows.Scan(
			&comsumeTopSql.UserId,
			&comsumeTopSql.DbId,
			&comsumeTopSql.Calls,
			&comsumeTopSql.MinExecTime,
			&comsumeTopSql.MaxExecTime,
			&comsumeTopSql.MeanExecTime,
			&comsumeTopSql.TotalExecTime,
			&comsumeTopSql.SharedBlksHit,
			&comsumeTopSql.SharedBlksRead,
			&comsumeTopSql.SharedBlksDirtied,
			&comsumeTopSql.SharedBlksWritten,
			&comsumeTopSql.TempBlksRead,
			&comsumeTopSql.TempBlksWritten,
			&comsumeTopSql.BlkReadTime,
			&comsumeTopSql.BlkWriteTime,
			&comsumeTopSql.Query)
		if err != nil {
			return nil, err
		}
		if len(comsumeTopSql.Query) > 60 {
			comsumeTopSql.Query = comsumeTopSql.Query[0:61] + "..."
		}
		comsumeTopSqlSet.AddItems(comsumeTopSql)
	}
	comsumeTopSqlSet.Total = len(comsumeTopSqlSet.ComsumeTopSqlItems)
	return comsumeTopSqlSet, nil
}

// 获取当前消耗temp空间的SQL
func (i *impl) GetComsumeTempSql(ctx context.Context) (*awr.ComsumeTopSqlSet, error) {
	sql := `
select 
userid::regrole,
dbid,
calls,
min_exec_time,
max_exec_time,
mean_exec_time,
total_exec_time,
shared_blks_hit,
shared_blks_read,
shared_blks_dirtied,
shared_blks_written,
temp_blks_read,
temp_blks_written,
blk_read_time,
blk_write_time,
query
from pg_stat_statements
order by temp_blks_written desc limit 10
`
	rows, err := i.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comsumeTopSqlSet := awr.NewComsumeTopSqlSet()
	for rows.Next() {
		comsumeTopSql := awr.NewComsumeTopSql()
		err = rows.Scan(
			&comsumeTopSql.UserId,
			&comsumeTopSql.DbId,
			&comsumeTopSql.Calls,
			&comsumeTopSql.MinExecTime,
			&comsumeTopSql.MaxExecTime,
			&comsumeTopSql.MeanExecTime,
			&comsumeTopSql.TotalExecTime,
			&comsumeTopSql.SharedBlksHit,
			&comsumeTopSql.SharedBlksRead,
			&comsumeTopSql.SharedBlksDirtied,
			&comsumeTopSql.SharedBlksWritten,
			&comsumeTopSql.TempBlksRead,
			&comsumeTopSql.TempBlksWritten,
			&comsumeTopSql.BlkReadTime,
			&comsumeTopSql.BlkWriteTime,
			&comsumeTopSql.Query)
		if err != nil {
			return nil, err
		}
		if len(comsumeTopSql.Query) > 60 {
			comsumeTopSql.Query = comsumeTopSql.Query[0:61] + "..."
		}
		comsumeTopSqlSet.AddItems(comsumeTopSql)
	}
	comsumeTopSqlSet.Total = len(comsumeTopSqlSet.ComsumeTopSqlItems)
	return comsumeTopSqlSet, nil
}

// 获取当前数据库聚簇WAL Files信息
func (i *impl) GetPgWalFileInfo(ctx context.Context) (*awr.WalFileInfo, error) {
	walFileInfo := awr.NewWalFileInfo()
	row := i.db.QueryRowContext(ctx, `show archive_mode`)
	err := row.Scan(&walFileInfo.ArchiveMode)
	if err != nil {
		return nil, err
	}
	row = i.db.QueryRowContext(ctx, `SELECT count(*) AS wal_file_count FROM pg_ls_waldir()`)
	err = row.Scan(&walFileInfo.WalFileCount)
	if err != nil {
		return nil, err
	}
	row = i.db.QueryRowContext(ctx, `select archived_count,COALESCE(last_archived_wal,''),COALESCE(last_failed_wal,''),failed_count from pg_stat_archiver`)
	err = row.Scan(&walFileInfo.ArchivedFileCount, &walFileInfo.LastArchived, &walFileInfo.LastFailure, &walFileInfo.ArchivedFailCount)
	if err != nil {
		return nil, err
	}
	_, err = i.db.ExecContext(ctx, `DROP TABLE IF EXISTS temp_archiver_stats`)
	if err != nil {
		return nil, err
	}
	_, err = i.db.ExecContext(ctx, `SELECT now() AS current_time,archived_count AS total_archived INTO temp_archiver_stats FROM pg_stat_archiver`)
	if err != nil {
		return nil, err
	}
	_, err = i.db.ExecContext(ctx, `select pg_sleep(10)`)
	if err != nil {
		return nil, err
	}
	sql := `SELECT (a.archived_count - t.total_archived) / EXTRACT(EPOCH FROM (now() - t.current_time))/60 AS archive_rate_per_min FROM pg_stat_archiver a,temp_archiver_stats t`
	row = i.db.QueryRowContext(ctx, sql)
	err = row.Scan(&walFileInfo.ArchiveRate)
	if err != nil {
		return nil, err
	}
	walFileInfo.ArchiveRate = fmt.Sprintf("%s per min", walFileInfo.ArchiveRate)
	walFileInfo.Total = fmt.Sprintf("%s succeeded, %s failed", walFileInfo.ArchivedFileCount, walFileInfo.ArchivedFailCount)
	args := []string{"wal_level",
		"archive_timeout",
		"wal_compression",
		"max_wal_size",
		"min_wal_size",
		"checkpoint_timeout",
		"full_page_writes",
		"wal_keep_size",
	}
	paramSet := awr.NewParamSet()
	for _, arg := range args {
		param := awr.NewParam()
		row = i.db.QueryRowContext(ctx, "select setting from pg_settings where name=$1", arg)
		err = row.Scan(&param.Value)
		if err != nil {
			return nil, err
		}
		param.Name = arg
		paramSet.AddItems(param)
	}
	paramSet.Total = len(paramSet.ParamItems)
	walFileInfo.ParamSet = paramSet
	return walFileInfo, nil
}

// 获取当前所有锁信息
func (i *impl) GetPgLockInfo(ctx context.Context) (*awr.LockInfoSet, error) {
	rows, err := i.db.QueryContext(ctx, `select locktype,granted,count(*) total from pg_locks group by locktype,granted`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	lockInfoSet := awr.NewLockInfoSet()
	for rows.Next() {
		lockInfo := awr.NewLockInfo()
		err = rows.Scan(&lockInfo.LockType, &lockInfo.Granted, &lockInfo.Total)
		if err != nil {
			return nil, err
		}
		lockInfoSet.AddItems(lockInfo)
	}
	lockInfoSet.Total = len(lockInfoSet.LockInfoItems)
	return lockInfoSet, nil
}

// 获取当前VACUUM信息
func (i *impl) GetPgVacuumInfo(ctx context.Context) (*awr.VacuumInfoSet, error) {
	sql := `
SELECT
    pid,
    datname,
    relid::regclass AS table_name,
    phase,
    heap_blks_total,
    heap_blks_scanned,
    heap_blks_vacuumed,
    index_vacuum_count,
    max_dead_tuples,
    num_dead_tuples
FROM
    pg_stat_progress_vacuum`
	rows, err := i.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	vacuumInfoSet := awr.NewVacuumInfoSet()
	for rows.Next() {
		vacuumInfo := awr.NewVacuumInfo()
		err = rows.Scan(
			&vacuumInfo.Pid,
			&vacuumInfo.Datname,
			&vacuumInfo.TableName,
			&vacuumInfo.Phase,
			&vacuumInfo.HeapBlksTotal,
			&vacuumInfo.HeapBlksScanned,
			&vacuumInfo.HeapBlksVacuumed,
			&vacuumInfo.IndexVacuumCount,
			&vacuumInfo.MaxDeadTuples,
			&vacuumInfo.NumDeadTuples,
		)
		if err != nil {
			return nil, err
		}
		vacuumInfoSet.AddItems(vacuumInfo)
	}
	vacuumInfoSet.Total = len(vacuumInfoSet.VacuumInfoItems)
	args := []string{"maintenance_work_mem",
		"autovacuum",
		"autovacuum_analyze_threshold",
		"autovacuum_vacuum_threshold",
		"autovacuum_freeze_max_age",
		"autovacuum_max_workers",
		"autovacuum_naptime",
		"vacuum_freeze_min_age",
		"vacuum_freeze_table_age",
	}
	paramSet := awr.NewParamSet()
	for _, arg := range args {
		param := awr.NewParam()
		row := i.db.QueryRowContext(ctx, "select setting from pg_settings where name=$1", arg)
		err = row.Scan(&param.Value)
		if err != nil {
			return nil, err
		}
		param.Name = arg
		paramSet.AddItems(param)
	}
	paramSet.Total = len(paramSet.ParamItems)
	vacuumInfoSet.ParamSet = paramSet
	return vacuumInfoSet, nil
}

// 获取当前角色信息
func (i *impl) GetPgRoleInfo(ctx context.Context) (*awr.RoleInfoSet, error) {
	sql := `
SELECT
R.rolname as "Name", 
R.rolcanlogin as "Login", 
R.rolreplication as "Repl", 
R.rolsuper as "Super", 
R.rolcreaterole as "Creat Rol",
R.rolcreatedb as "Creat DB", 
R.rolbypassrls as "Bypass RLS",
R.rolinherit as "Inherit", 
R.rolconnlimit as "Conn Limit",
COALESCE(EXTRACT(EPOCH FROM R.rolvaliduntil), 0) as "Expires",
ARRAY(SELECT pg_get_userbyid(M.roleid) FROM pg_auth_members AS M WHERE M.member = R.oid) as "Member Of"
FROM pg_roles AS R
ORDER BY R.oid ASC`
	rows, err := i.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	roleInfoSet := awr.NewRoleInfoSet()
	for rows.Next() {
		roleInfo := awr.NewRoleInfo()
		err = rows.Scan(
			&roleInfo.Name,
			&roleInfo.Login,
			&roleInfo.Repl,
			&roleInfo.Super,
			&roleInfo.CreatRol,
			&roleInfo.CreatDb,
			&roleInfo.BypassRls,
			&roleInfo.Inherit,
			&roleInfo.ConnLimit,
			&roleInfo.Expires,
			&roleInfo.MemberOf,
		)
		if err != nil {
			return nil, err
		}
		roleInfoSet.AddItems(roleInfo)
	}
	roleInfoSet.Total = len(roleInfoSet.RoleInfoItems)
	return roleInfoSet, nil
}

// 获取当前后端会话信息
func (i *impl) GetPgBackendInfo(ctx context.Context) (*awr.BackendInfo, error) {
	backendInfo := awr.NewBackendInfo()
	row := i.db.QueryRowContext(ctx, `show max_connections`)
	err := row.Scan(&backendInfo.MaxConnect)
	if err != nil {
		return nil, err
	}
	row = i.db.QueryRowContext(ctx, `select count(*) from pg_stat_activity`)
	err = row.Scan(&backendInfo.TotalBackends)
	if err != nil {
		return nil, err
	}
	row = i.db.QueryRowContext(ctx, `select count(*) from pg_stat_activity where wait_event_type='Lock'`)
	err = row.Scan(&backendInfo.WaitOnLocks)
	if err != nil {
		return nil, err
	}
	row = i.db.QueryRowContext(ctx, `select count(*) from pg_stat_activity WHERE xact_start IS NOT NULL AND (now() - query_start) > interval '10 minutes'`)
	err = row.Scan(&backendInfo.LongXact)
	if err != nil {
		return nil, err
	}
	row = i.db.QueryRowContext(ctx, `select count(*) from pg_stat_activity where state='idle'`)
	err = row.Scan(&backendInfo.IdleInXact)
	if err != nil {
		return nil, err
	}
	return backendInfo, nil
}

// 获取当前表空间信息
func (i *impl) GetPgTablespaceInfo(ctx context.Context) (*awr.TablespaceInfoSet, error) {
	rootDiskSize, err := exec.Command("bash", "-c", `df -h / | tail -n +2 |  awk '{print $2}'`).CombinedOutput()
	if err != nil {
		return nil, err
	}
	totalInodes, err := exec.Command("bash", "-c", `df -i / | tail -n +2 | awk '{print $2}'`).CombinedOutput()
	if err != nil {
		return nil, err
	}
	usedInodes, err := exec.Command("bash", "-c", `df -i / | tail -n +2 | awk '{print $3}'`).CombinedOutput()
	if err != nil {
		return nil, err
	}
	sql := `
	SELECT
	spcname "name",
	pg_get_userbyid(spcowner) "owner",
	COALESCE(pg_tablespace_location(oid),'') "location",
	pg_size_pretty(pg_tablespace_size(oid)) "size"
	FROM pg_tablespace
	ORDER BY oid ASC`
	rows, err := i.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tablespaceInfoSet := awr.NewTablespaceInfoSet()
	for rows.Next() {
		tablespaceInfo := awr.NewTablespaceInfo()
		err = rows.Scan(&tablespaceInfo.Name, &tablespaceInfo.Owner, &tablespaceInfo.Location, &tablespaceInfo.Size)
		if err != nil {
			return nil, err
		}
		if tablespaceInfo.Name == "pg_default" || tablespaceInfo.Name == "pg_global" {
			pgData, err := exec.Command("bash", "-c", `su - postgres -c "echo \$PGDATA"`).CombinedOutput()
			if err != nil {
				return nil, err
			}
			tablespaceInfo.Location = strings.Trim(string(pgData), "\n")
		}
		cmd := fmt.Sprintf(`du -sh %s | awk '{print $1}'`, tablespaceInfo.Location)
		tbsUsed, err := exec.Command("bash", "-c", cmd).CombinedOutput()
		if err != nil {
			return nil, err
		}
		tablespaceInfo.DiskUsed = fmt.Sprintf("%s of %s", strings.Trim(string(tbsUsed), "\n"), strings.Trim(string(rootDiskSize), "\n"))
		tablespaceInfo.InodeUsed = fmt.Sprintf("%s of %s", strings.Trim(string(usedInodes), "\n"), strings.Trim(string(totalInodes), "\n"))
		tablespaceInfoSet.AddItems(tablespaceInfo)
	}
	tablespaceInfoSet.Total = len(tablespaceInfoSet.TablespaceInfoItems)
	return tablespaceInfoSet, nil
}

// 生成AWR数据
func (i *impl) GenAwrData(ctx context.Context) (*awr.AwrData, error) {
	systemInfo, err := i.GetSystemInfo(ctx)
	if err != nil {
		return nil, err
	}
	pgClusterInfo, err := i.GetPgClusterInfo(ctx)
	if err != nil {
		return nil, err
	}
	comsumeIoSqlSet, err := i.GetComsumeIoSql(ctx)
	if err != nil {
		return nil, err
	}
	comsumeTimeSqlSet, err := i.GetComsumeTimeSql(ctx)
	if err != nil {
		return nil, err
	}
	comsumeBufferSqlSet, err := i.GetComsumeBufferSql(ctx)
	if err != nil {
		return nil, err
	}
	comsumeTempSqlSet, err := i.GetComsumeTempSql(ctx)
	if err != nil {
		return nil, err
	}
	walFileInfo, err := i.GetPgWalFileInfo(ctx)
	if err != nil {
		return nil, err
	}
	lockInfoSet, err := i.GetPgLockInfo(ctx)
	if err != nil {
		return nil, err
	}
	vaccumInfoSet, err := i.GetPgVacuumInfo(ctx)
	if err != nil {
		return nil, err
	}
	roleInfoSet, err := i.GetPgRoleInfo(ctx)
	if err != nil {
		return nil, err
	}
	backendInfo, err := i.GetPgBackendInfo(ctx)
	if err != nil {
		return nil, err
	}
	tablespaceInfoSet, err := i.GetPgTablespaceInfo(ctx)
	if err != nil {
		return nil, err
	}
	awrData := awr.NewAwrData()
	awrData.SystemInfo = systemInfo
	awrData.PgClusterInfo = pgClusterInfo
	awrData.ComsumeIoSqlSet = comsumeIoSqlSet
	awrData.ComsumeTimeSqlSet = comsumeTimeSqlSet
	awrData.ComsumeBufferSqlSet = comsumeBufferSqlSet
	awrData.ComsumeTempSqlSet = comsumeTempSqlSet
	awrData.WalFileInfo = walFileInfo
	awrData.LockInfoSet = lockInfoSet
	awrData.VacuumInfoSet = vaccumInfoSet
	awrData.RoleInfoSet = roleInfoSet
	awrData.BackendInfo = backendInfo
	awrData.TablespaceInfoSet = tablespaceInfoSet
	return awrData, nil
}

// 生成AWR报告
func (i *impl) GenAwrReport(ctx context.Context) error {
	awrData, err := i.GenAwrData(ctx)
	if err != nil {
		return err
	}
	fmap := template.FuncMap{"mod": awr.Mod}
	tpl := template.New("awr").Funcs(fmap)
	parse, err := tpl.Parse(awr.AwrTpl)
	if err != nil {
		return err
	}
	awrFileName := fmt.Sprintf("pg_awr_report_%s.html", time.Now().Format("20060102150405"))
	fs, err := os.OpenFile(awrFileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	err = parse.Execute(fs, awrData)
	if err != nil {
		return err
	}
	fmt.Printf("当前目录生成PostgreSQL AWR报告[%s]成功\n", awrFileName)
	return nil
}
