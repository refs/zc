# Testing that the multicast works:

```console
~/code/mdns master*
❯ iperf -c 224.0.0.1 -p 5353 -u -T 32 -t 3 -i 1
------------------------------------------------------------
Client connecting to 224.0.0.1, UDP port 5353
Sending 1470 byte datagrams, IPG target: 11215.21 us (kalman adjust)
Setting multicast TTL to 32
UDP buffer size: 9.00 KByte (default)
------------------------------------------------------------
[  4] local 192.168.0.17 port 54931 connected with 224.0.0.1 port 5353
[ ID] Interval       Transfer     Bandwidth
[  4]  0.0- 1.0 sec   131 KBytes  1.07 Mbits/sec
[  4]  1.0- 2.0 sec   128 KBytes  1.05 Mbits/sec
[  4]  0.0- 3.0 sec   385 KBytes  1.05 Mbits/sec
[  4] Sent 268 datagrams
```

Another session with the multicast server running:

```console
~/code/mdns master*
❯ go run main.go
2019-12-26T19:38:00+01:00 INF 83 bytes written to 192.168.0.17:54931
2019-12-26T19:38:00+01:00 INF 83 bytes written to 192.168.0.17:54931
2019-12-26T19:38:00+01:00 INF 83 bytes written to 192.168.0.17:54931
2019-12-26T19:38:00+01:00 INF 83 bytes written to 192.168.0.17:54931
2019-12-26T19:38:00+01:00 INF 83 bytes written to 192.168.0.17:54931
2019-12-26T19:38:00+01:00 INF 83 bytes written to 192.168.0.17:54931
[...]
```