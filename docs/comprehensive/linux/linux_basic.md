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

```
拒绝所有针对该端口的访问
iptables -I INPUT -p tcp --dport ${PORT} -j DROP 

允许所有针对该端口的访问
iptables -I INPUT -p tcp --dport ${PORT} -j ACCEPT
```
### 


## 参考
[1].https://www.shuzhiduo.com/A/rV57WBVR5P/
[2].https://github.com/jlevy/the-art-of-command-line/blob/master/README-zh.md