## shell 脚本收集

1. 备份最后一天修改的文件 [参考：https://teliute.org/linux/abs-3.9.1/special-chars.html]
```

  1 #!/bin/bash
  2 
  3 #  在一个"tarball"中(经过tar和gzip处理过的文件)
  4 #+ 备份最后24小时当前目录下d所有修改的文件. 
  5 
  6 BACKUPFILE=backup-$(date +%m-%d-%Y)
  7 #                 在备份文件中嵌入时间.
  8 #                 Thanks, Joshua Tschida, for the idea.
  9 archive=${1:-$BACKUPFILE}
 10 #  如果在命令行中没有指定备份文件的文件名,
 11 #+ 那么将默认使用"backup-MM-DD-YYYY.tar.gz".
 12 
 13 tar cvf - `find . -mtime -1 -type f -print` > $archive.tar
 14 gzip $archive.tar
 15 echo "Directory $PWD backed up in archive file \"$archive.tar.gz\"."
 16 
 17 
 18 #  Stephane Chazelas指出上边代码,
 19 #+ 如果在发现太多的文件的时候, 或者是如果文件
 20 #+ 名包括空格的时候, 将执行失败.
 21 
 22 # Stephane Chazelas建议使用下边的两种代码之一:
 23 # -------------------------------------------------------------------
 24 #   find . -mtime -1 -type f -print0 | xargs -0 tar rvf "$archive.tar"
 25 #      使用gnu版本的"find".
 26 
 27 
 28 #   find . -mtime -1 -type f -exec tar rvf "$archive.tar" '{}' \;
 29 #         对于其他风格的UNIX便于移植, 但是比较慢.
 30 # -------------------------------------------------------------------
 31 
 32 
 33 exit 0
```

2. 扫描远程机器上的identd服务进程
```


  1 #! /bin/sh
  2 ## 使用netcat工具写的和DaveG写的ident-scan扫描器有同等功能的东西. 噢, 他会被气死的. 
  3 ## 参数: target port [port port port ...]
  4 ## 标准输出和标准输入被混到一块.
  5 ##
  6 ##  优点: 运行起来比ident-scan慢, 这样使远程机器inetd进程更不易注意而不会产生警告, 
  7 ##+ 并且只有很少的知名端口会被指定. 
  8 ##  缺点: 要求数字端口参数, 输出中无法区分标准输出和标准错误, 
  9 ##+ 并且当远程服务监听在很高的端口时无法工作. 
 10 # 脚本作者: Hobbit <hobbit@avian.org>
 11 # 已征得作者同意在ABS指南中使用. 
 12 
 13 # ---------------------------------------------------
 14 E_BADARGS=65       # 至少需要两个参数. 
 15 TWO_WINKS=2        # 睡眠多长时间. 
 16 THREE_WINKS=3
 17 IDPORT=113         # indent协议的认证端口. 
 18 RAND1=999
 19 RAND2=31337
 20 TIMEOUT0=9
 21 TIMEOUT1=8
 22 TIMEOUT2=4
 23 # ---------------------------------------------------
 24 
 25 case "${2}" in
 26   "" ) echo "Need HOST and at least one PORT." ; exit $E_BADARGS ;;
 27 esac
 28 
 29 # 测试目标主机看是否运行了identd守护进程.
 30 nc -z -w $TIMEOUT0 "$1" $IDPORT || { echo "Oops, $1 isn't running identd." ; exit 0 ; }
 31 #  -z 选项扫描监听进程.
 32 #     -w $TIMEOUT = 尝试连接多长时间.
 33 
 34 # 产生一个随机的本地起点源端口.
 35 RP=`expr $$ % $RAND1 + $RAND2`
 36 
 37 TRG="$1"
 38 shift
 39 
 40 while test "$1" ; do
 41   nc -v -w $TIMEOUT1 -p ${RP} "$TRG" ${1} < /dev/null > /dev/null &
 42   PROC=$!
 43   sleep $THREE_WINKS
 44   echo "${1},${RP}" | nc -w $TIMEOUT2 -r "$TRG" $IDPORT 2>&1
 45   sleep $TWO_WINKS
 46 
 47 # 这看上去是不是像个残疾的脚本或是其他类似的东西... ?
 48 # ABS作者评注 : "这不是真的那么糟糕,
 49 #+               事实上, 做得非常聪明."
 50 
 51   kill -HUP $PROC
 52   RP=`expr ${RP} + 1`
 53   shift
 54 done
 55 
 56 exit $?
 57 
 58 #  注意事项:
 59 #  ---------
 60 
 61 #  试着把第30行去掉, 
 62 #+ 然后以"localhost.localdomain 25"为参数来运行脚本.
 63 
 64 #  关于Hobbit写的更多'nc'例子脚本,
 65 #+ 可以在以下文档中找到:
 66 #+ /usr/share/doc/nc-X.XX/scripts 目录
```