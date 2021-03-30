# linux 基础
<!-- TOC -->

- [linux 基础](#linux-基础)
        - [linux  shell脚本](#linux--shell脚本)
            - [shell 并发](#shell-并发)
        - [](#)
    - [参考](#参考)

<!-- /TOC -->
### linux  shell脚本
#### shell 并发

1. iptables操作
```
拒绝所有针对该端口的访问
iptables -I INPUT -p tcp --dport ${PORT} -j DROP 

允许所有针对该端口的访问
iptables -I INPUT -p tcp --dport ${PORT} -j ACCEPT
```

2. top性能分析
```
# 查找内存消耗最⼤的进程
top -c -b -o +%MEM | head -n 20 | tail -15 

```

3. ps进程管理

```
# 显示进程关系
ps  auxf 

# 查找僵尸进程
ps aux | awk '{print $8 " " $2} ' | grep -w Z

# 查找内存消耗最大的进程

ps aux --sort -rss | head
```
4. 数据备份
```
#!/bin/sh
umount /mnt/backup
mount /dev/sdb1 /mnt/backup
if [ `date +%d` = '01' ] #每⽉1号进⾏完全备份
then
    bakdir="/mnt/bak/daybak/month/"`date +%m%d`
    zl="" #进⾏完全备份
else
    backup_dir="/mnt/backup/"`date +%d`
    zl="-N "`date +'%Y-%m-01 00:00:01'`; #差异备份
#zl="-N "`date -d '-1 day' +'%Y-%m-%d 00:00:01'` #⽇增量备份
fi
tar "${zl}" -czf ${backup_dir}/www.tgz /var/www
umount /mnt/backup
```

5. 系统信息获取
```
# cpu 核心数
cat /proc/cpuinfo | grep processor | wc -l

```

6. 随机数应用
```
# 随机密码
cat /dev/urandom | head -1 | md5sum | head -c 8

# 随机密码
od -N 4 -t x4 /dev/random | head -1 | awk '{print $2}'


```

### 


## 参考
[1].https://www.shuzhiduo.com/A/rV57WBVR5P/
[2].https://github.com/jlevy/the-art-of-command-line/blob/master/README-zh.md