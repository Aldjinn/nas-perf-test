# nas-perf-test

NAS Performance Tester

```bash
docker build . -t nas-perf-test
docker run --rm -v ${PWD}:/nas-perf-test-folder -e NUM_FILES=4096 nas-perf-test
```
