<h1>CMD-IP</h1>

This is a simple cmd application **writing in golang** to retrieve ips and host names from given url

Command example to get ips from github.com:

```
go run main.go ip --host github.com
```
return 
20.201.28.151

to get server name:
```
go run main.go hostname --host github.com
```
return:

dns1.p08.nsone.net.

dns2.p08.nsone.net.

dns3.p08.nsone.net.

dns4.p08.nsone.net.

ns-1283.awsdns-32.org.

ns-1707.awsdns-21.co.uk.

ns-421.awsdns-52.com.

ns-520.awsdns-01.net.
