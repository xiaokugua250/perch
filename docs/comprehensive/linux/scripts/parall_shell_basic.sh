#!/bin/bash
#********************************************************************
#Author: liangdu1992@gmail.com
#website： www.z-gour.com
#Date： 2021-01-20
#FileName： parall_shell_basick.sh
#Description： & 和 wait实现多线程#********************************************************************

date

for ((i = 1; i < 1000; i++)); do
    {
        echo "$i finished"
    } &
done
wait
date
