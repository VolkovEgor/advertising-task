## Нагрузочное тестирование

```
$ ab -c 10 -n 10000 localhost:9000/api/adverts?page=1
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)


Server Software:        
Server Hostname:        localhost
Server Port:            9000

Document Path:          /api/adverts?page=1
Document Length:        643 bytes

Concurrency Level:      10
Time taken for tests:   151.522 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      7670000 bytes
HTML transferred:       6430000 bytes
Requests per second:    66.00 [#/sec] (mean)
Time per request:       151.522 [ms] (mean)
Time per request:       15.152 [ms] (mean, across all concurrent requests)
Transfer rate:          49.43 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.4      0       1
Processing:    19  151  26.0    145     471
Waiting:       17  151  25.9    145     471
Total:         20  151  26.0    145     472

Percentage of the requests served within a certain time (ms)
  50%    145
  66%    153
  75%    162
  80%    168
  90%    181
  95%    195
  98%    214
  99%    240
 100%    472 (longest request)
```