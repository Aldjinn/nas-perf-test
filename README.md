# nas-perf-test

NAS Performance Tester

```bash
docker build . -t nas-perf-test
docker run --rm -v ${pwd}/:/nas-perf-test-folder -e NUM_FILES=512 nas-perf-test
```
