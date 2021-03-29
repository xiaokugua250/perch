## shell  cheet 卡片

1. 特殊变量

| 变量     | 含义 |
| ----------- | ----------- |
|$0	|脚本名字|
|$1	|位置参数 #1|
|$2 - $9|	位置参数 #2 - #9|
|${10}|	位置参数 #10|
|$#|	位置参数的个数|
|"$*"|	所有的位置参数(作为单个字符串) *|
|"$@"|	所有的位置参数(每个都作为独立的字符串)|
|${#*}	|传递到脚本中的命令行参数的个数|
|${#@}	|传递到脚本中的命令行参数的个数|
|$?	|返回值|
|$$|	脚本的进程ID(PID)|
|$-	|传递到脚本中的标志(使用set)|
|$_	|之前命令的最后一个参数|
|$!	|运行在后台的最后一个作业的进程ID(PID)|



2. 比较操作

|操作|描述|---|操作|描述|
| --- | --- | --- | --- | ---|	 	 	 	 
|算术比较	||| 	 	字符串比较|	 
|-eq	|等于	 ||	=	|等于|
| 	 	| 	||==||	等于
|-ne|	不等于	 	||!=|	不等于
|-lt|	小于	 	||\<	|小于 (ASCII) *
|-le|	小于等于	 	 	 
|-gt|	大于	 	||\>	|大于 (ASCII) *
|-ge|	大于等于	 	 	 
| 	 |	|| 	-z|	字符串为空
| 	 |||	 	-n|	字符串不为空
| 	 	 	 	 
|算术比较	|双括号(( ... ))结构	 	 	 
|>	|大于	 |
|>=	|大于等于|
|<|	小于	 |	 	 
|<=	|小于等于|


3. 文件类型操作

|操作|测试条件|---|操作|测试条件|
| --- | --- | --- | --- | ---|
|-e	|文件是否存在||	 	-s	|文件大小不为0|
|-f	|是一个标准文件||	| 	 	 
|-d	|是一个目录	 ||	-r|	文件具有读权限|
|-h	|文件是一个符号链接	 ||	-w|	文件具有写权限|
|-L	|文件是一个符号链接	 ||	-x|	文件具有执行权限|
|-b	|文件是一个块设备	 	|| 	 ||
|-c	|文件是一个字符设备	 	||-g	|设置了sgid标记|
|-p	|文件是一个管道	|| 	-u	|设置了suid标记|
|-S	|文件是一个socket||	 	-k|	设置了"粘贴位"|
|-t	|文件与一个终端相关联	|||| 	 	 
| 	| 	 	 	 
|-N	|从这个文件最后一次被读取之后, 它被修改过	|| 	F1 |-nt F2|	文件F1比文件F2新 *|
|-O	|这个文件的宿主是你	 ||	F1 -ot F2	|文件F1比文件F2旧 *|
|-G	|文件的组id与你所属的组相同	 	||F1 -ef F2|	文件F1|和文件F2都是同一个文件的硬链接 *
| 	 |	 	 	 
|!	|"非" (反转上边的测试结果)	 	 |	 |

4.参数扩展与替换

|表达式|含义|
| --- | --- |
|${var} |变量var的值, 与$var相同|
| 	 ||
|${var-DEFAULT}	|如果var没有被声明, 那么就以$DEFAULT作为其值 *|
|${var:-DEFAULT}	|如果var没有被声明, 或者其值为空, 那么就以$DEFAULT作为其值 *|
| 	 |
|${var=DEFAULT}|	如果var没有被声明, 那么就以$DEFAULT作为其值 *|
|${var:=DEFAULT}|	如果var没有被声明, 或者其值为空, 那么就以$DEFAULT作为其值 *|
| 	 |
|${var+OTHER}|	如果var声明了, 那么其值就是$OTHER, 否则就为null字符串|
|${var:+OTHER}|	如果var被设置了, 那么其值就是$OTHER, 否则就为null字符串|
| 	 |
|${var?ERR_MSG}|	如果var没被声明, 那么就打印$ERR_MSG *|
|${var:?ERR_MSG}	|如果var没被设置, 那么就打印$ERR_MSG *|
| 	 |
|${!varprefix*}	|匹配之前所有以varprefix开头进行声明的变量|
|${!varprefix@}	|匹配之前所有以varprefix开头进行声明的变量|

5.字符串操作
|表达式|含义|
| --- | --- |
|${#string}	|$string的长度 |
|${string:position}	|在$string中, 从位置$position开始提取子串|
|${string:position:length}	|在$string中, 从位置$position开始提取长度为$length的子串|
| 	 |
|${string#substring}	|从变量$string的开头, 删除最短匹配|$substring的子串
|${string##substring}	|从变量$string的开头, 删除最长匹配|$substring的子串
|${string%substring}	|从变量$string的结尾, 删除最短匹配|$substring的子串
|${string%%substring}	|从变量$string的结尾, 删除最长匹配|$substring的子串
| 	 
|${string/substring/replacement}	|使用$replacement, 来代替第一个匹配的$substring|
|${string//substring/replacement}	|使用$replacement, 代替所有匹配的$substring|
|${string/#substring/replacement}	|如果$string的前缀匹配$substring, 那么就用$replacement来代替匹配到的$substring|
|${string/%substring/replacement}	|如果$string的后缀匹配|$substring, 那么就用$replacement来代替匹配到的$substring|
| 	 |
| 	 |
|expr match "$string" '$substring'|	匹配\$string开头的$substring*的长度|
|expr "$string" : '$substring'	|匹配\$string开头的$substring*的长度|
|expr index "$string" $substring	|在\$string中匹配到的$substring的第一个字符出现的位置|
|expr substr $string $position $length	|在\$string中从位置\$position开始提取长度为$length的子串|
|expr match "$string" '\\(\$substring\)'|	从\$string的开头位置提取$substring*|
|expr "$string" : '\\(\$substring\\)'	|从\$string的开头位置提取$substring*|
|expr match "$string" '.*\\(\$substring\\)'|	从\$string的结尾提取$substring*|
|expr "$string" : '.*\\(\$substring\\)'|	从\$string的结尾提取$substring*|


6.结构汇总
|表达式|含义|
| --- | --- |
|中括号	 ||
|if [ CONDITION ]	|测试结构|
|if [[ CONDITION ]]	|扩展的测试结构|
|Array[1]=element1|	数组初始化|
|[a-z]	|正则表达式的字符范围	 |
|大括号	 ||
|${variable}|	参数替换|
|${!variable}	|间接变量引用|
|{ command1; command2; . . . commandN; }	|代码块|
|{string1,string2,string3,...}	|大括号扩展|
|圆括号	 ||
|( command1; command2 )	|子shell中执行的命令组|
|Array=(element1 element2 element3)	|数组初始化|
|result=$(COMMAND)	|在子shell中执行命令, 并将结果赋值给变量|
|>(COMMAND)	|进程替换|
|<(COMMAND)|	进程替换|
|双圆括号	 ||
|(( var = 78 ))|	整型运算|
|var=$(( 20 + 5 ))|	整型运算, 并将结果赋值给变量|
| 	 |
|引号	 ||
|"$variable"|	"弱"引用|
|'string'	|"强"引用|
| 	 
|后置引用	 ||
|result=`COMMAND`	在子shell中运行命令, 并将结果赋值给变量|

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