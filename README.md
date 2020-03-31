plate
=====
Serves files from a local directory.

```
$ go get -u rmazur.io/plate

$ plate
2020/03/31 20:34:09 Staring the HTTP server...
^C2020/03/31 20:34:13 Shutting down...

$ plate --help
Usage of plate:
  -address string
    	Network address to listen on (all addresses by default)
  -dir string
    	Directory to serve files from (default ".")
  -port int
    	TCP port to listen on (default 8080)
```
