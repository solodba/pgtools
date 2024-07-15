# pgtools
postgresql工具

# 使用方法
```sh
# 检查服务器postgres进程
./pgtools chkps -u <系统用户> -w <密码> -m <ip地址> -p <ssh端口号>

# 检查数据库主备关系
./pgtools chkms -U <数据库用户> -W <密码> -M <ip地址> -P <数据库端口号> -D <数据库名称>

# 查看服务器切换日志
./pgtools chklog -u <系统用户> -w <密码> -m <ip地址> -p <ssh端口号> -T <数据库类型pg11/pg13>

# 修改记录主备状态信息表
./pgtools chtab <数据库用户> -W <密码> -M <ip地址> -P <数据库端口号> -D <数据库名称> -T <数据库类型pg13>

# 修复新备库
./pgtools repairms -u <系统用户> -w <密码> -m <新备库ip地址> -p <ssh端口号> -a <新主库ip地址> -b <新主库数据库端口号> -T <数据库类型pg11/pg13>

# pgreind同步主库数据再进行主从修复
./pgtools pgrewind -u <系统用户> -w <密码> -m <新备库ip地址> -p <ssh端口号> -a <新主库ip地址> -b <新主库数据库端口号> -T <数据库类型pg11/pg13>

# rebuild重构备库再进行主从修复
./pgtools rebuild -u <系统用户> -w <密码> -m <新备库ip地址> -p <ssh端口号> -a <新主库ip地址> -b <新主库数据库端口号> -T <数据库类型pg11/pg13>

# 修复新备库keepalived服务
./pgtools repairska -u <系统用户> -w <密码> -m <新备库ip地址> -p <ssh端口号> -T <数据库类型pg11/pg13>

# 修复新主库keepalived服务
./pgtools repairmka -u <系统用户> -w <密码> -m <新主库ip地址> -p <ssh端口号> -T <数据库类型pg11/pg13>
```