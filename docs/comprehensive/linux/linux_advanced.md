# linux 高级
<!-- TOC -->

- [linux 高级](#linux-高级)
        - [linux 网络操作](#linux-网络操作)
            - [iptables 相关操作](#iptables-相关操作)
        - [](#)
    - [参考](#参考)

<!-- /TOC -->
### linux 网络操作
#### iptables 相关操作
```
拒绝所有针对该端口的访问
iptables -I INPUT -p tcp --dport ${PORT} -j DROP 

允许所有针对该端口的访问
iptables -I INPUT -p tcp --dport ${PORT} -j ACCEPT
```
### 


## 参考
[1].https://www.shuzhiduo.com/A/rV57WBVR5P/