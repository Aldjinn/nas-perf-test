# nas-perf-test

NAS Performance Tester

````bash
# SSD
F:\Projects\go\nas-perf-test>go run main.go C:\Temp 
Wrote 10% of files to filesystem
Wrote 20% of files to filesystem
Wrote 30% of files to filesystem
Wrote 40% of files to filesystem
Wrote 50% of files to filesystem
Wrote 60% of files to filesystem
Wrote 70% of files to filesystem
Wrote 80% of files to filesystem
Wrote 90% of files to filesystem
Wrote 100% of files to filesystem
Created 1000 files of size 128KB and wrote them to filesystem in 1.8912938s
Deleted 1000 files and folder from filesystem in 115.972ms

# HDD
F:\Projects\go\nas-perf-test>go run main.go D:\Temp
Wrote 10% of files to filesystem
Wrote 20% of files to filesystem
Wrote 30% of files to filesystem
Wrote 40% of files to filesystem
Wrote 50% of files to filesystem
Wrote 60% of files to filesystem
Wrote 70% of files to filesystem
Wrote 80% of files to filesystem
Wrote 90% of files to filesystem
Wrote 100% of files to filesystem
Created 1000 files of size 128KB and wrote them to filesystem in 1.6605145s
Deleted 1000 files and folder from filesystem in 287.115ms

# NAS with single HDD
F:\Projects\go\nas-perf-test>go run main.go N:\Temp 
Wrote 10% of files to filesystem
Wrote 20% of files to filesystem
Wrote 30% of files to filesystem
Wrote 40% of files to filesystem
Wrote 50% of files to filesystem
Wrote 60% of files to filesystem
Wrote 70% of files to filesystem
Wrote 80% of files to filesystem
Wrote 90% of files to filesystem
Wrote 100% of files to filesystem
Created 1000 files of size 128KB and wrote them to filesystem in 53.2053453s
Deleted 1000 files and folder from filesystem in 17.2874992s
```
