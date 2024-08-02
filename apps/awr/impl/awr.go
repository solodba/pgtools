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

// 生成AWR数据
func (i *impl) GenAwrData(ctx context.Context) (*awr.AwrData, error) {
	systemInfo, err := i.GetSystemInfo(ctx)
	if err != nil {
		return nil, err
	}
	awrData := awr.NewAwrData()
	awrData.SystemInfo = systemInfo
	return awrData, nil
}

// 生成AWR报告
func (i *impl) GenAwrReport(ctx context.Context) error {
	awrData, err := i.GenAwrData(ctx)
	if err != nil {
		return err
	}
	tpl := template.New("awr")
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
