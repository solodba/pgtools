package impl

import (
	"context"
	"fmt"
	"strings"

	"github.com/solodba/pgtools/apps/recover"
)

// 获取当前lsn号和当前lsn号所在的wal文件
func (i *impl) GetLsnAndFile(ctx context.Context) (*recover.FileInfo, error) {
	cmd := `/data/postgres/bin/psql -h localhost -d postgres -U postgres -A -t -c "select pg_current_wal_lsn(),pg_walfile_name(pg_current_wal_lsn()),pg_walfile_name_offset(pg_current_wal_lsn());"`
	result, err := i.cmdConf.RunShell(cmd)
	if err != nil {
		return nil, err
	}
	result = strings.Trim(result, "\n")
	lsnNum := strings.Split(result, "|")[0]
	walFileName := strings.Split(result, "|")[1]
	walFileNameAndOffset := strings.Split(result, "|")[2]
	fileInfo := recover.NewFileInfo()
	fileInfo.FilePath = "/data/postgres/data/pg_wal/" + walFileName
	fileInfo.LsnNum = lsnNum
	fmt.Println("=======================================================================================")
	fmt.Printf("当前的lsn号:%s\n", lsnNum)
	fmt.Printf("当前的wal文件和偏移量:%s\n", walFileNameAndOffset)
	return fileInfo, err
}

// 切换lsn日志
func (i *impl) SwitchWalLog(context.Context) error {
	cmd := `/data/postgres/bin/psql -h localhost -d postgres -U postgres -A -t -c "select pg_switch_wal();"`
	_, err := i.cmdConf.RunShell(cmd)
	if err != nil {
		return err
	}
	return nil
}

// 获取误删事务号
func (i *impl) GetDeleteTxid(ctx context.Context, fileInfo *recover.FileInfo) (*recover.TxInfo, error) {
	walDumpCmd := fmt.Sprintf(`/data/postgres/bin/pg_waldump -e %s %s`, fileInfo.LsnNum, fileInfo.FilePath)
	for {
		result, err := i.cmdConf.RunShell(walDumpCmd)
		if err != nil {
			err = i.SwitchWalLog(ctx)
			if err != nil {
				return nil, err
			}
			fmt.Println("=======================================================================================")
			fmt.Println("切换wal日志成功")
			continue
		}
		walLogList := strings.Split(result, "\n")
		deleteLogList := make([]string, 0)
		for _, item := range walLogList {
			if strings.Contains(item, "DELETE off") {
				deleteLogList = append(deleteLogList, item)
			}
		}
		fmt.Println("=======================================================================================")
		fmt.Println("误删除操作的日志")
		fmt.Println(deleteLogList[len(deleteLogList)-1])
		txStr := strings.Split(deleteLogList[len(deleteLogList)-1], ",")[1]
		txId := strings.Split(txStr, ":")[1]
		txInfo := recover.NewTxInfo()
		txInfo.TxId = txId
		return txInfo, nil
	}
}

// 关闭数据库
func (i *impl) StopDb(ctx context.Context) error {
	fmt.Println("=======================================================================================")
	fmt.Println("停止数据库")
	cmd := `su - postgres -c "pg_ctl -D /data/postgres/data stop -mf"`
	_, err := i.cmdConf.RunShell(cmd)
	if err != nil {
		return err
	}
	return nil
}

// 恢复数据库到指定事务号
func (i *impl) RecoverDbToTxid(ctx context.Context, txInfo *recover.TxInfo) error {
	fmt.Println("=======================================================================================")
	fmt.Println("误删除恢复")
	cmd := fmt.Sprintf(`su - postgres -c "/data/postgres/bin/pg_resetwal -x %s -D /data/postgres/data"`, txInfo.TxId)
	result, err := i.cmdConf.RunShell(cmd)
	if err != nil {
		return err
	}
	fmt.Print(result)
	return nil
}

// 启动数据库
func (i *impl) StartDb(ctx context.Context) error {
	fmt.Println("=======================================================================================")
	fmt.Println("启动数据库")
	cmd := `su - postgres -c "pg_ctl -D /data/postgres/data start"`
	_, err := i.cmdConf.RunShell(cmd)
	if err != nil {
		return err
	}
	return nil
}
