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
	awrData := awr.NewAwrData()
	awrData.SystemInfo = systemInfo
	awrData.PgClusterInfo = pgClusterInfo
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
