# nas-perf-test

Creates a lot of small files and writes (+deletes) them to the mounted filesystem with duration measurement.

```bash
docker build . -t nas-perf-test
docker run --rm -v ${PWD}:/nas-perf-test-folder -e NUM_FILES=4096 nas-perf-test
```

## Examples

```bash
nc11701d@vc60:~/nas/vc60 » docker run --rm -v ~/nas/vc60:/nas-perf-test-folder -e NUM_FILES=4096 nas-perf-test
Wrote 2% of files to filesystem
Wrote 5% of files to filesystem
Wrote 7% of files to filesystem
Wrote 10% of files to filesystem
Wrote 12% of files to filesystem
Wrote 15% of files to filesystem
Wrote 17% of files to filesystem
Wrote 20% of files to filesystem
Wrote 22% of files to filesystem
Wrote 24% of files to filesystem
Wrote 27% of files to filesystem
Wrote 29% of files to filesystem
Wrote 32% of files to filesystem
Wrote 34% of files to filesystem
Wrote 37% of files to filesystem
Wrote 39% of files to filesystem
Wrote 42% of files to filesystem
Wrote 44% of files to filesystem
Wrote 46% of files to filesystem
Wrote 49% of files to filesystem
Wrote 51% of files to filesystem
Wrote 54% of files to filesystem
Wrote 56% of files to filesystem
Wrote 59% of files to filesystem
Wrote 61% of files to filesystem
Wrote 63% of files to filesystem
Wrote 66% of files to filesystem
Wrote 68% of files to filesystem
Wrote 71% of files to filesystem
Wrote 73% of files to filesystem
Wrote 76% of files to filesystem
Wrote 78% of files to filesystem
Wrote 81% of files to filesystem
Wrote 83% of files to filesystem
Wrote 85% of files to filesystem
Wrote 88% of files to filesystem
Wrote 90% of files to filesystem
Wrote 93% of files to filesystem
Wrote 95% of files to filesystem
Wrote 98% of files to filesystem
Created 4096 files of size 128KB and wrote them to filesystem in 3m1.657610638s
Deleted 4096 files and folder from filesystem in 30.767473897s
```
