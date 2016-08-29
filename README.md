#My Unique ID

This tool generates unique ids similar to uuidv1 (based on timestamp and generator machine id).

Uniqueness depends on two factors:
  - The muid binary may never see the same system clock time twice
  - More than one muid binary may not be running simultaneously with the same machineid.

Machine ID's may be randomly chosen when the muid generator container is created, or a MAC address may be used.

###Usage
Put the code in your `$GOPATH`:

```
go get github.com/therealplato/muid
# (or...)
cp -r . $GOPATH/src/github.com/therealplato/muid

cd $GOPATH/src/github.com/therealplato/muid
```

Build the binary:

```
go build cmd/muid.go
```

Run the binary with default configuration:

```
./muid    #./muid.exe on windows
```

Inspect memory usage: see [MEMORY](https://github.com/therealplato/muid/blob/master/MEMORY)

### Configuration
The `muid` binary supports these flags:


```
$ ./muid.exe -h
Usage of C:\Users\there\go\src\github.com\therealplato\muid\muid.exe:
  -mid string
        hexadecimal machine id (default "1234567890abcdef")
  -n int
        generate this many ID's (default 1000)
  -profile
        enable cpu and memory profiling
  -q    mute stdout
  -sizemid int
        use this many bytes for machine ID (default 8)
  -sizets int
        use this many bytes for LSBs of nanosecond timestamp (default 8)
 ```
