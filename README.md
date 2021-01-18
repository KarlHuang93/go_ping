# go_ping 


网络监控，包括丢包率、发送的往返时间的标准偏差、是通过此pinger发送的平均往返时间、是通过此pinger发送的最短往返时间等

`注意`: 需要sudo root

ping命令在运行中采用了ICMP协议，需要发送ICMP报文。但是只有root用户才能建立ICMP报文。
