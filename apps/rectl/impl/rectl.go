package impl

import (
	"context"
	"sort"
	"strings"

	"github.com/solodba/pgtools/apps/rectl"
)

// 获取最新wal日志名
func (i *impl) GetNextWal(ctx context.Context) (string, error) {
	cmd := `cd /data/postgres/data/pg_wal;stat -c "%Y %n" * | sort -n | tail -1 | awk '{print $2}'`
	result, err := i.cmdConf.RunShell(cmd)
	if err != nil {
		return "", err
	}
	result = strings.Trim(result, "\n")
	timeline := result[0:8]
	logicalId := result[8:16]
	logSeqNum := result[16:]
	var nextWalName string
	if rectl.HexToDec(logSeqNum) == 255 {
		logicalStr := rectl.DecToHexWal(rectl.HexToDec(logicalId) + 1)
		nextWalName = timeline + logicalStr + "00000001"
		return nextWalName, nil
	}
	logSeqStr := rectl.DecToHexWal(rectl.HexToDec(logSeqNum) + 1)
	nextWalName = timeline + logicalId + logSeqStr
	return nextWalName, nil
}

// 获取下一个多事务ID和最旧多事务ID
func (i *impl) GetMxid(ctx context.Context) (*rectl.Mxid, error) {
	cmd := `ls /data/postgres/data/pg_multixact/offsets`
	result, err := i.cmdConf.RunShell(cmd)
	if err != nil {
		return nil, err
	}
	result = strings.Trim(result, "\n")
	resultList := strings.Split(result, " ")
	mxidList := make([]int, 0)
	for _, item := range resultList {
		mxidList = append(mxidList, int(rectl.HexToDec(item)))
	}
	sort.Ints(mxidList)
	minMxid := mxidList[0]
	maxMxid := mxidList[len(mxidList)-1]
	maxid := rectl.NewMxid()
	if minMxid == maxMxid && minMxid == 0 {
		maxid.OldestMxid = "0x" + rectl.DecToHexMxid(65536)
		maxid.NextMxid = "0x" + rectl.DecToHexMxid(65536)
	} else {
		maxid.OldestMxid = "0x" + rectl.DecToHexMxid(uint64(minMxid)*65536)
		maxid.NextMxid = "0x" + rectl.DecToHexMxid((uint64(maxMxid)+1)*65536)
	}
	return maxid, nil
}

// 获取下一个多事务处理偏移量
func (i *impl) GetNextMxidOffset(ctx context.Context) (string, error) {
	return "", nil
}

// 获取下一个事务ID
func (i *impl) GetNextXid(ctx context.Context) (string, error) {
	return "", nil
}
