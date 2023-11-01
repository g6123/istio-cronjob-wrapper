# istio-cronjob-wrapper

CronJob container entrypoint with istio-proxy support.

See https://stackoverflow.com/q/72645650, or https://github.com/istio/istio/issues/11659.

## Usage

```dockerfile
COPY bin/istio-cronjob-wrapper-linux-amd64 /istio-cronjob-wrapper
# ...
ENTRYPOINT ["/istio-cronjob-wrapper"]
```
