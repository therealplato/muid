there@iamconfuse MINGW64 ~/go/src/github.com/therealplato/muid/cmd (master)
$ time ./muid.exe -profile -q -n 50000000

real    0m5.491s
user    0m0.000s
sys     0m0.015s

there@iamconfuse MINGW64 ~/go/src/github.com/therealplato/muid/cmd (master)
$ go tool pprof muid.exe memprofile.out
Entering interactive mode (type "help" for commands)
(pprof) top2
1.12GB of 1.12GB total (99.91%)
Dropped 1 node (cum <= 0.01GB)
Showing top 2 nodes out of 4 (cum >= 1.12GB)
      flat  flat%   sum%        cum   cum%
    1.12GB 99.91% 99.91%     1.12GB   100%  github.com/therealplato/muid.(*Generator).Bulk
         0     0% 99.91%     1.12GB   100%  main.main