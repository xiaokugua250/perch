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


7. sed 操作 cheet

7.1 基本操作
|操作符 |  名字 | 执行效果|
| --- | --- |--- |
|[地址范围]/p	|打印|	打印[指定的地址范围]|
|[地址范围]/d	|删除	|删除[指定的地址范围]|
|s/pattern1/pattern2/	|替换	|将指定行中, 将第一个匹配到的pattern1, 替换|为pattern2.|
|[地址范围]/s/pattern1/pattern2/	|替换|	在地址范围指定的每一行中, 将第一个匹|配到的pattern1, 替换为pattern2.|
|[地址范围]/y/pattern1/pattern2/	|transform|	在地址范围指定的每一行中, |将pattern1中的每个匹配到pattern2的字符都使用pattern2的相应字符作替换. |(等|价于tr命令)|
|g	|全局|	在每个匹配的输入行中, 将每个模式匹配都作相应的操作. (译者注: 不只局限于第一个匹配)

7.2 操作示例
|操作|效果|
| --- | --- |
|8d	|删除输入的第8行.|
|/^$/d|	删除所有空行.|
|1,/^$/d	|从输入的开头一直删除到第1个空行(第一个空行也删除掉).|
|/Jones/p	|只打印那些包含"Jones"的行(使用-n选项).|
|s/Windows/Linux/	|在每个输入行中, 将第一个出现的"Windows"实例替换为|"Linux".|
|s/BSOD/stability/g	|在每个输入行中, 将所有"BSOD"都替换为"stability".|
|s/ *$//	|删除掉每行结尾的所有空格.|
|s/00*/0/g|	将所有连续出现的0都压缩成单个的0.|
|/GUI/d	|删除掉所有包含"GUI"的行.|
|s/GUI//g	|将所有"GUI"都删除掉, 并保持剩余部分的完整性.|

8 脚本退出码 cheet
| 退出码|含义|示例|注释|
| --- | --- | --- | --- |
|1	|通用错误|	let "var1 = 1/0"	|各种各样的错误都可能使用这个退出码, 比如"除0错误"|
|2	|shell内建命令使用错误(Bash文档上有说明)|	| 	很少看到, 通常情况下退出码都为1|
|126|	命令调用不能执行	|| 	程序或命令的权限是不可执行的|
|127|	"command not found"	 ||	估计是$PATH不对, 或者是拼写错误|
|128|	exit的参数错误	|exit 3.14159	|exit只能以整数作为参数, 范围是0 - |255(见脚注)
|128+n|	信号"n"的致命错误|	kill -9 脚本的$PPID|	$? 返回137(128 + 9)
|130	|用Control-C来结束脚本	| |	Control-C是信号2的致命错误, (130 = 128 + |2, 见上边)|
|255*|	超出范围的退出状态	|exit -1	|exit命令只能够接受范围是0 - 255的整数作为参数|

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

3 shell并行计算

3.1 &+wait 模式
`Shell实现并发就是通过&命令符将循环体的命令放入后台运行，但是这种方法对线程并发数不可控，系统也会随着高并发压力的不断攀升，处理速度会变得越来越慢，所以这种方法针对少量的文件可行，但是一旦文件数量大，处理速度是很慢的。`
```
 ## for 循环版本
#!/bin/bash
# bam to bed

date # 脚本开始时间

for ((i=1;i<=1000;i++))
do
{
    bam2bed  #这里执行自己的脚本
        echo " $i finished! "  
 }&              #用{}把循环体括起来，后加一个&符号，代表每次循环都把命令放入后台运行
                 #一旦放入后台，就意味着{}里面的命令交给操作系统的一个线程处理了
                 #循环了1000次，就有1000个&将任务放入后台，操作系统会并发1000个线程来处理     
done    
wait             #wait命令表示。等待上面的命令（放入后台的任务）都执行完毕了再往下执行
     
date # 脚本结束时间
```

3.2 管道版本
`
先新建一个FIFO，写入一些字符。一个进程开始前会先从这个FIFO中读走一个字符，执行完之后再写回一个字符。如果FIFO中没有字符，该线程就会等待，fifo就成了一个锁。
`
```
#!/bin/bash
# bam to bed

start_time=`date +%s`  #定义脚本运行的开始时间

tmp_fifofile="/tmp/$$.fifo"
mkfifo $tmp_fifofile   # 新建一个FIFO类型的文件
exec 6<>$tmp_fifofile  # 将FD6指向FIFO类型, 这里6也可以是其它数字
rm $tmp_fifofile  #删也可以，

thread_num=32  # 定义最大线程数

#根据线程总数量设置令牌个数
#事实上就是在fd6中放置了$thread_num个回车符
for ((i=0;i<${thread_num};i++));do
    echo
done >&6

for i in data/*.bam # 找到data文件夹下所有bam格式的文件
do
    # 一个read -u6命令执行一次，就从FD6中减去一个回车符，然后向下执行
    # 当FD6中没有回车符时，就停止，从而实现线程数量控制
    read -u6
    {
        bam2bed # 可以用实际命令代替
        echo >&6 # 当进程结束以后，再向FD6中加上一个回车符，即补上了read -u6减去的那个
    } &
done

wait # 要有wait，等待所有线程结束

stop_time=`date +%s` # 定义脚本运行的结束时间
echo "TIME:`expr $stop_time - $start_time`" # 输出脚本运行时间

exec 6>&- # 关闭FD6，最后一定要记得关闭FIFO
echo "over" # 表示脚本运行结束

```
参考
[1].https://taoyan.netlify.app/post/2020-01-02.%E5%A4%9A%E7%BA%BF%E7%A8%8B%E5%B9%B6%E8%A1%8C%E8%AE%A1%E7%AE%97/
[2].https://jerkwin.github.io/2013/12/14/Bash%E8%84%9A%E6%9C%AC%E5%AE%9E%E7%8E%B0%E6%89%B9%E9%87%8F%E4%BD%9C%E4%B8%9A%E5%B9%B6%E8%A1%8C%E5%8C%96/

4 系统信息查看
4.1 获取系统基本信息
```
#!/bin/bash  
# 获取要监控的本地服务器IP地址  
IP=`ifconfig | grep inet | grep -vE 'inet6|127.0.0.1' | awk '{print $2}'`  
echo "IP地址："$IP  
   
# 获取cpu总核数  
cpu_num=`grep -c "model name" /proc/cpuinfo`  
echo "cpu总核数："$cpu_num  
   
# 1、获取CPU利用率  
################################################  
#us 用户空间占用CPU百分比  
#sy 内核空间占用CPU百分比  
#ni 用户进程空间内改变过优先级的进程占用CPU百分比  
#id 空闲CPU百分比  
#wa 等待输入输出的CPU时间百分比  
#hi 硬件中断  
#si 软件中断  
#################################################  
# 获取用户空间占用CPU百分比  
cpu_user=`top -b -n 1 | grep Cpu | awk '{print $2}' | cut -f 1 -d "%"`  
echo "用户空间占用CPU百分比："$cpu_user  
   
# 获取内核空间占用CPU百分比  
cpu_system=`top -b -n 1 | grep Cpu | awk '{print $4}' | cut -f 1 -d "%"`  
echo "内核空间占用CPU百分比："$cpu_system  
   
# 获取空闲CPU百分比  
cpu_idle=`top -b -n 1 | grep Cpu | awk '{print $8}' | cut -f 1 -d "%"`  
echo "空闲CPU百分比："$cpu_idle  
   
# 获取等待输入输出占CPU百分比  
cpu_iowait=`top -b -n 1 | grep Cpu | awk '{print $10}' | cut -f 1 -d "%"`  
echo "等待输入输出占CPU百分比："$cpu_iowait  
   
#2、获取CPU上下文切换和中断次数  
# 获取CPU中断次数  
cpu_interrupt=`vmstat -n 1 1 | sed -n 3p | awk '{print $11}'`  
echo "CPU中断次数："$cpu_interrupt  
   
# 获取CPU上下文切换次数  
cpu_context_switch=`vmstat -n 1 1 | sed -n 3p | awk '{print $12}'`  
echo "CPU上下文切换次数："$cpu_context_switch  
   
#3、获取CPU负载信息  
# 获取CPU15分钟前到现在的负载平均值  
cpu_load_15min=`uptime | awk '{print $11}' | cut -f 1 -d ','`  
echo "CPU 15分钟前到现在的负载平均值："$cpu_load_15min  
   
# 获取CPU5分钟前到现在的负载平均值  
cpu_load_5min=`uptime | awk '{print $10}' | cut -f 1 -d ','`  
echo "CPU 5分钟前到现在的负载平均值："$cpu_load_5min  
   
# 获取CPU1分钟前到现在的负载平均值  
cpu_load_1min=`uptime | awk '{print $9}' | cut -f 1 -d ','`  
echo "CPU 1分钟前到现在的负载平均值："$cpu_load_1min  
   
# 获取任务队列(就绪状态等待的进程数)  
cpu_task_length=`vmstat -n 1 1 | sed -n 3p | awk '{print $1}'`  
echo "CPU任务队列长度："$cpu_task_length  
   
#4、获取内存信息  
# 获取物理内存总量  
mem_total=`free | grep Mem | awk '{print $2}'`  
echo "物理内存总量："$mem_total  
   
# 获取操作系统已使用内存总量  
mem_sys_used=`free | grep Mem | awk '{print $3}'`  
echo "已使用内存总量(操作系统)："$mem_sys_used  
   
# 获取操作系统未使用内存总量  
mem_sys_free=`free | grep Mem | awk '{print $4}'`  
echo "剩余内存总量(操作系统)："$mem_sys_free  
   
# 获取应用程序已使用的内存总量  
mem_user_used=`free | sed -n 3p | awk '{print $3}'`  
echo "已使用内存总量(应用程序)："$mem_user_used  
   
# 获取应用程序未使用内存总量  
mem_user_free=`free | sed -n 3p | awk '{print $4}'`  
echo "剩余内存总量(应用程序)："$mem_user_free  
   
   
# 获取交换分区总大小  
mem_swap_total=`free | grep Swap | awk '{print $2}'`  
echo "交换分区总大小："$mem_swap_total  
   
# 获取已使用交换分区大小  
mem_swap_used=`free | grep Swap | awk '{print $3}'`  
echo "已使用交换分区大小："$mem_swap_used  
   
# 获取剩余交换分区大小  
mem_swap_free=`free | grep Swap | awk '{print $4}'`  
echo "剩余交换分区大小："$mem_swap_free  
   
  
#5、获取磁盘I/O统计信息  
echo "指定设备(/dev/sda)的统计信息"  
# 每秒向设备发起的读请求次数  
disk_sda_rs=`iostat -kx | grep sda| awk '{print $4}'`  
echo "每秒向设备发起的读请求次数："$disk_sda_rs  
   
# 每秒向设备发起的写请求次数  
disk_sda_ws=`iostat -kx | grep sda| awk '{print $5}'`  
echo "每秒向设备发起的写请求次数："$disk_sda_ws  
   
# 向设备发起的I/O请求队列长度平均值  
disk_sda_avgqu_sz=`iostat -kx | grep sda| awk '{print $9}'`  
echo "向设备发起的I/O请求队列长度平均值"$disk_sda_avgqu_sz  
   
# 每次向设备发起的I/O请求平均时间  
disk_sda_await=`iostat -kx | grep sda| awk '{print $10}'`  
echo "每次向设备发起的I/O请求平均时间："$disk_sda_await  
   
# 向设备发起的I/O服务时间均值  
disk_sda_svctm=`iostat -kx | grep sda| awk '{print $11}'`  
echo "向设备发起的I/O服务时间均值："$disk_sda_svctm  
   
# 向设备发起I/O请求的CPU时间百分占比  
disk_sda_util=`iostat -kx | grep sda| awk '{print $12}'`  
echo "向设备发起I/O请求的CPU时间百分占比："$disk_sda_util  
```