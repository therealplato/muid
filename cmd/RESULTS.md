#### First try with generator.Bulk():
$ time ./muid.exe -n 100000000
real    0m14.758s

$ go tool pprof muid.exe memprofile.out
Entering interactive mode (type "help" for commands)
(pprof) top1
2.24GB of 2.24GB total (  100%)
Dropped 5 nodes (cum <= 0.01GB)
Showing top 1 nodes out of 4 (cum >= 2.24GB)
      flat  flat%   sum%        cum   cum%
    2.24GB   100%   100%     2.24GB   100%  github.com/therealplato/muid.(*Generator).Bulk


// Using 4 bytes mid instead of default 8: